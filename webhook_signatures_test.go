package helo

import (
	"errors"
	"strconv"
	"testing"
	"time"
)

const (
	testSigningKey  = "whsec_test"
	testWebhookBody = `{"event":"message.delivered"}`
)

func currentWebhookTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func TestVerifyWebhookSignatureAcceptsValidHeader(t *testing.T) {
	timestamp := currentWebhookTimestamp()
	signature := WebhookSignature([]byte(testWebhookBody), testSigningKey, timestamp)
	header := "t=" + timestamp + ",v1=" + signature

	if err := VerifyWebhookSignature(header, []byte(testWebhookBody), testSigningKey); err != nil {
		t.Fatalf("expected signature %q to be valid, got %v", header, err)
	}
}

// A sender rolling out a new signing scheme emits every version at once. This
// SDK must keep verifying the versions it knows and ignore the rest, otherwise
// the rollout breaks every receiver that has not upgraded yet.
func TestVerifyWebhookSignatureIgnoresUnknownVersionsAndElements(t *testing.T) {
	timestamp := currentWebhookTimestamp()
	signature := WebhookSignature([]byte(testWebhookBody), testSigningKey, timestamp)

	headers := []string{
		"t=" + timestamp + ",v1=" + signature + ",v2=8badf00d",
		"t=" + timestamp + ",v2=8badf00d,v1=" + signature,
		"t=" + timestamp + ",v1=" + signature + ",alg=sha512",
		"t=" + timestamp + ", v1=" + signature,
		"v1=" + signature + ",t=" + timestamp,
		"t=" + timestamp + ",v1=8badf00d,v1=" + signature,
	}

	for _, header := range headers {
		t.Run(header, func(t *testing.T) {
			if err := VerifyWebhookSignature(header, []byte(testWebhookBody), testSigningKey); err != nil {
				t.Fatalf("expected signature %q to be valid, got %v", header, err)
			}
		})
	}
}

func TestVerifyWebhookSignatureRejectsInvalidHeaders(t *testing.T) {
	timestamp := currentWebhookTimestamp()
	signature := WebhookSignature([]byte(testWebhookBody), testSigningKey, timestamp)

	staleTimestamp := strconv.FormatInt(time.Now().Add(-10*time.Minute).Unix(), 10)
	staleSignature := WebhookSignature([]byte(testWebhookBody), testSigningKey, staleTimestamp)

	cases := []struct {
		name    string
		header  string
		body    string
		key     string
		wantErr error
	}{
		{"wrong key", "t=" + timestamp + ",v1=" + signature, testWebhookBody, "wrong-key", ErrSignatureMismatch},
		{"tampered body", "t=" + timestamp + ",v1=" + signature, `{"event":"message.bounced"}`, testSigningKey, ErrSignatureMismatch},
		{"stale timestamp", "t=" + staleTimestamp + ",v1=" + staleSignature, testWebhookBody, testSigningKey, ErrSignatureTimestampSkew},
		{"only unknown versions", "t=" + timestamp + ",v2=" + signature, testWebhookBody, testSigningKey, ErrUnsupportedSignatureVersion},
		{"unknown version with valid v1 for another key", "t=" + timestamp + ",v2=" + signature + ",v1=8badf00d", testWebhookBody, testSigningKey, ErrSignatureMismatch},
		{"malformed header", "garbage", testWebhookBody, testSigningKey, ErrMalformedSignatureHeader},
		{"uppercase signature", "t=" + timestamp + ",v1=ABCDEF", testWebhookBody, testSigningKey, ErrMalformedSignatureHeader},
		{"short signature", "t=" + timestamp + ",v1=abc", testWebhookBody, testSigningKey, ErrSignatureMismatch},
		{"no signature element", "t=" + timestamp, testWebhookBody, testSigningKey, ErrMalformedSignatureHeader},
		{"no timestamp element", "v1=" + signature, testWebhookBody, testSigningKey, ErrMalformedSignatureHeader},
		{"non-numeric timestamp", "t=yesterday,v1=" + signature, testWebhookBody, testSigningKey, ErrMalformedSignatureHeader},
		{"empty header", "", testWebhookBody, testSigningKey, ErrMalformedSignatureHeader},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := VerifyWebhookSignature(tc.header, []byte(tc.body), tc.key)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("expected %v for %q, got %v", tc.wantErr, tc.header, err)
			}
		})
	}
}
