package helo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWebhookEndpoints_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "GET"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/webhook-endpoints") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/webhook-endpoints")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &WebhookEndpointsListParams{
		Limit:      10,
		Offset:     10,
		ChannelIds: []string{"550e8400-e29b-41d4-a716-446655440000"},
	}
	result, err := client.WebhookEndpoints.List(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestWebhookEndpoints_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "POST"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/webhook-endpoints") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/webhook-endpoints")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &CreateWebhookEndpointRequest{
		URL:       "test-url",
		Events:    []WebhookEvent{WebhookEventAccepted, WebhookEventProcessed},
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		Enabled:   true,
	}
	result, err := client.WebhookEndpoints.Create(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestWebhookEndpoints_Retrieve(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "GET"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/webhook-endpoints/550e8400-e29b-41d4-a716-446655440000") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/webhook-endpoints/550e8400-e29b-41d4-a716-446655440000")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	result, err := client.WebhookEndpoints.Retrieve(context.Background(), "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestWebhookEndpoints_Update(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "PATCH"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/webhook-endpoints/550e8400-e29b-41d4-a716-446655440000") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/webhook-endpoints/550e8400-e29b-41d4-a716-446655440000")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &UpdateWebhookEndpointRequest{
		URL:       "test-url",
		Events:    []WebhookEvent{WebhookEventAccepted, WebhookEventProcessed},
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		Enabled:   true,
	}
	result, err := client.WebhookEndpoints.Update(context.Background(), "550e8400-e29b-41d4-a716-446655440000", params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestWebhookEndpoints_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "DELETE"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/webhook-endpoints/550e8400-e29b-41d4-a716-446655440000") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/webhook-endpoints/550e8400-e29b-41d4-a716-446655440000")
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	if err := client.WebhookEndpoints.Delete(context.Background(), "550e8400-e29b-41d4-a716-446655440000"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestWebhookEndpoints_RegenerateSigningKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "POST"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/webhook-endpoints/550e8400-e29b-41d4-a716-446655440000/regenerate-signing-key") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/webhook-endpoints/550e8400-e29b-41d4-a716-446655440000/regenerate-signing-key")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	result, err := client.WebhookEndpoints.RegenerateSigningKey(context.Background(), "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}
