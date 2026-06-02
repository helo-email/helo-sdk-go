package helo

import (
	"context"
)

// SendingService exposes operations on the Sending resource.
type SendingService struct {
	client *Client
}

// Transactional Send a transactional email
func (s *SendingService) Transactional(ctx context.Context, params *SendMessageRequest) (*SendMessageAcceptedResponse, error) {
	out := new(SendMessageAcceptedResponse)
	if err := s.client.request(ctx, "POST", "/send/transactional", out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionalBatch Send transactional emails in batch
func (s *SendingService) TransactionalBatch(ctx context.Context, params *SendMessageBatchRequest) (*SendMessageBatchResponse, error) {
	out := new(SendMessageBatchResponse)
	if err := s.client.request(ctx, "POST", "/send/transactional/batch", out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// Broadcast Send a broadcast email
func (s *SendingService) Broadcast(ctx context.Context, params *SendBroadcastRequest) (*SendBroadcastResponse, error) {
	out := new(SendBroadcastResponse)
	if err := s.client.request(ctx, "POST", "/send/broadcast", out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcastMessage Send a single broadcast email
func (s *SendingService) BroadcastMessage(ctx context.Context, params *SendMessageRequest) (*SendMessageAcceptedResponse, error) {
	out := new(SendMessageAcceptedResponse)
	if err := s.client.request(ctx, "POST", "/send/broadcast/message", out, withBody(params)); err != nil {
		return nil, err
	}
	return out, nil
}
