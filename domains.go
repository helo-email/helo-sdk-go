package helo

import (
	"context"
	"fmt"
)

// DomainsService exposes operations on the Domains resource.
type DomainsService struct {
	client *Client
}

// List List all domains
func (s *DomainsService) List(ctx context.Context, params *DomainsListParams) (*PaginatedResponseOfDomainResponse, error) {
	out := new(PaginatedResponseOfDomainResponse)
	if err := s.client.request(ctx, "GET", "/domains", out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// Create Create a domain
func (s *DomainsService) Create(ctx context.Context, params *CreateDomainRequest) (*DomainWithDnsResponse, error) {
	out := new(DomainWithDnsResponse)
	if err := s.client.request(ctx, "POST", "/domains", out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// Retrieve Retrieve a domain
func (s *DomainsService) Retrieve(ctx context.Context, id string) (*DomainWithDnsResponse, error) {
	out := new(DomainWithDnsResponse)
	if err := s.client.request(ctx, "GET", fmt.Sprintf("/domains/%v", id), out); err != nil {
		return nil, err
	}
	return out, nil
}

// Update Update a domain
func (s *DomainsService) Update(ctx context.Context, id string, params *UpdateDomainRequest) (*DomainResponse, error) {
	out := new(DomainResponse)
	if err := s.client.request(ctx, "PATCH", fmt.Sprintf("/domains/%v", id), out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// Delete Delete a domain
func (s *DomainsService) Delete(ctx context.Context, id string) error {
	if err := s.client.request(ctx, "DELETE", fmt.Sprintf("/domains/%v", id), nil); err != nil {
		return err
	}
	return nil
}

// Verify Verify a domain
func (s *DomainsService) Verify(ctx context.Context, id string) (*DnsRecordsResponse, error) {
	out := new(DnsRecordsResponse)
	if err := s.client.request(ctx, "POST", fmt.Sprintf("/domains/%v/verify", id), out); err != nil {
		return nil, err
	}
	return out, nil
}

// RotateKey Rotate a domain key
func (s *DomainsService) RotateKey(ctx context.Context, id string) (*DnsRecordResponse, error) {
	out := new(DnsRecordResponse)
	if err := s.client.request(ctx, "POST", fmt.Sprintf("/domains/%v/rotate-key", id), out); err != nil {
		return nil, err
	}
	return out, nil
}

// DomainsListParams are the query parameters for List.
type DomainsListParams struct {
	Limit      int      `json:"limit,omitempty"`
	Offset     int      `json:"offset,omitempty"`
	Name       string   `json:"name,omitempty"`
	ChannelIds []string `json:"channelIds,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *DomainsListParams) toQuery() map[string]any {
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
