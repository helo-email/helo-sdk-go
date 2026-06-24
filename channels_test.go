package helo

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestChannels_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "GET"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/channels") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/channels")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &ChannelsListParams{
		Limit:        10,
		Offset:       10,
		Name:         "example",
		ChannelIds:   []string{"550e8400-e29b-41d4-a716-446655440000"},
		DeliveryType: "live",
	}
	result, err := client.Channels.List(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestChannels_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "POST"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/channels") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/channels")
		}
		if body, _ := io.ReadAll(r.Body); strings.Contains(string(body), ":{}") {
			t.Errorf("request body contains an empty object (omitempty not working): %s", body)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &CreateChannelRequest{
		Name:         "test-name",
		DeliveryType: DeliveryTypeLive,
	}
	result, err := client.Channels.Create(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestChannels_Retrieve(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "GET"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/channels/550e8400-e29b-41d4-a716-446655440000") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/channels/550e8400-e29b-41d4-a716-446655440000")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	result, err := client.Channels.Retrieve(context.Background(), "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestChannels_Update(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "PATCH"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/channels/550e8400-e29b-41d4-a716-446655440000") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/channels/550e8400-e29b-41d4-a716-446655440000")
		}
		if body, _ := io.ReadAll(r.Body); strings.Contains(string(body), ":{}") {
			t.Errorf("request body contains an empty object (omitempty not working): %s", body)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &UpdateChannelRequest{
		Name:         "test-name",
		DeliveryType: DeliveryTypeLive,
	}
	result, err := client.Channels.Update(context.Background(), "550e8400-e29b-41d4-a716-446655440000", params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestChannels_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "DELETE"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/channels/550e8400-e29b-41d4-a716-446655440000") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/channels/550e8400-e29b-41d4-a716-446655440000")
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	if err := client.Channels.Delete(context.Background(), "550e8400-e29b-41d4-a716-446655440000"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
