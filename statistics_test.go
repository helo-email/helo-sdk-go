package helo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestStatistics_RetrieveHourly(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "GET"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/activity/statistics/hourly") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/activity/statistics/hourly")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &StatisticsRetrieveHourlyParams{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		From:      time.Now(),
		To:        time.Now(),
		Tags:      []string{"example1", "example2"},
	}
	result, err := client.Statistics.RetrieveHourly(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestStatistics_RetrieveDaily(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "GET"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/activity/statistics/daily") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/activity/statistics/daily")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &StatisticsRetrieveDailyParams{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		From:      "2024-01-01",
		To:        "2024-01-01",
		Tags:      []string{"example1", "example2"},
		Timezone:  "America/New_York",
	}
	result, err := client.Statistics.RetrieveDaily(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}

func TestStatistics_RetrieveTotals(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got, want := r.Method, "GET"; got != want {
			t.Errorf("method = %q, want %q", got, want)
		}
		if got, want := r.Header.Get("Authorization"), "Bearer test-token-123"; got != want {
			t.Errorf("authorization = %q, want %q", got, want)
		}
		if !strings.HasPrefix(r.URL.Path, "/activity/statistics/totals") {
			t.Errorf("path = %q, want prefix %q", r.URL.Path, "/activity/statistics/totals")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := NewHelo("test-token-123", WithBaseURL(server.URL))

	params := &StatisticsRetrieveTotalsParams{
		ChannelID: "550e8400-e29b-41d4-a716-446655440000",
		From:      time.Now(),
		To:        time.Now(),
		Tags:      []string{"example1", "example2"},
	}
	result, err := client.Statistics.RetrieveTotals(context.Background(), params)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
}
