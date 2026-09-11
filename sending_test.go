package helo

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSending_SendTransactional(t *testing.T) {
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
		if body, _ := io.ReadAll(r.Body); strings.Contains(string(body), ":{}") {
			t.Errorf("request body contains an empty object (omitempty not working): %s", body)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &SendMessageRequest{
		From:    MailAddress{Email: "from@yourdomain.com", Name: "From name"},
		To:      []MailAddress{{Email: "to@example.com", Name: "To name"}},
		Subject: "Hello from Helo",
		Html:    "<html><body><h1>Hi there, new friend.</h1><p>This is a test message, delivered with <3 by Helo. </p></body></html>",
		Text:    "This is a test message, delivered with <3 by Helo.",
		Tags:    []string{"welcome", "onboarding"},
	}
	opts := &SendingSendTransactionalOptions{
		ChannelID:      "550e8400-e29b-41d4-a716-446655440000",
		IdempotencyKey: "example",
	}
	result, err := client.Sending.SendTransactional(context.Background(), params, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestSending_SendTransactionalBatch(t *testing.T) {
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
		if body, _ := io.ReadAll(r.Body); strings.Contains(string(body), ":{}") {
			t.Errorf("request body contains an empty object (omitempty not working): %s", body)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &SendMessageBatchRequest{
		Requests: []SendMessageRequest{{From: MailAddress{Email: "from@yourdomain.com", Name: "From name"}, To: []MailAddress{{Email: "to@example.com", Name: "To name"}}, Subject: "Hello from Helo", Html: "<html><body><h1>Hi there, new friend.</h1><p>This is a test message, delivered with <3 by Helo. </p></body></html>", Text: "This is a test message, delivered with <3 by Helo.", Tags: []string{"welcome", "onboarding"}}},
	}
	opts := &SendingSendTransactionalBatchOptions{
		ChannelID:      "550e8400-e29b-41d4-a716-446655440000",
		IdempotencyKey: "example",
	}
	result, err := client.Sending.SendTransactionalBatch(context.Background(), params, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestSending_SendBroadcast(t *testing.T) {
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
		if body, _ := io.ReadAll(r.Body); strings.Contains(string(body), ":{}") {
			t.Errorf("request body contains an empty object (omitempty not working): %s", body)
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
	opts := &SendingSendBroadcastOptions{
		ChannelID:      "550e8400-e29b-41d4-a716-446655440000",
		IdempotencyKey: "example",
	}
	result, err := client.Sending.SendBroadcast(context.Background(), params, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestSending_SendBroadcastMessage(t *testing.T) {
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
		if body, _ := io.ReadAll(r.Body); strings.Contains(string(body), ":{}") {
			t.Errorf("request body contains an empty object (omitempty not working): %s", body)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &SendMessageRequest{
		From:    MailAddress{Email: "from@yourdomain.com", Name: "From name"},
		To:      []MailAddress{{Email: "to@example.com", Name: "To name"}},
		Subject: "Hello from Helo",
		Html:    "<html><body><h1>Hi there, new friend.</h1><p>This is a test message, delivered with <3 by Helo. </p></body></html>",
		Text:    "This is a test message, delivered with <3 by Helo.",
		Tags:    []string{"welcome", "onboarding"},
	}
	opts := &SendingSendBroadcastMessageOptions{
		ChannelID:      "550e8400-e29b-41d4-a716-446655440000",
		IdempotencyKey: "example",
	}
	result, err := client.Sending.SendBroadcastMessage(context.Background(), params, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}
