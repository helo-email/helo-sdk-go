package helo

import (
	"context"
	"fmt"
)

// WebhooksService exposes operations on the Webhooks resource.
type WebhooksService struct {
	client *Client
}

// List List all webhooks
func (s *WebhooksService) List(ctx context.Context, params *WebhooksListParams) (*PaginationResultOfWebhookResponse, error) {
	out := new(PaginationResultOfWebhookResponse)
	if err := s.client.request(ctx, "GET", "/webhooks", out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// Create Create a webhook
func (s *WebhooksService) Create(ctx context.Context, params *CreateWebhookRequest) (*WebhookResponse, error) {
	out := new(WebhookResponse)
	if err := s.client.request(ctx, "POST", "/webhooks", out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// Retrieve Retrieve a webhook
func (s *WebhooksService) Retrieve(ctx context.Context, id string) (*WebhookResponse, error) {
	out := new(WebhookResponse)
	if err := s.client.request(ctx, "GET", fmt.Sprintf("/webhooks/%v", id), out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update Update a webhook
func (s *WebhooksService) Update(ctx context.Context, id string, params *UpdateWebhookRequest) (*WebhookResponse, error) {
	out := new(WebhookResponse)
	if err := s.client.request(ctx, "PATCH", fmt.Sprintf("/webhooks/%v", id), out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete Delete a webhook
func (s *WebhooksService) Delete(ctx context.Context, id string) error {
	if err := s.client.request(ctx, "DELETE", fmt.Sprintf("/webhooks/%v", id), nil); err != nil {
		return err
	}
	return nil
}

// RegenerateSigningKey Regenerate webhook signing key
func (s *WebhooksService) RegenerateSigningKey(ctx context.Context, id string) (*WebhookResponse, error) {
	out := new(WebhookResponse)
	if err := s.client.request(ctx, "POST", fmt.Sprintf("/webhooks/%v/regenerate-signing-key", id), out); err != nil {
		return nil, err
	}
	return out, nil
}

// WebhooksListParams are the query parameters for List.
type WebhooksListParams struct {
	Limit      int      `json:"limit,omitempty"`
	Offset     int      `json:"offset,omitempty"`
	ChannelIds []string `json:"channelIds,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *WebhooksListParams) toQuery() map[string]any {
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
