package helo

import (
	"context"
	"fmt"
)

// TemplatesService exposes operations on the Templates resource.
type TemplatesService struct {
	client *Client
}

// List List operation
func (s *TemplatesService) List(ctx context.Context, params *TemplatesListParams) (*PaginatedResponseOfTemplateSummaryResponse, error) {
	out := new(PaginatedResponseOfTemplateSummaryResponse)
	if err := s.client.request(ctx, "GET", "/templates", out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// Create Create operation
func (s *TemplatesService) Create(ctx context.Context, params *CreateTemplateRequest) (*TemplateResponse, error) {
	out := new(TemplateResponse)
	if err := s.client.request(ctx, "POST", "/templates", out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// RenderDraft RenderDraft operation
func (s *TemplatesService) RenderDraft(ctx context.Context, params *RenderDraftRequest) (*RenderTemplateResponse, error) {
	out := new(RenderTemplateResponse)
	if err := s.client.request(ctx, "POST", "/templates/render", out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// Retrieve Retrieve operation
func (s *TemplatesService) Retrieve(ctx context.Context, id string) (*TemplateResponse, error) {
	out := new(TemplateResponse)
	if err := s.client.request(ctx, "GET", fmt.Sprintf("/templates/%v", id), out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update Update operation
func (s *TemplatesService) Update(ctx context.Context, id string, params *UpdateTemplateRequest) (*TemplateResponse, error) {
	out := new(TemplateResponse)
	if err := s.client.request(ctx, "PATCH", fmt.Sprintf("/templates/%v", id), out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete Delete operation
func (s *TemplatesService) Delete(ctx context.Context, id string) error {
	if err := s.client.request(ctx, "DELETE", fmt.Sprintf("/templates/%v", id), nil); err != nil {
		return err
	}
	return nil
}

// Render Render operation
func (s *TemplatesService) Render(ctx context.Context, id string, params *RenderTemplateRequest) (*RenderTemplateResponse, error) {
	out := new(RenderTemplateResponse)
	if err := s.client.request(ctx, "POST", fmt.Sprintf("/templates/%v/render", id), out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// TemplatesListParams are the query parameters for List.
type TemplatesListParams struct {
	Limit      int      `json:"limit,omitempty"`
	Offset     int      `json:"offset,omitempty"`
	Name       string   `json:"name,omitempty"`
	ChannelIds []string `json:"channelIds,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *TemplatesListParams) toQuery() map[string]any {
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
	if p.Name != "" {
		q["name"] = p.Name
	}
	if len(p.ChannelIds) > 0 {
		q["channelIds"] = p.ChannelIds
	}
	return q
}
