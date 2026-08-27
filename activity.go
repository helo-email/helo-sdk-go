package helo

import (
	"context"
	"fmt"
	"time"
)

// ActivityService exposes operations on the Activity resource.
type ActivityService struct {
	client *Client
}

// ListEvents List activity events
func (s *ActivityService) ListEvents(ctx context.Context, params *ActivityListEventsParams) (*PaginatedEventsResponse, error) {
	out := new(PaginatedEventsResponse)
	if err := s.client.request(ctx, "GET", "/activity/events", out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// ListMessages List messages
func (s *ActivityService) ListMessages(ctx context.Context, params *ActivityListMessagesParams) (*PaginatedMessagesResponse, error) {
	out := new(PaginatedMessagesResponse)
	if err := s.client.request(ctx, "GET", "/activity/messages", out, withQuery(params.toQuery())); err != nil {
		return nil, err
	}
	return out, nil
}

// RetrieveMessage Retrieve message details
func (s *ActivityService) RetrieveMessage(ctx context.Context, id string) (*MessageDetailsResponse, error) {
	out := new(MessageDetailsResponse)
	if err := s.client.request(ctx, "GET", fmt.Sprintf("/activity/messages/%v", id), out); err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityListEventsParams are the query parameters for ListEvents.
type ActivityListEventsParams struct {
	ChannelID  string      `json:"channelId,omitempty"`
	MessageID  string      `json:"messageId,omitempty"`
	After      int64       `json:"after,omitempty"`
	StartDate  time.Time   `json:"startDate,omitempty"`
	EndDate    time.Time   `json:"endDate,omitempty"`
	Limit      int         `json:"limit,omitempty"`
	Recipient  string      `json:"recipient,omitempty"`
	Subject    string      `json:"subject,omitempty"`
	Tags       []string    `json:"tags,omitempty"`
	MailType   string      `json:"mailType,omitempty"`
	EventTypes []EventType `json:"eventTypes,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *ActivityListEventsParams) toQuery() map[string]any {
	if p == nil {
		return nil
	}
	q := map[string]any{}
	if p.ChannelID != "" {
		q["channelId"] = p.ChannelID
	}
	if p.MessageID != "" {
		q["messageId"] = p.MessageID
	}
	if p.After != 0 {
		q["after"] = p.After
	}
	if !p.StartDate.IsZero() {
		q["startDate"] = p.StartDate.Format(time.RFC3339)
	}
	if !p.EndDate.IsZero() {
		q["endDate"] = p.EndDate.Format(time.RFC3339)
	}
	if p.Limit != 0 {
		q["limit"] = p.Limit
	}
	if p.Recipient != "" {
		q["recipient"] = p.Recipient
	}
	if p.Subject != "" {
		q["subject"] = p.Subject
	}
	if len(p.Tags) > 0 {
		q["tags"] = p.Tags
	}
	if p.MailType != "" {
		q["mailType"] = p.MailType
	}
	if len(p.EventTypes) > 0 {
		q["eventTypes"] = p.EventTypes
	}
	return q
}

// ActivityListMessagesParams are the query parameters for ListMessages.
type ActivityListMessagesParams struct {
	ChannelID string        `json:"channelId,omitempty"`
	After     int64         `json:"after,omitempty"`
	StartDate time.Time     `json:"startDate,omitempty"`
	EndDate   time.Time     `json:"endDate,omitempty"`
	Limit     int           `json:"limit,omitempty"`
	Recipient string        `json:"recipient,omitempty"`
	Subject   string        `json:"subject,omitempty"`
	Tags      []string      `json:"tags,omitempty"`
	MailType  string        `json:"mailType,omitempty"`
	Status    MessageStatus `json:"status,omitempty"`
}

// toQuery converts the params struct into a map suitable for the HTTP layer.
// Zero-valued optional fields are omitted, mirroring the JSON tag `omitempty` behaviour.
func (p *ActivityListMessagesParams) toQuery() map[string]any {
	if p == nil {
		return nil
	}
	q := map[string]any{}
	if p.ChannelID != "" {
		q["channelId"] = p.ChannelID
	}
	if p.After != 0 {
		q["after"] = p.After
	}
	if !p.StartDate.IsZero() {
		q["startDate"] = p.StartDate.Format(time.RFC3339)
	}
	if !p.EndDate.IsZero() {
		q["endDate"] = p.EndDate.Format(time.RFC3339)
	}
	if p.Limit != 0 {
		q["limit"] = p.Limit
	}
	if p.Recipient != "" {
		q["recipient"] = p.Recipient
	}
	if p.Subject != "" {
		q["subject"] = p.Subject
	}
	if len(p.Tags) > 0 {
		q["tags"] = p.Tags
	}
	if p.MailType != "" {
		q["mailType"] = p.MailType
	}
	if p.Status != "" {
		q["status"] = p.Status
	}
	return q
}
