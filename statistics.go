package helo

import (
	"context"
	"time"
)

// StatisticsService exposes operations on the Statistics resource.
type StatisticsService struct {
	client *Client
}

// RetrieveHourly Retrieve hourly statistics
func (s *StatisticsService) RetrieveHourly(ctx context.Context, params *StatisticsRetrieveHourlyParams) (*StatisticsHourlyResponse, error) {
	out := new(StatisticsHourlyResponse)
	if err := s.client.request(ctx, "GET", "/statistics/hourly", out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// RetrieveDaily Retrieve daily statistics
func (s *StatisticsService) RetrieveDaily(ctx context.Context, params *StatisticsRetrieveDailyParams) (*StatisticsDailyResponse, error) {
	out := new(StatisticsDailyResponse)
	if err := s.client.request(ctx, "GET", "/statistics/daily", out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// RetrieveTotals Retrieve all time statistics
func (s *StatisticsService) RetrieveTotals(ctx context.Context, params *StatisticsRetrieveTotalsParams) (*StatisticsTotalsResponse, error) {
	out := new(StatisticsTotalsResponse)
	if err := s.client.request(ctx, "GET", "/statistics/totals", out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// StatisticsRetrieveHourlyParams are the query parameters for RetrieveHourly.
type StatisticsRetrieveHourlyParams struct {
	ChannelID string    `json:"channelId,omitempty"`
	From      time.Time `json:"from,omitempty"`
	To        time.Time `json:"to,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *StatisticsRetrieveHourlyParams) toQuery() map[string]any {
	if p == nil {
		return nil
	}
	q := map[string]any{}
	if p.ChannelID != "" {
		q["channelId"] = p.ChannelID
	}
	q["from"] = p.From.Format(time.RFC3339)
	q["to"] = p.To.Format(time.RFC3339)
	if len(p.Tags) > 0 {
		q["tags"] = p.Tags
	}
	return q
}

// StatisticsRetrieveDailyParams are the query parameters for RetrieveDaily.
type StatisticsRetrieveDailyParams struct {
	ChannelID string   `json:"channelId,omitempty"`
	From      string   `json:"from,omitempty"`
	To        string   `json:"to,omitempty"`
	Tags      []string `json:"tags,omitempty"`
	Timezone  string   `json:"timezone,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *StatisticsRetrieveDailyParams) toQuery() map[string]any {
	if p == nil {
		return nil
	}
	q := map[string]any{}
	if p.ChannelID != "" {
		q["channelId"] = p.ChannelID
	}
	q["from"] = p.From
	q["to"] = p.To
	if len(p.Tags) > 0 {
		q["tags"] = p.Tags
	}
	q["timezone"] = p.Timezone
	return q
}

// StatisticsRetrieveTotalsParams are the query parameters for RetrieveTotals.
type StatisticsRetrieveTotalsParams struct {
	ChannelID string    `json:"channelId,omitempty"`
	From      time.Time `json:"from,omitempty"`
	To        time.Time `json:"to,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *StatisticsRetrieveTotalsParams) toQuery() map[string]any {
	if p == nil {
		return nil
	}
	q := map[string]any{}
	if p.ChannelID != "" {
		q["channelId"] = p.ChannelID
	}
	q["from"] = p.From.Format(time.RFC3339)
	q["to"] = p.To.Format(time.RFC3339)
	if len(p.Tags) > 0 {
		q["tags"] = p.Tags
	}
	return q
}
