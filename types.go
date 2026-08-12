package helo

import "time"

// ChannelDetailsResponse is a generated schema type.
type ChannelDetailsResponse struct {
	ID           string           `json:"id,omitempty"`
	Name         string           `json:"name,omitempty"`
	DeliveryType DeliveryType     `json:"deliveryType,omitempty"`
	CreatedAt    time.Time        `json:"createdAt,omitempty"`
	UpdatedAt    time.Time        `json:"updatedAt,omitempty"`
	Tracking     *ChannelTracking `json:"tracking,omitempty"`
}

// ChannelBasicResponse is a generated schema type.
type ChannelBasicResponse struct {
	ID           string       `json:"id,omitempty"`
	Name         string       `json:"name,omitempty"`
	DeliveryType DeliveryType `json:"deliveryType,omitempty"`
	CreatedAt    time.Time    `json:"createdAt,omitempty"`
}

// ChannelTracking is a generated schema type.
type ChannelTracking struct {
	Links bool `json:"links"`
	Opens bool `json:"opens"`
}

// CreateChannelRequest is a generated schema type.
type CreateChannelRequest struct {
	Name         string                 `json:"name"`
	DeliveryType DeliveryType           `json:"deliveryType"`
	Tracking     *CreateChannelTracking `json:"tracking,omitempty"`
}

// CreateChannelTracking is a generated schema type.
type CreateChannelTracking struct {
	Links bool `json:"links,omitempty"`
	Opens bool `json:"opens,omitempty"`
}

// ValidationError A single field validation error.
type ValidationError struct {
	Message string `json:"message"`
}

// ErrorResponse Standard error response body.
type ErrorResponse struct {
	Type      string         `json:"type,omitempty"`
	Title     string         `json:"title,omitempty"`
	Instance  string         `json:"instance,omitempty"`
	Status    int            `json:"status,omitempty"`
	Code      string         `json:"code,omitempty"`
	Detail    string         `json:"detail,omitempty"`
	RequestID string         `json:"requestId,omitempty"`
	Errors    map[string]any `json:"errors,omitempty"`
}

// PaginationResultOfChannelBasicResponse is a generated schema type.
type PaginationResultOfChannelBasicResponse struct {
	Results    []ChannelBasicResponse `json:"results"`
	TotalCount int                    `json:"totalCount"`
}

// UpdateChannelRequest is a generated schema type.
type UpdateChannelRequest struct {
	Name         string                 `json:"name,omitempty"`
	DeliveryType DeliveryType           `json:"deliveryType,omitempty"`
	Tracking     *UpdateChannelTracking `json:"tracking,omitempty"`
}

// UpdateChannelTracking is a generated schema type.
type UpdateChannelTracking struct {
	Links bool `json:"links,omitempty"`
	Opens bool `json:"opens,omitempty"`
}

// ActivityMailAddress is a generated schema type.
type ActivityMailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

// ActivityEvent is a generated schema type.
type ActivityEvent struct {
	MessageID  string         `json:"messageId"`
	ChannelID  string         `json:"channelId"`
	MailType   string         `json:"mailType"`
	MailSource string         `json:"mailSource,omitempty"`
	EventType  EventType      `json:"eventType"`
	Timestamp  time.Time      `json:"timestamp"`
	Subject    string         `json:"subject"`
	Recipients []string       `json:"recipients"`
	Tags       []string       `json:"tags,omitempty"`
	Metadata   map[string]any `json:"metadata,omitempty"`
	Details    map[string]any `json:"details,omitempty"`
}

// PaginatedEventsResponse is a generated schema type.
type PaginatedEventsResponse struct {
	After      int64           `json:"after,omitempty"`
	TotalCount float64         `json:"totalCount"`
	Results    []ActivityEvent `json:"results"`
}

// PaginatedMessagesResponse is a generated schema type.
type PaginatedMessagesResponse struct {
	After      int64     `json:"after,omitempty"`
	TotalCount float64   `json:"totalCount"`
	Results    []Message `json:"results"`
}

// Message is a generated schema type.
type Message struct {
	MessageID    string            `json:"messageId"`
	ChannelID    string            `json:"channelId"`
	Timestamp    time.Time         `json:"timestamp"`
	MailType     string            `json:"mailType"`
	MailSource   string            `json:"mailSource"`
	DeliveryType string            `json:"deliveryType"`
	Status       string            `json:"status"`
	Subject      string            `json:"subject"`
	Recipients   []string          `json:"recipients"`
	Tags         []string          `json:"tags,omitempty"`
	Statistics   MessageStatistics `json:"statistics"`
}

// MessageStatistics is a generated schema type.
type MessageStatistics struct {
	Delivered    int `json:"delivered"`
	Bounced      int `json:"bounced"`
	Opened       int `json:"opened"`
	Clicked      int `json:"clicked"`
	Complained   int `json:"complained"`
	Unsubscribed int `json:"unsubscribed"`
}

// MessageDetailsResponse is a generated schema type.
type MessageDetailsResponse struct {
	MessageID    string                             `json:"messageId"`
	ChannelID    string                             `json:"channelId"`
	Timestamp    time.Time                          `json:"timestamp"`
	MailType     string                             `json:"mailType"`
	MailSource   string                             `json:"mailSource"`
	DeliveryType string                             `json:"deliveryType"`
	Status       string                             `json:"status"`
	Subject      string                             `json:"subject"`
	From         ActivityMailAddress                `json:"from"`
	To           []ActivityMailAddress              `json:"to"`
	Cc           []ActivityMailAddress              `json:"cc,omitempty"`
	Bcc          []ActivityMailAddress              `json:"bcc,omitempty"`
	ReplyTo      []ActivityMailAddress              `json:"replyTo,omitempty"`
	Text         string                             `json:"text,omitempty"`
	Html         string                             `json:"html,omitempty"`
	Body         string                             `json:"body,omitempty"`
	Tags         []string                           `json:"tags,omitempty"`
	Headers      map[string]any                     `json:"headers,omitempty"`
	Metadata     map[string]any                     `json:"metadata,omitempty"`
	Attachments  []MessageDetailsResponseAttachment `json:"attachments,omitempty"`
	Tracking     MessageDetailsResponseTracking     `json:"tracking"`
	Events       []MessageDetailsResponseEvent      `json:"events"`
	Statistics   *MessageStatistics                 `json:"statistics,omitempty"`
}

// CreateDomainRequest is a generated schema type.
type CreateDomainRequest struct {
	Name       string   `json:"name"`
	ChannelIds []string `json:"channelIds,omitempty"`
}

// DnsRecordResponse is a generated schema type.
type DnsRecordResponse struct {
	Type          DnsRecordType   `json:"type,omitempty"`
	Host          string          `json:"host,omitempty"`
	Value         string          `json:"value,omitempty"`
	Status        DnsRecordStatus `json:"status,omitempty"`
	LastCheckedAt time.Time       `json:"lastCheckedAt,omitempty"`
}

// DnsRecordsResponse is a generated schema type.
type DnsRecordsResponse struct {
	DomainKeyActive  *DnsRecordResponse  `json:"domainKeyActive,omitempty"`
	DomainKeyPending *DnsRecordResponse  `json:"domainKeyPending,omitempty"`
	ReturnPath       []DnsRecordResponse `json:"returnPath,omitempty"`
}

// DomainChannelResponse is a generated schema type.
type DomainChannelResponse struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	DeliveryType DeliveryType `json:"deliveryType"`
	Deleted      bool         `json:"deleted,omitempty"`
}

// DomainResponse is a generated schema type.
type DomainResponse struct {
	ID        string                  `json:"id"`
	CreatedAt time.Time               `json:"createdAt"`
	Name      string                  `json:"name"`
	Verified  bool                    `json:"verified"`
	Channels  []DomainChannelResponse `json:"channels,omitempty"`
}

// DomainWithDnsResponse is a generated schema type.
type DomainWithDnsResponse struct {
	ID         string                  `json:"id"`
	CreatedAt  time.Time               `json:"createdAt"`
	Name       string                  `json:"name"`
	Verified   bool                    `json:"verified"`
	Channels   []DomainChannelResponse `json:"channels,omitempty"`
	DnsRecords DnsRecordsResponse      `json:"dnsRecords"`
}

// PaginatedResponseOfDomainResponse is a generated schema type.
type PaginatedResponseOfDomainResponse struct {
	TotalCount int              `json:"totalCount"`
	Results    []DomainResponse `json:"results"`
}

// UpdateDomainRequest is a generated schema type.
type UpdateDomainRequest struct {
	ChannelIds []string `json:"channelIds,omitempty"`
}

// MailAddress is a generated schema type.
type MailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

// Attachment is a generated schema type.
type Attachment struct {
	Content     string                `json:"content"`
	ContentID   string                `json:"contentId,omitempty"`
	ContentType string                `json:"contentType,omitempty"`
	FileName    string                `json:"fileName"`
	Disposition AttachmentDisposition `json:"disposition"`
}

// SendBroadcastRequest is a generated schema type.
type SendBroadcastRequest struct {
	From        MailAddress                   `json:"from"`
	ReplyTo     []MailAddress                 `json:"replyTo,omitempty"`
	Template    SendBroadcastRequestTemplate  `json:"template"`
	Tracking    *SendBroadcastRequestTracking `json:"tracking,omitempty"`
	Attachments []Attachment                  `json:"attachments,omitempty"`
	Tags        []string                      `json:"tags,omitempty"`
	Headers     map[string]any                `json:"headers,omitempty"`
	Metadata    map[string]any                `json:"metadata,omitempty"`
	Messages    []SendBroadcastRequestMessage `json:"messages"`
}

// SendBroadcastResponse is a generated schema type.
type SendBroadcastResponse struct {
	Status      string `json:"status,omitempty"`
	BroadcastID string `json:"broadcastId,omitempty"`
}

// SendMessageRequest is a generated schema type.
type SendMessageRequest struct {
	From        MailAddress                 `json:"from"`
	To          []MailAddress               `json:"to"`
	Cc          []MailAddress               `json:"cc,omitempty"`
	Bcc         []MailAddress               `json:"bcc,omitempty"`
	ReplyTo     []MailAddress               `json:"replyTo,omitempty"`
	Subject     string                      `json:"subject,omitempty"`
	Html        string                      `json:"html,omitempty"`
	Text        string                      `json:"text,omitempty"`
	Template    *SendMessageRequestTemplate `json:"template,omitempty"`
	Tracking    *SendMessageRequestTracking `json:"tracking,omitempty"`
	Attachments []Attachment                `json:"attachments,omitempty"`
	Tags        []string                    `json:"tags,omitempty"`
	Headers     map[string]any              `json:"headers,omitempty"`
	Metadata    map[string]any              `json:"metadata,omitempty"`
}

// SendMessageResponse is a generated schema type.
type SendMessageResponse struct {
}

// SendMessageAcceptedResponse is a generated schema type.
type SendMessageAcceptedResponse struct {
	Status       string   `json:"status,omitempty"`
	MessageID    string   `json:"messageId,omitempty"`
	Suppressions []string `json:"suppressions,omitempty"`
}

// SendMessageFailedResponse is a generated schema type.
type SendMessageFailedResponse struct {
	Status       string `json:"status"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

// SendMessageBatchRequest is a generated schema type.
type SendMessageBatchRequest struct {
	Requests []SendMessageRequest `json:"requests"`
}

// SendMessageBatchResponse is a generated schema type.
type SendMessageBatchResponse struct {
	Responses []SendMessageResponse `json:"responses"`
}

// BroadcastResponse is a generated schema type.
type BroadcastResponse struct {
	ID         string          `json:"id"`
	CreatedAt  time.Time       `json:"createdAt"`
	Status     BroadcastStatus `json:"status"`
	Subject    string          `json:"subject"`
	Completion string          `json:"completion"`
	Messages   int             `json:"messages"`
}

// BroadcastDetailsResponse is a generated schema type.
type BroadcastDetailsResponse struct {
	ID         string              `json:"id"`
	CreatedAt  time.Time           `json:"createdAt"`
	Status     BroadcastStatus     `json:"status"`
	Subject    string              `json:"subject"`
	Completion string              `json:"completion"`
	Messages   int                 `json:"messages"`
	Failed     int                 `json:"failed"`
	Suppressed int                 `json:"suppressed"`
	Content    BroadcastContent    `json:"content"`
	Tracking   BroadcastTracking   `json:"tracking"`
	Statistics BroadcastStatistics `json:"statistics"`
}

// BroadcastContent is a generated schema type.
type BroadcastContent struct {
	From        *MailAddress                 `json:"from,omitempty"`
	ReplyTo     []MailAddress                `json:"replyTo,omitempty"`
	Template    *BroadcastContentTemplate    `json:"template,omitempty"`
	Attachments []BroadcastContentAttachment `json:"attachments,omitempty"`
	Tags        []string                     `json:"tags,omitempty"`
	Headers     map[string]any               `json:"headers,omitempty"`
	Metadata    map[string]any               `json:"metadata,omitempty"`
}

// BroadcastTracking is a generated schema type.
type BroadcastTracking struct {
	Opens bool `json:"opens"`
	Links bool `json:"links"`
}

// BroadcastStatistics is a generated schema type.
type BroadcastStatistics struct {
	Sent         int `json:"sent"`
	Delivered    int `json:"delivered"`
	Bounced      int `json:"bounced"`
	Opened       int `json:"opened"`
	Clicked      int `json:"clicked"`
	Complained   int `json:"complained"`
	Unsubscribed int `json:"unsubscribed"`
}

// BroadcastFailureResponse is a generated schema type.
type BroadcastFailureResponse struct {
	Recipients   RecipientHeaders `json:"recipients"`
	MessageIndex int              `json:"messageIndex"`
	ErrorCode    string           `json:"errorCode"`
	ErrorMessage string           `json:"errorMessage"`
}

// RecipientHeaders is a generated schema type.
type RecipientHeaders struct {
	To  []MailAddress `json:"to"`
	Cc  []MailAddress `json:"cc,omitempty"`
	Bcc []MailAddress `json:"bcc,omitempty"`
}

// PaginatedResponseOfBroadcast is a generated schema type.
type PaginatedResponseOfBroadcast struct {
	TotalCount int                 `json:"totalCount"`
	Results    []BroadcastResponse `json:"results"`
}

// PaginatedResponseOfBroadcastFailure is a generated schema type.
type PaginatedResponseOfBroadcastFailure struct {
	TotalCount int                        `json:"totalCount"`
	Results    []BroadcastFailureResponse `json:"results"`
}

// PaginatedResponseOfBroadcastSuppression is a generated schema type.
type PaginatedResponseOfBroadcastSuppression struct {
	TotalCount int      `json:"totalCount"`
	Results    []string `json:"results"`
}

// DeliveryStats is a generated schema type.
type DeliveryStats struct {
	Sent         int `json:"sent,omitempty"`
	Delivered    int `json:"delivered,omitempty"`
	Opened       int `json:"opened,omitempty"`
	Clicked      int `json:"clicked,omitempty"`
	Bounced      int `json:"bounced,omitempty"`
	Unsubscribed int `json:"unsubscribed,omitempty"`
	Complained   int `json:"complained,omitempty"`
}

// StatisticsTotalsResponse is a generated schema type.
type StatisticsTotalsResponse struct {
	Transactional *DeliveryStats `json:"transactional,omitempty"`
	Broadcast     *DeliveryStats `json:"broadcast,omitempty"`
}

// StatisticsDailyResponse is a generated schema type.
type StatisticsDailyResponse struct {
	Results []StatisticsDailyResponseResult `json:"results,omitempty"`
}

// StatisticsHourlyResponse is a generated schema type.
type StatisticsHourlyResponse struct {
	Results []StatisticsHourlyResponseResult `json:"results,omitempty"`
}

// PaginatedResponseOfSuppressionResponse is a generated schema type.
type PaginatedResponseOfSuppressionResponse struct {
	TotalCount int                   `json:"totalCount"`
	Results    []SuppressionResponse `json:"results"`
}

// SuppressionResponse is a generated schema type.
type SuppressionResponse struct {
	Email     string            `json:"email"`
	Reason    SuppressionReason `json:"reason"`
	CreatedAt time.Time         `json:"createdAt"`
}

// CreateSuppressionsRequest is a generated schema type.
type CreateSuppressionsRequest struct {
	ChannelID string   `json:"channelId"`
	MailType  MailType `json:"mailType"`
	Emails    []string `json:"emails"`
}

// CreateSuppressionsResponse is a generated schema type.
type CreateSuppressionsResponse struct {
	Results []SuppressionResult `json:"results"`
}

// SuppressionResult is a generated schema type.
type SuppressionResult struct {
	Email   string `json:"email"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// RemoveSuppressionsRequest is a generated schema type.
type RemoveSuppressionsRequest struct {
	ChannelID string   `json:"channelId"`
	MailType  MailType `json:"mailType"`
	Emails    []string `json:"emails"`
}

// RemoveSuppressionsResponse is a generated schema type.
type RemoveSuppressionsResponse struct {
	Results []RemoveSuppressionResult `json:"results"`
}

// RemoveSuppressionResult is a generated schema type.
type RemoveSuppressionResult struct {
	Email   string `json:"email"`
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// CreateWebhookRequest Request body for creating a new webhook.
type CreateWebhookRequest struct {
	URL               string          `json:"url"`
	Events            []WebhookEvent  `json:"events"`
	ChannelID         string          `json:"channelId,omitempty"`
	AdditionalHeaders []WebhookHeader `json:"additionalHeaders,omitempty"`
	Enabled           bool            `json:"enabled,omitempty"`
}

// WebhookHeader A custom HTTP header to include in webhook deliveries.
type WebhookHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// PaginationResultOfWebhookResponse Paginated list of webhooks.
type PaginationResultOfWebhookResponse struct {
	Results    []WebhookResponse `json:"results"`
	TotalCount int               `json:"totalCount"`
}

// UpdateWebhookRequest Request body for updating a webhook. Only provided fields are changed.
type UpdateWebhookRequest struct {
	URL               string          `json:"url,omitempty"`
	Events            []WebhookEvent  `json:"events,omitempty"`
	ChannelID         string          `json:"channelId,omitempty"`
	AdditionalHeaders []WebhookHeader `json:"additionalHeaders,omitempty"`
	Enabled           bool            `json:"enabled,omitempty"`
}

// WebhookResponse Webhook configuration properties.
type WebhookResponse struct {
	ID                string               `json:"id,omitempty"`
	ChannelID         string               `json:"channelId,omitempty"`
	URL               string               `json:"url,omitempty"`
	PayloadSigningKey string               `json:"payloadSigningKey,omitempty"`
	Enabled           bool                 `json:"enabled,omitempty"`
	AdditionalHeaders []WebhookHeader      `json:"additionalHeaders,omitempty"`
	Events            []WebhookEvent       `json:"events,omitempty"`
	LastResponse      *WebhookLastResponse `json:"lastResponse,omitempty"`
}

// WebhookLastResponse The most recent delivery outcome recorded for a webhook.
type WebhookLastResponse struct {
	StatusCode int       `json:"statusCode,omitempty"`
	Error      string    `json:"error,omitempty"`
	At         time.Time `json:"at,omitempty"`
}

// WebhookPayloadCommon Fields common to every webhook payload.
type WebhookPayloadCommon struct {
	Timestamp time.Time      `json:"timestamp,omitempty"`
	MessageID string         `json:"messageId,omitempty"`
	ChannelID string         `json:"channelId,omitempty"`
	MailType  MailType       `json:"mailType,omitempty"`
	Subject   string         `json:"subject,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

// AcceptedWebhookPayload Payload delivered for the `accepted` event.
type AcceptedWebhookPayload struct {
	EventType  string         `json:"eventType"`
	Recipients []string       `json:"recipients"`
	Timestamp  time.Time      `json:"timestamp"`
	MessageID  string         `json:"messageId"`
	ChannelID  string         `json:"channelId"`
	MailType   MailType       `json:"mailType"`
	Subject    string         `json:"subject,omitempty"`
	Tags       []string       `json:"tags,omitempty"`
	Metadata   map[string]any `json:"metadata,omitempty"`
}

// ProcessedWebhookPayload Payload delivered for the `processed` event.
type ProcessedWebhookPayload struct {
	EventType  string         `json:"eventType"`
	Recipients []string       `json:"recipients"`
	Timestamp  time.Time      `json:"timestamp"`
	MessageID  string         `json:"messageId"`
	ChannelID  string         `json:"channelId"`
	MailType   MailType       `json:"mailType"`
	Subject    string         `json:"subject,omitempty"`
	Tags       []string       `json:"tags,omitempty"`
	Metadata   map[string]any `json:"metadata,omitempty"`
}

// DeliveredWebhookPayload Payload delivered for the `delivered` event.
type DeliveredWebhookPayload struct {
	EventType string         `json:"eventType"`
	Recipient string         `json:"recipient"`
	Details   any            `json:"details,omitempty"`
	Timestamp time.Time      `json:"timestamp"`
	MessageID string         `json:"messageId"`
	ChannelID string         `json:"channelId"`
	MailType  MailType       `json:"mailType"`
	Subject   string         `json:"subject,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

// BouncedWebhookPayload Payload delivered for the `bounced` event.
type BouncedWebhookPayload struct {
	EventType string         `json:"eventType"`
	Recipient string         `json:"recipient"`
	Details   any            `json:"details,omitempty"`
	Timestamp time.Time      `json:"timestamp"`
	MessageID string         `json:"messageId"`
	ChannelID string         `json:"channelId"`
	MailType  MailType       `json:"mailType"`
	Subject   string         `json:"subject,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

// OpenedWebhookPayload Payload delivered for the `opened` event.
type OpenedWebhookPayload struct {
	EventType string         `json:"eventType"`
	Recipient string         `json:"recipient"`
	Details   any            `json:"details,omitempty"`
	Timestamp time.Time      `json:"timestamp"`
	MessageID string         `json:"messageId"`
	ChannelID string         `json:"channelId"`
	MailType  MailType       `json:"mailType"`
	Subject   string         `json:"subject,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

// ClickedWebhookPayload Payload delivered for the `clicked` event.
type ClickedWebhookPayload struct {
	EventType string         `json:"eventType"`
	Recipient string         `json:"recipient"`
	Details   any            `json:"details,omitempty"`
	Timestamp time.Time      `json:"timestamp"`
	MessageID string         `json:"messageId"`
	ChannelID string         `json:"channelId"`
	MailType  MailType       `json:"mailType"`
	Subject   string         `json:"subject,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

// ComplainedWebhookPayload Payload delivered for the `complained` event.
type ComplainedWebhookPayload struct {
	EventType string         `json:"eventType"`
	Recipient string         `json:"recipient"`
	Details   any            `json:"details,omitempty"`
	Timestamp time.Time      `json:"timestamp"`
	MessageID string         `json:"messageId"`
	ChannelID string         `json:"channelId"`
	MailType  MailType       `json:"mailType"`
	Subject   string         `json:"subject,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

// UnsubscribedWebhookPayload Payload delivered for the `unsubscribed` event.
type UnsubscribedWebhookPayload struct {
	EventType string         `json:"eventType"`
	Recipient string         `json:"recipient"`
	Details   any            `json:"details,omitempty"`
	Timestamp time.Time      `json:"timestamp"`
	MessageID string         `json:"messageId"`
	ChannelID string         `json:"channelId"`
	MailType  MailType       `json:"mailType"`
	Subject   string         `json:"subject,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

// ResubscribedWebhookPayload Payload delivered for the `resubscribed` event.
type ResubscribedWebhookPayload struct {
	EventType string         `json:"eventType"`
	Recipient string         `json:"recipient"`
	Details   any            `json:"details,omitempty"`
	Timestamp time.Time      `json:"timestamp"`
	MessageID string         `json:"messageId"`
	ChannelID string         `json:"channelId"`
	MailType  MailType       `json:"mailType"`
	Subject   string         `json:"subject,omitempty"`
	Tags      []string       `json:"tags,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
}

// RecipientEventFields Fields shared by per-recipient events (delivered, bounced, opened, clicked, complained, unsubscribed, resubscribed).
type RecipientEventFields struct {
	Recipient string `json:"recipient"`
	Details   any    `json:"details,omitempty"`
}

// MessageDetailsResponseAttachment is an inline schema extracted from its parent type.
type MessageDetailsResponseAttachment struct {
	FileName    string  `json:"fileName"`
	Disposition string  `json:"disposition"`
	Size        float64 `json:"size"`
}

// MessageDetailsResponseTracking is an inline schema extracted from its parent type.
type MessageDetailsResponseTracking struct {
	Links bool `json:"links"`
	Opens bool `json:"opens"`
}

// MessageDetailsResponseEvent is an inline schema extracted from its parent type.
type MessageDetailsResponseEvent struct {
	EventType  EventType      `json:"eventType"`
	Timestamp  time.Time      `json:"timestamp"`
	Recipients []string       `json:"recipients"`
	Details    map[string]any `json:"details,omitempty"`
}

// SendBroadcastRequestTemplate is an inline schema extracted from its parent type.
type SendBroadcastRequestTemplate struct {
	Subject      string         `json:"subject"`
	Html         string         `json:"html,omitempty"`
	Text         string         `json:"text,omitempty"`
	InlineStyles bool           `json:"inlineStyles,omitempty"`
	Data         map[string]any `json:"data,omitempty"`
}

// SendBroadcastRequestTracking is an inline schema extracted from its parent type.
type SendBroadcastRequestTracking struct {
	Opens bool `json:"opens,omitempty"`
	Links bool `json:"links,omitempty"`
}

// SendBroadcastRequestMessage is an inline schema extracted from its parent type.
type SendBroadcastRequestMessage struct {
	To       []MailAddress  `json:"to"`
	Cc       []MailAddress  `json:"cc,omitempty"`
	Bcc      []MailAddress  `json:"bcc,omitempty"`
	Tags     []string       `json:"tags,omitempty"`
	Headers  map[string]any `json:"headers,omitempty"`
	Metadata map[string]any `json:"metadata,omitempty"`
	Data     map[string]any `json:"data,omitempty"`
}

// SendMessageRequestTemplate is an inline schema extracted from its parent type.
type SendMessageRequestTemplate struct {
	Subject      string         `json:"subject,omitempty"`
	Html         string         `json:"html,omitempty"`
	Text         string         `json:"text,omitempty"`
	InlineStyles bool           `json:"inlineStyles,omitempty"`
	Data         map[string]any `json:"data,omitempty"`
}

// SendMessageRequestTracking is an inline schema extracted from its parent type.
type SendMessageRequestTracking struct {
	Opens bool `json:"opens,omitempty"`
	Links bool `json:"links,omitempty"`
}

// BroadcastContentTemplate is an inline schema extracted from its parent type.
type BroadcastContentTemplate struct {
	Subject string `json:"subject"`
	Html    string `json:"html,omitempty"`
	Text    string `json:"text,omitempty"`
}

// BroadcastContentAttachment is an inline schema extracted from its parent type.
type BroadcastContentAttachment struct {
	FileName    string                `json:"fileName"`
	Disposition AttachmentDisposition `json:"disposition"`
	Size        int                   `json:"size"`
}

// StatisticsDailyResponseResult is an inline schema extracted from its parent type.
type StatisticsDailyResponseResult struct {
	Timestamp     string         `json:"timestamp,omitempty"`
	Transactional *DeliveryStats `json:"transactional,omitempty"`
	Broadcast     *DeliveryStats `json:"broadcast,omitempty"`
}

// StatisticsHourlyResponseResult is an inline schema extracted from its parent type.
type StatisticsHourlyResponseResult struct {
	Timestamp     time.Time      `json:"timestamp,omitempty"`
	Transactional *DeliveryStats `json:"transactional,omitempty"`
	Broadcast     *DeliveryStats `json:"broadcast,omitempty"`
}
