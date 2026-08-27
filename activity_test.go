package helo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestActivity_ListEvents(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "GET"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/activity/events") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/activity/events")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &ActivityListEventsParams{
		ChannelID:  "550e8400-e29b-41d4-a716-446655440000",
		MessageID:  "550e8400-e29b-41d4-a716-446655440000",
		After:      10,
		StartDate:  time.Now(),
		EndDate:    time.Now(),
		Limit:      10,
		Recipient:  "example",
		Subject:    "example",
		Tags:       []string{"example1", "example2"},
		MailType:   "transactional",
		EventTypes: []EventType{EventTypeAccepted, EventTypeProcessed},
	}
	result, err := client.Activity.ListEvents(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestActivity_ListMessages(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "GET"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/activity/messages") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/activity/messages")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &ActivityListMessagesParams{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		After:     10,
		StartDate: time.Now(),
		EndDate:   time.Now(),
		Limit:     10,
		Recipient: "example",
		Subject:   "example",
		Tags:      []string{"example1", "example2"},
		MailType:  "transactional",
		Status:    MessageStatusQueued,
	}
	result, err := client.Activity.ListMessages(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestActivity_RetrieveMessage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "GET"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/activity/messages/550e8400-e29b-41d4-a716-446655440000") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/activity/messages/550e8400-e29b-41d4-a716-446655440000")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	result, err := client.Activity.RetrieveMessage(context.Background(), "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}
