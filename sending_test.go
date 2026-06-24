package helo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSending_Transactional(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "POST"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/send/transactional") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/send/transactional")
		}
		if got, want := r.Header.Get("X-Helo-Channel-Id"), "550e8400-e29b-41d4-a716-446655440000"; got != want {
			t.Errorf("header X-Helo-Channel-Id = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("X-Helo-Idempotency-Key"), "example"; got != want {
			t.Errorf("header X-Helo-Idempotency-Key = %q, want %q", got, want)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &SendMessageRequest{
		From:    MailAddress{Email: "test@example.com", Name: "test-name"},
		To:      []MailAddress{{Email: "test@example.com", Name: "test-name"}},
		Subject: "test-subject",
		Html:    "test-html",
		Text:    "test-text",
		Tags:    []string{"example1", "example2"},
	}
	opts := &SendingTransactionalOptions{
		ChannelID:      "550e8400-e29b-41d4-a716-446655440000",
		IdempotencyKey: "example",
	}
	result, err := client.Sending.Transactional(context.Background(), params, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestSending_TransactionalBatch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "POST"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/send/transactional/batch") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/send/transactional/batch")
		}
		if got, want := r.Header.Get("X-Helo-Channel-Id"), "550e8400-e29b-41d4-a716-446655440000"; got != want {
			t.Errorf("header X-Helo-Channel-Id = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("X-Helo-Idempotency-Key"), "example"; got != want {
			t.Errorf("header X-Helo-Idempotency-Key = %q, want %q", got, want)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &SendMessageBatchRequest{
		Requests: []SendMessageRequest{{From: MailAddress{Email: "test@example.com", Name: "test-name"}, To: []MailAddress{{Email: "test@example.com", Name: "test-name"}}, Subject: "test-subject", Html: "test-html", Text: "test-text", Tags: []string{"example1", "example2"}}},
	}
	opts := &SendingTransactionalBatchOptions{
		ChannelID:      "550e8400-e29b-41d4-a716-446655440000",
		IdempotencyKey: "example",
	}
	result, err := client.Sending.TransactionalBatch(context.Background(), params, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestSending_Broadcast(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "POST"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/send/broadcast") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/send/broadcast")
		}
		if got, want := r.Header.Get("X-Helo-Channel-Id"), "550e8400-e29b-41d4-a716-446655440000"; got != want {
			t.Errorf("header X-Helo-Channel-Id = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("X-Helo-Idempotency-Key"), "example"; got != want {
			t.Errorf("header X-Helo-Idempotency-Key = %q, want %q", got, want)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &SendBroadcastRequest{
		From:     MailAddress{Email: "test@example.com", Name: "test-name"},
		Template: SendBroadcastRequestTemplate{Subject: "test-subject", Html: "test-html", Text: "test-text", InlineStyles: true},
		Tags:     []string{"example1", "example2"},
		Messages: []SendBroadcastRequestMessage{{To: []MailAddress{{Email: "test@example.com", Name: "test-name"}}, Tags: []string{"example1", "example2"}}},
	}
	opts := &SendingBroadcastOptions{
		ChannelID:      "550e8400-e29b-41d4-a716-446655440000",
		IdempotencyKey: "example",
	}
	result, err := client.Sending.Broadcast(context.Background(), params, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestSending_BroadcastMessage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "POST"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/send/broadcast/message") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/send/broadcast/message")
		}
		if got, want := r.Header.Get("X-Helo-Channel-Id"), "550e8400-e29b-41d4-a716-446655440000"; got != want {
			t.Errorf("header X-Helo-Channel-Id = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("X-Helo-Idempotency-Key"), "example"; got != want {
			t.Errorf("header X-Helo-Idempotency-Key = %q, want %q", got, want)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &SendMessageRequest{
		From:    MailAddress{Email: "test@example.com", Name: "test-name"},
		To:      []MailAddress{{Email: "test@example.com", Name: "test-name"}},
		Subject: "test-subject",
		Html:    "test-html",
		Text:    "test-text",
		Tags:    []string{"example1", "example2"},
	}
	opts := &SendingBroadcastMessageOptions{
		ChannelID:      "550e8400-e29b-41d4-a716-446655440000",
		IdempotencyKey: "example",
	}
	result, err := client.Sending.BroadcastMessage(context.Background(), params, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}
