package helo

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTemplates_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "GET"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/templates") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/templates")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &TemplatesListParams{
		Limit:      10,
		Offset:     10,
		Name:       "example",
		ChannelIds: []string{"550e8400-e29b-41d4-a716-446655440000"},
	}
	result, err := client.Templates.List(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestTemplates_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "POST"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/templates") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/templates")
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

	params := &CreateTemplateRequest{
		Name:       "test-name",
		Subject:    "test-subject",
		Html:       "test-html",
		Mjml:       "test-mjml",
		Text:       "test-text",
		ChannelIds: []string{"550e8400-e29b-41d4-a716-446655440000"},
	}
	result, err := client.Templates.Create(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestTemplates_RenderDraft(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "POST"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/templates/render") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/templates/render")
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

	params := &RenderDraftRequest{
		Subject:      "test-subject",
		Html:         "test-html",
		Mjml:         "test-mjml",
		Text:         "test-text",
		InlineStyles: Bool(true),
	}
	result, err := client.Templates.RenderDraft(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestTemplates_Retrieve(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "GET"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/templates/550e8400-e29b-41d4-a716-446655440000") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/templates/550e8400-e29b-41d4-a716-446655440000")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	result, err := client.Templates.Retrieve(context.Background(), "550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestTemplates_Update(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "PATCH"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/templates/550e8400-e29b-41d4-a716-446655440000") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/templates/550e8400-e29b-41d4-a716-446655440000")
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

	params := &UpdateTemplateRequest{
		Name:       "test-name",
		Subject:    "test-subject",
		Html:       "test-html",
		Mjml:       "test-mjml",
		Text:       "test-text",
		ChannelIds: []string{"550e8400-e29b-41d4-a716-446655440000"},
		Version:    123,
	}
	result, err := client.Templates.Update(context.Background(), "550e8400-e29b-41d4-a716-446655440000", params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestTemplates_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "DELETE"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/templates/550e8400-e29b-41d4-a716-446655440000") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/templates/550e8400-e29b-41d4-a716-446655440000")
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	if err := client.Templates.Delete(context.Background(), "550e8400-e29b-41d4-a716-446655440000"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTemplates_Render(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "POST"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/templates/550e8400-e29b-41d4-a716-446655440000/render") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/templates/550e8400-e29b-41d4-a716-446655440000/render")
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

	params := &RenderTemplateRequest{
		InlineStyles: Bool(true),
	}
	result, err := client.Templates.Render(context.Background(), "550e8400-e29b-41d4-a716-446655440000", params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}
