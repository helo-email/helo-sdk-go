package helo

import (
	"context"
)

// SendingService exposes operations on the Sending resource.
type SendingService struct {
	client *Client
}

// Transactional Send a transactional email
func (s *SendingService) Transactional(ctx context.Context, params *SendMessageRequest, opts *SendingTransactionalOptions) (*SendMessageAcceptedResponse, error) {
	out := new(SendMessageAcceptedResponse)
	if err := s.client.request(ctx, "POST", "/send/transactional", out, withBody(params), withHeaders(opts.toHeaders())); err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionalBatch Send transactional emails in batch
func (s *SendingService) TransactionalBatch(ctx context.Context, params *SendMessageBatchRequest, opts *SendingTransactionalBatchOptions) (*SendMessageBatchResponse, error) {
	out := new(SendMessageBatchResponse)
	if err := s.client.request(ctx, "POST", "/send/transactional/batch", out, withBody(params), withHeaders(opts.toHeaders())); err != nil {
		return nil, err
	}
	return out, nil
}

// Broadcast Send a broadcast email
func (s *SendingService) Broadcast(ctx context.Context, params *SendBroadcastRequest, opts *SendingBroadcastOptions) (*SendBroadcastResponse, error) {
	out := new(SendBroadcastResponse)
	if err := s.client.request(ctx, "POST", "/send/broadcast", out, withBody(params), withHeaders(opts.toHeaders())); err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcastMessage Send a single broadcast email
func (s *SendingService) BroadcastMessage(ctx context.Context, params *SendMessageRequest, opts *SendingBroadcastMessageOptions) (*SendMessageAcceptedResponse, error) {
	out := new(SendMessageAcceptedResponse)
	if err := s.client.request(ctx, "POST", "/send/broadcast/message", out, withBody(params), withHeaders(opts.toHeaders())); err != nil {
		return nil, err
	}
	return out, nil
}

// SendingTransactionalOptions are the optional headers for Transactional.
type SendingTransactionalOptions struct {
	// ChannelID Used to specify a channel ID for sending when using an account-level API credential.
	ChannelID string `json:"-"`
	// IdempotencyKey A unique identifier used to prevent duplicate messages being sent when retrying failed requests.
	IdempotencyKey string `json:"-"`
}

// toHeaders converts the options struct into a header map for the HTTP layer.
// A nil receiver or zero-valued fields contribute no headers.
func (o *SendingTransactionalOptions) toHeaders() map[string]string {
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

// SendingTransactionalBatchOptions are the optional headers for TransactionalBatch.
type SendingTransactionalBatchOptions struct {
	// ChannelID Used to specify a channel ID for sending when using an account-level API credential.
	ChannelID string `json:"-"`
	// IdempotencyKey A unique identifier used to prevent duplicate messages being sent when retrying failed requests.
	IdempotencyKey string `json:"-"`
}

// toHeaders converts the options struct into a header map for the HTTP layer.
// A nil receiver or zero-valued fields contribute no headers.
func (o *SendingTransactionalBatchOptions) toHeaders() map[string]string {
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

// SendingBroadcastOptions are the optional headers for Broadcast.
type SendingBroadcastOptions struct {
	// ChannelID Used to specify a channel ID for sending when using an account-level API credential.
	ChannelID string `json:"-"`
	// IdempotencyKey A unique identifier used to prevent duplicate messages being sent when retrying failed requests.
	IdempotencyKey string `json:"-"`
}

// toHeaders converts the options struct into a header map for the HTTP layer.
// A nil receiver or zero-valued fields contribute no headers.
func (o *SendingBroadcastOptions) toHeaders() map[string]string {
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

// SendingBroadcastMessageOptions are the optional headers for BroadcastMessage.
type SendingBroadcastMessageOptions struct {
	// ChannelID Used to specify a channel ID for sending when using an account-level API credential.
	ChannelID string `json:"-"`
	// IdempotencyKey A unique identifier used to prevent duplicate messages being sent when retrying failed requests.
	IdempotencyKey string `json:"-"`
}

// toHeaders converts the options struct into a header map for the HTTP layer.
// A nil receiver or zero-valued fields contribute no headers.
func (o *SendingBroadcastMessageOptions) toHeaders() map[string]string {
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
