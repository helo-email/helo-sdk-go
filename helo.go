// Package helo provides a Go SDK for the Helo API.
package helo

import "net/http"

// DefaultBaseURL is the API endpoint used unless overridden via WithBaseURL.
const DefaultBaseURL = "https://api.helohq.com"

// Helo is the entry-point client for the SDK.
type Helo struct {
	client *Client

	Channels     *ChannelsService
	Activity     *ActivityService
	Domains      *DomainsService
	Sending      *SendingService
	Broadcasts   *BroadcastsService
	Statistics   *StatisticsService
	Suppressions *SuppressionsService
	Webhooks     *WebhooksService
}

// HeloOption configures a Helo client.
type HeloOption func(*Helo)

// WithBaseURL overrides the default API base URL.
func WithBaseURL(baseURL string) HeloOption {
	return func(c *Helo) { c.client.BaseURL = baseURL }
}

// WithHTTPClient sets a custom *http.Client used for all requests.
func WithHTTPClient(httpClient *http.Client) HeloOption {
	return func(c *Helo) { c.client.HTTPClient = httpClient }
}

// NewHelo creates a new Helo client.
func NewHelo(apiKey string, opts ...HeloOption) *Helo {
	c := &Helo{
		client: &Client{
			BaseURL:    DefaultBaseURL,
			APIKey:     apiKey,
			HTTPClient: http.DefaultClient,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Channels = &ChannelsService{client: c.client}
	c.Activity = &ActivityService{client: c.client}
	c.Domains = &DomainsService{client: c.client}
	c.Sending = &SendingService{client: c.client}
	c.Broadcasts = &BroadcastsService{client: c.client}
	c.Statistics = &StatisticsService{client: c.client}
	c.Suppressions = &SuppressionsService{client: c.client}
	c.Webhooks = &WebhooksService{client: c.client}

	return c
}
