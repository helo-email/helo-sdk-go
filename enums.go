package helo

// MailType is a string-based enum.
type MailType string

const (
	MailTypeTransactional MailType = "transactional"
	MailTypeBroadcast     MailType = "broadcast"
)

// AllMailTypeValues returns all valid MailType values.
func AllMailTypeValues() []MailType {
	return []MailType{
		MailTypeTransactional,
		MailTypeBroadcast,
	}
}

// DeliveryType is a string-based enum.
type DeliveryType string

const (
	DeliveryTypeLive    DeliveryType = "live"
	DeliveryTypeSandbox DeliveryType = "sandbox"
)

// AllDeliveryTypeValues returns all valid DeliveryType values.
func AllDeliveryTypeValues() []DeliveryType {
	return []DeliveryType{
		DeliveryTypeLive,
		DeliveryTypeSandbox,
	}
}

// EventType is a string-based enum.
type EventType string

const (
	EventTypeAccepted     EventType = "accepted"
	EventTypeProcessed    EventType = "processed"
	EventTypeDelivered    EventType = "delivered"
	EventTypeBounced      EventType = "bounced"
	EventTypeOpened       EventType = "opened"
	EventTypeClicked      EventType = "clicked"
	EventTypeComplained   EventType = "complained"
	EventTypeUnsubscribed EventType = "unsubscribed"
	EventTypeResubscribed EventType = "resubscribed"
)

// AllEventTypeValues returns all valid EventType values.
func AllEventTypeValues() []EventType {
	return []EventType{
		EventTypeAccepted,
		EventTypeProcessed,
		EventTypeDelivered,
		EventTypeBounced,
		EventTypeOpened,
		EventTypeClicked,
		EventTypeComplained,
		EventTypeUnsubscribed,
		EventTypeResubscribed,
	}
}

// DnsRecordStatus is a string-based enum.
type DnsRecordStatus string

const (
	DnsRecordStatusPending  DnsRecordStatus = "pending"
	DnsRecordStatusVerified DnsRecordStatus = "verified"
	DnsRecordStatusFailing  DnsRecordStatus = "failing"
	DnsRecordStatusFailed   DnsRecordStatus = "failed"
)

// AllDnsRecordStatusValues returns all valid DnsRecordStatus values.
func AllDnsRecordStatusValues() []DnsRecordStatus {
	return []DnsRecordStatus{
		DnsRecordStatusPending,
		DnsRecordStatusVerified,
		DnsRecordStatusFailing,
		DnsRecordStatusFailed,
	}
}

// DnsRecordType is a string-based enum.
type DnsRecordType string

const (
	DnsRecordTypeTxt   DnsRecordType = "txt"
	DnsRecordTypeCname DnsRecordType = "cname"
)

// AllDnsRecordTypeValues returns all valid DnsRecordType values.
func AllDnsRecordTypeValues() []DnsRecordType {
	return []DnsRecordType{
		DnsRecordTypeTxt,
		DnsRecordTypeCname,
	}
}

// BroadcastStatus is a string-based enum.
type BroadcastStatus string

const (
	BroadcastStatusAccepted   BroadcastStatus = "accepted"
	BroadcastStatusProcessing BroadcastStatus = "processing"
	BroadcastStatusCompleted  BroadcastStatus = "completed"
	BroadcastStatusCancelled  BroadcastStatus = "cancelled"
)

// AllBroadcastStatusValues returns all valid BroadcastStatus values.
func AllBroadcastStatusValues() []BroadcastStatus {
	return []BroadcastStatus{
		BroadcastStatusAccepted,
		BroadcastStatusProcessing,
		BroadcastStatusCompleted,
		BroadcastStatusCancelled,
	}
}

// AttachmentDisposition is a string-based enum.
type AttachmentDisposition string

const (
	AttachmentDispositionAttachment AttachmentDisposition = "attachment"
	AttachmentDispositionInline     AttachmentDisposition = "inline"
)

// AllAttachmentDispositionValues returns all valid AttachmentDisposition values.
func AllAttachmentDispositionValues() []AttachmentDisposition {
	return []AttachmentDisposition{
		AttachmentDispositionAttachment,
		AttachmentDispositionInline,
	}
}

// SuppressionReason is a string-based enum.
type SuppressionReason string

const (
	SuppressionReasonBounce      SuppressionReason = "bounce"
	SuppressionReasonComplaint   SuppressionReason = "complaint"
	SuppressionReasonUnsubscribe SuppressionReason = "unsubscribe"
	SuppressionReasonManual      SuppressionReason = "manual"
)

// AllSuppressionReasonValues returns all valid SuppressionReason values.
func AllSuppressionReasonValues() []SuppressionReason {
	return []SuppressionReason{
		SuppressionReasonBounce,
		SuppressionReasonComplaint,
		SuppressionReasonUnsubscribe,
		SuppressionReasonManual,
	}
}

// WebhookEvent is a string-based enum.
type WebhookEvent string

const (
	WebhookEventAccepted     WebhookEvent = "accepted"
	WebhookEventProcessed    WebhookEvent = "processed"
	WebhookEventBounced      WebhookEvent = "bounced"
	WebhookEventDelivered    WebhookEvent = "delivered"
	WebhookEventOpened       WebhookEvent = "opened"
	WebhookEventClicked      WebhookEvent = "clicked"
	WebhookEventComplained   WebhookEvent = "complained"
	WebhookEventUnsubscribed WebhookEvent = "unsubscribed"
	WebhookEventResubscribed WebhookEvent = "resubscribed"
)

// AllWebhookEventValues returns all valid WebhookEvent values.
func AllWebhookEventValues() []WebhookEvent {
	return []WebhookEvent{
		WebhookEventAccepted,
		WebhookEventProcessed,
		WebhookEventBounced,
		WebhookEventDelivered,
		WebhookEventOpened,
		WebhookEventClicked,
		WebhookEventComplained,
		WebhookEventUnsubscribed,
		WebhookEventResubscribed,
	}
}
