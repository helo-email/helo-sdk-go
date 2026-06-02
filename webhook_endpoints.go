package helo

import (
	"context"
	"fmt"
)

// WebhookEndpointsService exposes operations on the WebhookEndpoints resource.
type WebhookEndpointsService struct {
	client *Client
}

// List List all webhook endpoints
func (s *WebhookEndpointsService) List(ctx context.Context, params *WebhookEndpointsListParams) (*PaginationResultOfWebhookEndpointResponse, error) {
	out := new(PaginationResultOfWebhookEndpointResponse)
	if err := s.client.request(ctx, "GET", "/webhook-endpoints", out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// Create Create a webhook endpoint
func (s *WebhookEndpointsService) Create(ctx context.Context, params *CreateWebhookEndpointRequest) (*WebhookEndpointResponse, error) {
	out := new(WebhookEndpointResponse)
	if err := s.client.request(ctx, "POST", "/webhook-endpoints", out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// Retrieve Retrieve a webhook endpoint
func (s *WebhookEndpointsService) Retrieve(ctx context.Context, id string) (*WebhookEndpointResponse, error) {
	out := new(WebhookEndpointResponse)
	if err := s.client.request(ctx, "GET", fmt.Sprintf("/webhook-endpoints/%v", id), out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update Update a webhook endpoint
func (s *WebhookEndpointsService) Update(ctx context.Context, id string, params *UpdateWebhookEndpointRequest) (*WebhookEndpointResponse, error) {
	out := new(WebhookEndpointResponse)
	if err := s.client.request(ctx, "PATCH", fmt.Sprintf("/webhook-endpoints/%v", id), out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete Delete a webhook endpoint
func (s *WebhookEndpointsService) Delete(ctx context.Context, id string) error {
	if err := s.client.request(ctx, "DELETE", fmt.Sprintf("/webhook-endpoints/%v", id), nil); err != nil {
		return err
	}
	return nil
}

// RegenerateSigningKey Regenerate webhook signing key
func (s *WebhookEndpointsService) RegenerateSigningKey(ctx context.Context, id string) (*WebhookEndpointResponse, error) {
	out := new(WebhookEndpointResponse)
	if err := s.client.request(ctx, "POST", fmt.Sprintf("/webhook-endpoints/%v/regenerate-signing-key", id), out); err != nil {
		return nil, err
	}
	return out, nil
}

// WebhookEndpointsListParams are the query parameters for List.
type WebhookEndpointsListParams struct {
	Limit      int      `json:"limit,omitempty"`
	Offset     int      `json:"offset,omitempty"`
	ChannelIds []string `json:"channelIds,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *WebhookEndpointsListParams) toQuery() map[string]any {
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
	if len(p.ChannelIds) > 0 {
		q["channelIds"] = p.ChannelIds
	}
	return q
}
