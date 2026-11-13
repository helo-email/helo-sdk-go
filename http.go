package helo

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math"
	mathrand "math/rand/v2"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

// moduleImportPath is this module's import path, used to find the SDK's own
// version in the build info of whatever binary embeds it.
const moduleImportPath = "github.com/helo-email/helo-sdk-go"

// userAgentProduct names the SDK in the User-Agent header.
const userAgentProduct = "helo-sdk-go"

// DefaultMaxRetries is how many times a request is retried unless overridden
// via WithMaxRetries.
const DefaultMaxRetries = 2

// DefaultTimeout is how long each attempt may take, response body included,
// unless WithHTTPClient supplies a client with its own Timeout. http.DefaultClient
// has none, so a hung connection would otherwise block forever.
const DefaultTimeout = 30 * time.Second

var defaultHTTPClient = &http.Client{Timeout: DefaultTimeout}

// The retry policy every Helo SDK shares.
const (
	initialRetryDelay = 0.5 * float64(time.Second)
	maxRetryDelay     = 8.0 * float64(time.Second)
)

var (
	retryStatuses     = map[int]bool{408: true, 429: true, 500: true, 502: true, 503: true, 504: true}
	idempotentMethods = map[string]bool{"GET": true, "HEAD": true, "OPTIONS": true, "PUT": true, "DELETE": true}
)

// Client is the HTTP client used by all resource services.
//
// A request is retried (up to MaxRetries times, with jittered backoff) after a
// network failure or a 408, 429 or 5xx response, but only when repeating it
// cannot do anything twice: its method is idempotent, or it carries an
// idempotency key. A 429 is retried regardless, since the server refused it
// without acting on it. A Retry-After longer than the SDK waits is not retried
// early; the error goes to the caller.
type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client

	// MaxRetries is how many times a failed request is retried; 0 disables retries.
	MaxRetries int

	// UserAgent is sent with every request. It defaults to defaultUserAgent();
	// a program wrapping the SDK (the CLI, say) sets its own via WithUserAgent.
	UserAgent string
}

// defaultUserAgent identifies the SDK and the Go runtime carrying it. There is
// no generated version constant — the module is versioned by git tag — so the
// version is read back from the build info the toolchain embeds.
var defaultUserAgent = buildUserAgent()

func buildUserAgent() string {
	version := "unknown"
	if info, ok := debug.ReadBuildInfo(); ok {
		if info.Main.Path == moduleImportPath && info.Main.Version != "" {
			version = info.Main.Version
		}
		for _, dep := range info.Deps {
			if dep.Path == moduleImportPath && dep.Version != "" {
				version = dep.Version
				break
			}
		}
	}
	return fmt.Sprintf("%s/%s (%s; %s/%s)", userAgentProduct, sanitizeVersion(version), runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

// A module version can be a placeholder like "(devel)" for an untagged local
// build. The parentheses there would open a comment the User-Agent never
// closes, so strip anything that is not a token character.
func sanitizeVersion(version string) string {
	cleaned := strings.Map(func(r rune) rune {
		if r == '(' || r == ')' || r == ' ' {
			return -1
		}
		return r
	}, version)
	if cleaned == "" {
		return "unknown"
	}
	return cleaned
}

// idempotencyKey returns a random (version 4) UUID, for an idempotency key the
// caller didn't supply.
func idempotencyKey() string {
	var b [16]byte
	_, _ = rand.Read(b[:])
	b[6] = b[6]&0x0f | 0x40
	b[8] = b[8]&0x3f | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// requestOptions carries optional body, query and header parameters for a single call.
type requestOptions struct {
	body    any
	query   map[string]any
	headers map[string]string
}

// requestOption mutates a requestOptions value.
type requestOption func(*requestOptions)

func withBody(body any) requestOption {
	return func(o *requestOptions) { o.body = body }
}

func withQuery(query map[string]any) requestOption {
	return func(o *requestOptions) { o.query = query }
}

func withHeaders(headers map[string]string) requestOption {
	return func(o *requestOptions) { o.headers = headers }
}

// request performs an HTTP request and decodes the JSON response into out (if non-nil).
func (c *Client) request(ctx context.Context, method, path string, out any, opts ...requestOption) error {
	options := &requestOptions{}
	for _, opt := range opts {
		opt(options)
	}

	endpoint, err := url.Parse(strings.TrimRight(c.BaseURL, "/") + path)
	if err != nil {
		return fmt.Errorf("invalid url: %w", err)
	}

	if len(options.query) > 0 {
		q := endpoint.Query()
		for key, value := range options.query {
			if value == nil {
				continue
			}
			switch v := value.(type) {
			case string:
				if v == "" {
					continue
				}
				q.Set(key, v)
			case []string:
				if len(v) == 0 {
					continue
				}
				q.Set(key, strings.Join(v, ","))
			default:
				// Any other slice, such as a []EventType, is comma-joined like a
				// []string (every array query parameter is explode: false), not
				// printed as "[a b]".
				if rv := reflect.ValueOf(v); rv.Kind() == reflect.Slice {
					if rv.Len() == 0 {
						continue
					}
					items := make([]string, rv.Len())
					for i := range items {
						items[i] = fmt.Sprint(rv.Index(i).Interface())
					}
					q.Set(key, strings.Join(items, ","))
					continue
				}
				q.Set(key, fmt.Sprintf("%v", v))
			}
		}
		endpoint.RawQuery = q.Encode()
	}

	var body []byte
	if options.body != nil {
		body, err = json.Marshal(options.body)
		if err != nil {
			return fmt.Errorf("marshal request body: %w", err)
		}
	}

	method = strings.ToUpper(method)
	canRepeat := idempotentMethods[method]
	for key, value := range options.headers {
		if value != "" && strings.HasSuffix(strings.ToLower(key), "idempotency-key") {
			canRepeat = true
		}
	}

	for attempt := 0; ; attempt++ {
		resp, err := c.send(ctx, method, endpoint.String(), body, options)
		if err != nil {
			if !canRepeat || attempt >= c.MaxRetries || ctx.Err() != nil {
				return &APIError{Message: fmt.Sprintf("connection failed: %v", err), Cause: err}
			}
			if err := sleep(ctx, backoffDelay(attempt)); err != nil {
				return &APIError{Message: fmt.Sprintf("connection failed: %v", err), Cause: err}
			}
			continue
		}

		if delay, retry := c.retryDelay(resp, attempt, canRepeat); retry {
			resp.Body.Close()
			if err := sleep(ctx, delay); err != nil {
				return &APIError{Message: fmt.Sprintf("connection failed: %v", err), Cause: err}
			}
			continue
		}
		return handleResponse(resp, out)
	}
}

// retryDelay reports how long to wait before retrying resp, and whether to.
func (c *Client) retryDelay(resp *http.Response, attempt int, canRepeat bool) (time.Duration, bool) {
	if attempt >= c.MaxRetries || !retryStatuses[resp.StatusCode] {
		return 0, false
	}
	if resp.StatusCode != http.StatusTooManyRequests && !canRepeat {
		return 0, false
	}
	retryAfter, ok := retryAfterDelay(resp)
	if !ok {
		return backoffDelay(attempt), true
	}
	return retryAfter, retryAfter <= time.Duration(maxRetryDelay)
}

// backoffDelay is exponential backoff with full jitter.
func backoffDelay(attempt int) time.Duration {
	ceiling := math.Min(maxRetryDelay, initialRetryDelay*math.Pow(2, float64(attempt)))
	return time.Duration(mathrand.Float64() * ceiling)
}

// retryAfterDelay reads Retry-After as either seconds or an HTTP date.
func retryAfterDelay(resp *http.Response) (time.Duration, bool) {
	value := resp.Header.Get("Retry-After")
	if value == "" {
		return 0, false
	}
	if seconds, err := strconv.ParseFloat(value, 64); err == nil {
		return time.Duration(math.Max(seconds, 0) * float64(time.Second)), true
	}
	if at, err := http.ParseTime(value); err == nil && time.Until(at) > 0 {
		return time.Until(at), true
	}
	return 0, false
}

// sleep waits for d, or returns the context's error if it ends first.
func sleep(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

// send makes one attempt at the request.
func (c *Client) send(ctx context.Context, method, endpoint string, body []byte, options *requestOptions) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequestWithContext(ctx, method, endpoint, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	} else {
		req.Header.Set("User-Agent", defaultUserAgent)
	}
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}
	if options.body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	for key, value := range options.headers {
		if value == "" {
			continue
		}
		req.Header.Set(key, value)
	}

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = defaultHTTPClient
	}
	return httpClient.Do(req)
}

// handleResponse decodes a final response into out, or returns its error.
func handleResponse(resp *http.Response, out any) error {
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("request failed with status %d", resp.StatusCode),
			Body:       respBody,
		}
		var parsed map[string]any
		if len(respBody) > 0 && json.Unmarshal(respBody, &parsed) == nil {
			apiErr.Response = parsed
		}
		return apiErr
	}

	if out == nil || len(respBody) == 0 {
		return nil
	}

	if err := json.Unmarshal(respBody, out); err != nil {
		return fmt.Errorf("decode response: %w", err)
	}
	return nil
}
