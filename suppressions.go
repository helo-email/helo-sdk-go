package helo

import (
	"context"
)

// SuppressionsService exposes operations on the Suppressions resource.
type SuppressionsService struct {
	client *Client
}

// List List suppressions
func (s *SuppressionsService) List(ctx context.Context, params *SuppressionsListParams) (*PaginatedResponseOfSuppressionResponse, error) {
	out := new(PaginatedResponseOfSuppressionResponse)
	if err := s.client.request(ctx, "GET", "/suppressions", out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// Create Create suppressions
func (s *SuppressionsService) Create(ctx context.Context, params *CreateSuppressionsRequest) (*CreateSuppressionsResponse, error) {
	out := new(CreateSuppressionsResponse)
	if err := s.client.request(ctx, "POST", "/suppressions", out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// Remove Remove suppressions
func (s *SuppressionsService) Remove(ctx context.Context, params *RemoveSuppressionsRequest) (*RemoveSuppressionsResponse, error) {
	out := new(RemoveSuppressionsResponse)
	if err := s.client.request(ctx, "POST", "/suppressions/remove", out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// SuppressionsListParams are the query parameters for List.
type SuppressionsListParams struct {
	ChannelID string            `json:"channelId,omitempty"`
	MailType  MailType          `json:"mailType,omitempty"`
	Reason    SuppressionReason `json:"reason,omitempty"`
	Email     string            `json:"email,omitempty"`
	Limit     int               `json:"limit,omitempty"`
	Offset    int               `json:"offset,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *SuppressionsListParams) toQuery() map[string]any {
	if p == nil {
		return nil
	}
	q := map[string]any{}
	q["channelId"] = p.ChannelID
	q["mailType"] = p.MailType
	if p.Reason != "" {
		q["reason"] = p.Reason
	}
	if p.Email != "" {
		q["email"] = p.Email
	}
	if p.Limit != 0 {
		q["limit"] = p.Limit
	}
	if p.Offset != 0 {
		q["offset"] = p.Offset
	}
	return q
}
