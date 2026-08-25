package helo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Errors returned by VerifyWebhookSignature. Compare them with errors.Is.
var (
	// ErrMalformedSignatureHeader means the header was not in the documented
	// t={timestamp},v{version}={signature} form.
	ErrMalformedSignatureHeader = errors.New("helo: malformed webhook signature header")

	// ErrUnsupportedSignatureVersion means the header carried only signing
	// schemes this SDK does not know how to verify. Upgrading the SDK is the
	// fix; see SupportedWebhookSignatureVersions.
	ErrUnsupportedSignatureVersion = errors.New("helo: unsupported webhook signature version")

	// ErrSignatureTimestampSkew means the signature was correctly formed but
	// its timestamp is too far from the current time, so it may be a replay.
	ErrSignatureTimestampSkew = errors.New("helo: webhook signature timestamp outside tolerance")

	// ErrSignatureMismatch means the signature did not match the body, either
	// because the body was tampered with or the signing key is wrong.
	ErrSignatureMismatch = errors.New("helo: webhook signature mismatch")
)

// SupportedWebhookSignatureVersions lists the signing schemes this SDK can
// verify. The signature header may carry several versions at once
// (t=...,v1=...,v2=...) so that a new scheme can be rolled out while receivers
// upgrade; verification uses the newest version present that appears in this
// list, and ignores the rest.
var SupportedWebhookSignatureVersions = []int{1}

var (
	timestampValueRegex = regexp.MustCompile(`^\d+$`)
	signatureKeyRegex   = regexp.MustCompile(`^v(\d+)$`)
	hexSignatureRegex   = regexp.MustCompile(`^[a-f0-9]+$`)
)

// maxTimestampSkew is how far a webhook's timestamp may drift from the current
// time before its signature is rejected.
const maxTimestampSkew = 5 * time.Minute

// VerifyWebhookSignature checks signatureHeader against requestBody, which must
// be the raw (unparsed) request body. It returns nil when the signature is
// valid, and otherwise one of the Err* values above describing why it was
// rejected.
func VerifyWebhookSignature(signatureHeader string, requestBody []byte, signingKey string) error {
	timestamp, signatures, err := parseWebhookSignatureHeader(signatureHeader)
	if err != nil {
		return err
	}

	version, ok := newestSupportedSignatureVersion(signatures)
	if !ok {
		return fmt.Errorf("%w: header carries only %s", ErrUnsupportedSignatureVersion, describeSignatureVersions(signatures))
	}

	seconds, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return ErrMalformedSignatureHeader
	}

	if skew := absDuration(time.Since(time.Unix(seconds, 0))); skew > maxTimestampSkew {
		return fmt.Errorf("%w: off by %s, tolerance is %s", ErrSignatureTimestampSkew, skew.Round(time.Second), maxTimestampSkew)
	}

	computedSignature, ok := signatureForVersion(version, requestBody, signingKey, timestamp)
	if !ok {
		return ErrUnsupportedSignatureVersion
	}

	for _, signature := range signatures[version] {
		if hmac.Equal([]byte(computedSignature), []byte(signature)) {
			return nil
		}
	}

	return ErrSignatureMismatch
}

// WebhookSignature returns the hex-encoded HMAC-SHA256 signature for a webhook
// payload, using the v1 signing scheme.
func WebhookSignature(payload []byte, key string, timestamp string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(timestamp + "."))
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil))
}

// signatureForVersion computes the signature for one signing scheme, and is the
// single place a new scheme needs to be added.
func signatureForVersion(version int, payload []byte, key string, timestamp string) (string, bool) {
	switch version {
	case 1:
		return WebhookSignature(payload, key, timestamp), true
	default:
		return "", false
	}
}

// parseWebhookSignatureHeader splits the header into its timestamp and its
// signatures keyed by version. Elements it does not recognize are ignored, so
// that a sender adding new elements does not break verification here.
func parseWebhookSignatureHeader(header string) (string, map[int][]string, error) {
	timestamp := ""
	signatures := map[int][]string{}

	for _, element := range strings.Split(header, ",") {
		key, value, found := strings.Cut(strings.TrimSpace(element), "=")
		if !found {
			continue
		}

		if key == "t" {
			if !timestampValueRegex.MatchString(value) {
				return "", nil, ErrMalformedSignatureHeader
			}
			timestamp = value
			continue
		}

		match := signatureKeyRegex.FindStringSubmatch(key)
		if match == nil {
			continue
		}

		version, err := strconv.Atoi(match[1])
		if err != nil {
			continue
		}

		// Only versions this SDK verifies have a signature format it can
		// insist on; anything else is recorded but left unchecked.
		if isSupportedSignatureVersion(version) && !hexSignatureRegex.MatchString(value) {
			return "", nil, ErrMalformedSignatureHeader
		}

		signatures[version] = append(signatures[version], value)
	}

	if timestamp == "" || len(signatures) == 0 {
		return "", nil, ErrMalformedSignatureHeader
	}

	return timestamp, signatures, nil
}

func isSupportedSignatureVersion(version int) bool {
	for _, supported := range SupportedWebhookSignatureVersions {
		if supported == version {
			return true
		}
	}
	return false
}

// newestSupportedSignatureVersion picks the highest version present that this
// SDK can verify, so that once a sender emits a newer scheme the older one
// stops being honored here.
func newestSupportedSignatureVersion(signatures map[int][]string) (int, bool) {
	newest, found := 0, false
	for version := range signatures {
		if isSupportedSignatureVersion(version) && (!found || version > newest) {
			newest, found = version, true
		}
	}
	return newest, found
}

func describeSignatureVersions(signatures map[int][]string) string {
	versions := make([]int, 0, len(signatures))
	for version := range signatures {
		versions = append(versions, version)
	}
	sort.Ints(versions)

	labels := make([]string, 0, len(versions))
	for _, version := range versions {
		labels = append(labels, "v"+strconv.Itoa(version))
	}
	return strings.Join(labels, ", ")
}

func absDuration(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}
