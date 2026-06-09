package helo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Client is the HTTP client used by all resource services.
type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// requestOptions carries optional body and query parameters for a single call.
type requestOptions struct {
	body  any
	query map[string]any
}

// requestOption mutates a requestOptions value.
type requestOption func(*requestOptions)

func withBody(body any) requestOption {
	return func(o *requestOptions) { o.body = body }
}

func withQuery(query map[string]any) requestOption {
	return func(o *requestOptions) { o.query = query }
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
				q.Set(key, fmt.Sprintf("%v", v))
			}
		}
		endpoint.RawQuery = q.Encode()
	}

	var bodyReader io.Reader
	if options.body != nil {
		buf, err := json.Marshal(options.body)
		if err != nil {
			return fmt.Errorf("marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(buf)
	}

	req, err := http.NewRequestWithContext(ctx, strings.ToUpper(method), endpoint.String(), bodyReader)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	if c.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}
	if options.body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return &APIError{Message: fmt.Sprintf("connection failed: %v", err), Cause: err}
	}
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
