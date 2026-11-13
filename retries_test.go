package helo

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

// The retry policy every Helo SDK shares.

type reply struct {
	status     int
	retryAfter string
}

// stub serves the replies in order and records each request it receives.
func stub(t *testing.T, replies ...reply) (*Client, *[]*http.Request) {
	t.Helper()
	var requests []*http.Request
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests = append(requests, r)
		next := replies[0]
		if len(replies) > 1 {
			replies = replies[1:]
		}
		if next.retryAfter != "" {
			w.Header().Set("Retry-After", next.retryAfter)
		}
		w.WriteHeader(next.status)
	}))
	t.Cleanup(server.Close)
	return &Client{BaseURL: server.URL, APIKey: "test-key", HTTPClient: server.Client(), MaxRetries: DefaultMaxRetries}, &requests
}

func TestRetriesAnIdempotentRequest(t *testing.T) {
	client, requests := stub(t, reply{status: 503, retryAfter: "0"}, reply{status: 204})

	if err := client.request(context.Background(), "GET", "/things", nil); err != nil {
		t.Fatal(err)
	}
	if len(*requests) != 2 {
		t.Fatalf("got %d requests, want 2", len(*requests))
	}
}

func TestRetriesAPostCarryingAnIdempotencyKey(t *testing.T) {
	client, requests := stub(t, reply{status: 502, retryAfter: "0"}, reply{status: 204})

	err := client.request(context.Background(), "POST", "/things", nil,
		withBody(map[string]any{}), withHeaders(map[string]string{"X-Example-Idempotency-Key": "key-1"}))
	if err != nil {
		t.Fatal(err)
	}
	if len(*requests) != 2 || (*requests)[1].Header.Get("X-Example-Idempotency-Key") != "key-1" {
		t.Fatalf("want the same key sent twice, got %d requests", len(*requests))
	}
}

// Without a key, the server may have acted on the first attempt before failing.
func TestDoesNotRetryAPostWithoutAnIdempotencyKey(t *testing.T) {
	client, requests := stub(t, reply{status: 503, retryAfter: "0"})

	err := client.request(context.Background(), "POST", "/things", nil, withBody(map[string]any{}))
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != 503 {
		t.Fatalf("want a 503 APIError, got %v", err)
	}
	if len(*requests) != 1 {
		t.Fatalf("got %d requests, want 1", len(*requests))
	}
}

func TestDoesNotRetryAPostWithoutAKeyAfterANetworkError(t *testing.T) {
	attempts := 0
	client := &Client{BaseURL: "http://localhost:8002", MaxRetries: DefaultMaxRetries, HTTPClient: &http.Client{
		Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			attempts++
			return nil, errors.New("connection refused")
		}),
	}}

	if err := client.request(context.Background(), "POST", "/things", nil, withBody(map[string]any{})); err == nil {
		t.Fatal("want an error")
	}
	if attempts != 1 {
		t.Fatalf("got %d attempts, want 1", attempts)
	}
}

func TestRetriesANetworkErrorOnAnIdempotentRequest(t *testing.T) {
	attempts := 0
	client := &Client{BaseURL: "http://localhost:8002", MaxRetries: DefaultMaxRetries, HTTPClient: &http.Client{
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			attempts++
			if attempts == 1 {
				return nil, errors.New("connection reset")
			}
			return &http.Response{StatusCode: 204, Body: http.NoBody, Header: http.Header{}, Request: r}, nil
		}),
	}}

	if err := client.request(context.Background(), "GET", "/things", nil); err != nil {
		t.Fatal(err)
	}
	if attempts != 2 {
		t.Fatalf("got %d attempts, want 2", attempts)
	}
}

// A 429 means the server refused the request without acting on it.
func TestRetriesARateLimitedPostEvenWithoutAKey(t *testing.T) {
	client, requests := stub(t, reply{status: 429, retryAfter: "0"}, reply{status: 204})

	if err := client.request(context.Background(), "POST", "/things", nil, withBody(map[string]any{})); err != nil {
		t.Fatal(err)
	}
	if len(*requests) != 2 {
		t.Fatalf("got %d requests, want 2", len(*requests))
	}
}

func TestDoesNotRetryEarlyWhenRetryAfterIsLongerThanItWaits(t *testing.T) {
	client, requests := stub(t, reply{status: 429, retryAfter: "60"})

	var apiErr *APIError
	if err := client.request(context.Background(), "GET", "/things", nil); !errors.As(err, &apiErr) || apiErr.StatusCode != 429 {
		t.Fatalf("want a 429 APIError, got %v", err)
	}
	if len(*requests) != 1 {
		t.Fatalf("got %d requests, want 1", len(*requests))
	}
}

func TestGivesUpAfterMaxRetries(t *testing.T) {
	client, requests := stub(t, reply{status: 500, retryAfter: "0"})
	client.MaxRetries = 1

	if err := client.request(context.Background(), "GET", "/things", nil); err == nil {
		t.Fatal("want an error")
	}
	if len(*requests) != 2 {
		t.Fatalf("got %d requests, want 2", len(*requests))
	}
}

func TestMaxRetriesZeroDisablesRetries(t *testing.T) {
	client, requests := stub(t, reply{status: 503, retryAfter: "0"})
	client.MaxRetries = 0

	if err := client.request(context.Background(), "GET", "/things", nil); err == nil {
		t.Fatal("want an error")
	}
	if len(*requests) != 1 {
		t.Fatalf("got %d requests, want 1", len(*requests))
	}
}

func TestNewClientRetriesByDefault(t *testing.T) {
	if got := NewHelo("key").client.MaxRetries; got != DefaultMaxRetries {
		t.Fatalf("MaxRetries = %d, want %d", got, DefaultMaxRetries)
	}
	if got := NewHelo("key", WithMaxRetries(0)).client.MaxRetries; got != 0 {
		t.Fatalf("MaxRetries = %d, want 0", got)
	}
}

// A send whose response is lost must not be sent twice: without a key from the
// caller, the SDK generates one per call, which the retry then reuses.
func TestGeneratesAnIdempotencyKeyWhenNotGiven(t *testing.T) {
	headers := (*SendingSendTransactionalOptions)(nil).toHeaders()

	if !regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`).MatchString(headers["X-Helo-Idempotency-Key"]) {
		t.Fatalf("want a generated UUID, got %q", headers["X-Helo-Idempotency-Key"])
	}
	if given := (&SendingSendTransactionalOptions{IdempotencyKey: "mine"}).toHeaders(); given["X-Helo-Idempotency-Key"] != "mine" {
		t.Fatalf("want the caller's key, got %q", given["X-Helo-Idempotency-Key"])
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestAbandonsAnAttemptThatTakesTooLongAndRetriesIt(t *testing.T) {
	attempts := 0
	client := &Client{BaseURL: "http://localhost:8002", MaxRetries: 1, HTTPClient: &http.Client{
		Timeout: 50 * time.Millisecond,
		Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			attempts++
			if attempts == 1 {
				<-r.Context().Done() // hangs until the client's Timeout cancels it
				return nil, r.Context().Err()
			}
			return &http.Response{StatusCode: 204, Body: http.NoBody, Header: http.Header{}, Request: r}, nil
		}),
	}}

	if err := client.request(context.Background(), "GET", "/things", nil); err != nil {
		t.Fatalf("want success after a retry, got %v", err)
	}
	if attempts != 2 {
		t.Fatalf("got %d attempts, want 2", attempts)
	}
}

func TestTimesOutByDefault(t *testing.T) {
	if got := NewHelo("test-key").client.HTTPClient.Timeout; got != DefaultTimeout {
		t.Fatalf("got timeout %v, want %v", got, DefaultTimeout)
	}
}
