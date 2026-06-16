package helo

import (
	"context"
	"fmt"
)

// BroadcastsService exposes operations on the Broadcasts resource.
type BroadcastsService struct {
	client *Client
}

// List List broadcasts
func (s *BroadcastsService) List(ctx context.Context, params *BroadcastsListParams) (*PaginatedResponseOfBroadcast, error) {
	out := new(PaginatedResponseOfBroadcast)
	if err := s.client.request(ctx, "GET", "/broadcasts", out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// Retrieve Retrieve a broadcast
func (s *BroadcastsService) Retrieve(ctx context.Context, id string) (*BroadcastDetailsResponse, error) {
	out := new(BroadcastDetailsResponse)
	if err := s.client.request(ctx, "GET", fmt.Sprintf("/broadcasts/%v", id), out); err != nil {
		return nil, err
	}
	return out, nil
}

// ListFailures List broadcast failures
func (s *BroadcastsService) ListFailures(ctx context.Context, id string, params *BroadcastsListFailuresParams) (*PaginatedResponseOfBroadcastFailure, error) {
	out := new(PaginatedResponseOfBroadcastFailure)
	if err := s.client.request(ctx, "GET", fmt.Sprintf("/broadcasts/%v/failures", id), out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// ListSuppressions List broadcast suppressions
func (s *BroadcastsService) ListSuppressions(ctx context.Context, id string, params *BroadcastsListSuppressionsParams) (*PaginatedResponseOfBroadcastSuppression, error) {
	out := new(PaginatedResponseOfBroadcastSuppression)
	if err := s.client.request(ctx, "GET", fmt.Sprintf("/broadcasts/%v/suppressions", id), out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcastsListParams are the query parameters for List.
type BroadcastsListParams struct {
	ChannelID string          `json:"channelId,omitempty"`
	Status    BroadcastStatus `json:"status,omitempty"`
	Subject   string          `json:"subject,omitempty"`
	Limit     int             `json:"limit,omitempty"`
	Offset    int             `json:"offset,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *BroadcastsListParams) toQuery() map[string]any {
	if p == nil {
		return nil
	}
	q := map[string]any{}
	q["channelId"] = p.ChannelID
	if p.Status != "" {
		q["status"] = p.Status
	}
	if p.Subject != "" {
		q["subject"] = p.Subject
	}
	if p.Limit != 0 {
		q["limit"] = p.Limit
	}
	if p.Offset != 0 {
		q["offset"] = p.Offset
	}
	return q
}

// BroadcastsListFailuresParams are the query parameters for ListFailures.
type BroadcastsListFailuresParams struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *BroadcastsListFailuresParams) toQuery() map[string]any {
	if p == nil {
		return nil
	}
	q := map[string]any{}
	if p.Limit != 0 {
		q["limit"] = p.Limit
	}
	if p.Offset != 0 {
		q["offset"] = p.Offset
	}
	return q
}

// BroadcastsListSuppressionsParams are the query parameters for ListSuppressions.
type BroadcastsListSuppressionsParams struct {
	Limit  int `json:"limit,omitempty"`
	Offset int `json:"offset,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *BroadcastsListSuppressionsParams) toQuery() map[string]any {
	if p == nil {
		return nil
	}
	q := map[string]any{}
	if p.Limit != 0 {
		q["limit"] = p.Limit
	}
	if p.Offset != 0 {
		q["offset"] = p.Offset
	}
	return q
}
