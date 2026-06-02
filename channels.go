package helo

import (
	"context"
	"fmt"
)

// ChannelsService exposes operations on the Channels resource.
type ChannelsService struct {
	client *Client
}

// List List all channels
func (s *ChannelsService) List(ctx context.Context, params *ChannelsListParams) (*PaginationResultOfChannelBasicResponse, error) {
	out := new(PaginationResultOfChannelBasicResponse)
	if err := s.client.request(ctx, "GET", "/channels", out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// Create Create a channel
func (s *ChannelsService) Create(ctx context.Context, params *CreateChannelRequest) (*ChannelDetailsResponse, error) {
	out := new(ChannelDetailsResponse)
	if err := s.client.request(ctx, "POST", "/channels", out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// Retrieve Retrieve a channel
func (s *ChannelsService) Retrieve(ctx context.Context, id string) (*ChannelDetailsResponse, error) {
	out := new(ChannelDetailsResponse)
	if err := s.client.request(ctx, "GET", fmt.Sprintf("/channels/%v", id), out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update Update a channel
func (s *ChannelsService) Update(ctx context.Context, id string, params *UpdateChannelRequest) (*ChannelDetailsResponse, error) {
	out := new(ChannelDetailsResponse)
	if err := s.client.request(ctx, "PATCH", fmt.Sprintf("/channels/%v", id), out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete Delete a channel
func (s *ChannelsService) Delete(ctx context.Context, id string) error {
	if err := s.client.request(ctx, "DELETE", fmt.Sprintf("/channels/%v", id), nil); err != nil {
		return err
	}
	return nil
}

// ChannelsListParams are the query parameters for List.
type ChannelsListParams struct {
	Limit        int      `json:"limit,omitempty"`
	Offset       int      `json:"offset,omitempty"`
	Name         string   `json:"name,omitempty"`
	ChannelIds   []string `json:"channelIds,omitempty"`
	DeliveryType string   `json:"deliveryType,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *ChannelsListParams) toQuery() map[string]any {
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
	if p.DeliveryType != "" {
		q["deliveryType"] = p.DeliveryType
	}
	return q
}
