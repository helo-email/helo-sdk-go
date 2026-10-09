package helo

import (
	"context"
)

// SendingService exposes operations on the Sending resource.
type SendingService struct {
	client *Client
}

// SendTransactional Send a transactional email
func (s *SendingService) SendTransactional(ctx context.Context, params *SendMessageRequest, opts *SendingSendTransactionalOptions) (*SendMessageAcceptedResponse, error) {
	out := new(SendMessageAcceptedResponse)
	if err := s.client.request(ctx, "POST", "/send/transactional", out, withBody(params), withHeaders(opts.toHeaders())); err != nil {
		return nil, err
	}
	return out, nil
}

// SendTransactionalBatch Send transactional emails in batch
func (s *SendingService) SendTransactionalBatch(ctx context.Context, params *SendMessageBatchRequest, opts *SendingSendTransactionalBatchOptions) (*SendMessageBatchResponse, error) {
	out := new(SendMessageBatchResponse)
	if err := s.client.request(ctx, "POST", "/send/transactional/batch", out, withBody(params), withHeaders(opts.toHeaders())); err != nil {
		return nil, err
	}
	return out, nil
}

// SendBroadcast Send a broadcast
func (s *SendingService) SendBroadcast(ctx context.Context, params *SendBroadcastRequest, opts *SendingSendBroadcastOptions) (*SendBroadcastResponse, error) {
	out := new(SendBroadcastResponse)
	if err := s.client.request(ctx, "POST", "/send/broadcast", out, withBody(params), withHeaders(opts.toHeaders())); err != nil {
		return nil, err
	}
	return out, nil
}

// SendBroadcastMessage Send a single broadcast email
func (s *SendingService) SendBroadcastMessage(ctx context.Context, params *SendMessageRequest, opts *SendingSendBroadcastMessageOptions) (*SendMessageAcceptedResponse, error) {
	out := new(SendMessageAcceptedResponse)
	if err := s.client.request(ctx, "POST", "/send/broadcast/message", out, withBody(params), withHeaders(opts.toHeaders())); err != nil {
		return nil, err
	}
	return out, nil
}

// SendingSendTransactionalOptions are the optional headers for SendTransactional.
type SendingSendTransactionalOptions struct {
	// ChannelID Used to specify a channel ID for sending when using an account-level API credential.
	ChannelID string `json:"-"`
	// IdempotencyKey A unique identifier used to prevent duplicate messages being sent when retrying failed requests.
	IdempotencyKey string `json:"-"`
}

// toHeaders converts the options struct into a header map for the HTTP layer.
// A nil receiver or zero-valued fields contribute no headers.
func (o *SendingSendTransactionalOptions) toHeaders() map[string]string {
	if o == nil {
		return nil
	}
	h := map[string]string{}
	if o.ChannelID != "" {
		h["X-Helo-Channel-Id"] = o.ChannelID
	}
	if o.IdempotencyKey != "" {
		h["X-Helo-Idempotency-Key"] = o.IdempotencyKey
	}
	return h
}

// SendingSendTransactionalBatchOptions are the optional headers for SendTransactionalBatch.
type SendingSendTransactionalBatchOptions struct {
	// ChannelID Used to specify a channel ID for sending when using an account-level API credential.
	ChannelID string `json:"-"`
	// IdempotencyKey A unique identifier used to prevent duplicate messages being sent when retrying failed requests.
	IdempotencyKey string `json:"-"`
}

// toHeaders converts the options struct into a header map for the HTTP layer.
// A nil receiver or zero-valued fields contribute no headers.
func (o *SendingSendTransactionalBatchOptions) toHeaders() map[string]string {
	if o == nil {
		return nil
	}
	h := map[string]string{}
	if o.ChannelID != "" {
		h["X-Helo-Channel-Id"] = o.ChannelID
	}
	if o.IdempotencyKey != "" {
		h["X-Helo-Idempotency-Key"] = o.IdempotencyKey
	}
	return h
}

// SendingSendBroadcastOptions are the optional headers for SendBroadcast.
type SendingSendBroadcastOptions struct {
	// ChannelID Used to specify a channel ID for sending when using an account-level API credential.
	ChannelID string `json:"-"`
	// IdempotencyKey A unique identifier used to prevent duplicate messages being sent when retrying failed requests.
	IdempotencyKey string `json:"-"`
}

// toHeaders converts the options struct into a header map for the HTTP layer.
// A nil receiver or zero-valued fields contribute no headers.
func (o *SendingSendBroadcastOptions) toHeaders() map[string]string {
	if o == nil {
		return nil
	}
	h := map[string]string{}
	if o.ChannelID != "" {
		h["X-Helo-Channel-Id"] = o.ChannelID
	}
	if o.IdempotencyKey != "" {
		h["X-Helo-Idempotency-Key"] = o.IdempotencyKey
	}
	return h
}

// SendingSendBroadcastMessageOptions are the optional headers for SendBroadcastMessage.
type SendingSendBroadcastMessageOptions struct {
	// ChannelID Used to specify a channel ID for sending when using an account-level API credential.
	ChannelID string `json:"-"`
	// IdempotencyKey A unique identifier used to prevent duplicate messages being sent when retrying failed requests.
	IdempotencyKey string `json:"-"`
}

// toHeaders converts the options struct into a header map for the HTTP layer.
// A nil receiver or zero-valued fields contribute no headers.
func (o *SendingSendBroadcastMessageOptions) toHeaders() map[string]string {
	if o == nil {
		return nil
	}
	h := map[string]string{}
	if o.ChannelID != "" {
		h["X-Helo-Channel-Id"] = o.ChannelID
	}
	if o.IdempotencyKey != "" {
		h["X-Helo-Idempotency-Key"] = o.IdempotencyKey
	}
	return h
}
