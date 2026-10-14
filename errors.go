package helo

import (
	"fmt"
	"sort"
	"strings"
)

// APIError represents an error response from the API, or a request that got
// no response at all (StatusCode 0, with the transport error as Cause).
type APIError struct {
	StatusCode int
	Message    string
	Body       []byte
	Response   map[string]any
	Cause      error
}

// Error implements the error interface. It includes the API's detail or title
// (RFC 7807 problem-details), the HTTP status, error code and request ID, and
// any field-level validation errors when the response body provides them.
func (e *APIError) Error() string {
	if e.StatusCode == 0 {
		return e.Message
	}

	var b strings.Builder
	detail := e.Detail()
	b.WriteString(detail)

	var meta []string
	if detail != e.Message { // Message already names the status
		meta = append(meta, fmt.Sprintf("status %d", e.StatusCode))
	}
	if code := e.Code(); code != "" {
		meta = append(meta, "code "+code)
	}
	if id := e.RequestID(); id != "" {
		meta = append(meta, "request "+id)
	}
	if len(meta) > 0 {
		fmt.Fprintf(&b, " (%s)", strings.Join(meta, ", "))
	}

	if fieldErrs := e.fieldErrors(); fieldErrs != "" {
		fmt.Fprintf(&b, " — %s", fieldErrs)
	} else if len(e.Body) > 0 && e.Response == nil {
		// Body wasn't JSON-decodable; surface the raw bytes so the caller has
		// something to work with.
		fmt.Fprintf(&b, ": %s", strings.TrimSpace(string(e.Body)))
	}

	return b.String()
}

// Unwrap returns the underlying cause for errors.Is/As support.
func (e *APIError) Unwrap() error {
	return e.Cause
}

// Status returns the HTTP status code, or 0 when no response was received.
func (e *APIError) Status() int {
	return e.StatusCode
}

// Problem returns the decoded problem-details response body, or nil when the
// body was empty or not JSON.
func (e *APIError) Problem() map[string]any {
	return e.Response
}

// Code returns the API's machine-readable error code (e.g. "validation_failed"),
// or "" when the response didn't include one.
func (e *APIError) Code() string {
	return e.responseString("code")
}

// RequestID returns the identifier the API assigned to the failed request,
// which support can use to find it, or "" when the response didn't include one.
func (e *APIError) RequestID() string {
	return e.responseString("requestId")
}

// Detail returns the human-readable detail from the response body if present,
// falling back to title and then to the generic message.
func (e *APIError) Detail() string {
	for _, key := range []string{"detail", "title", "message"} {
		if v := e.responseString(key); v != "" {
			return v
		}
	}
	return e.Message
}

func (e *APIError) responseString(key string) string {
	if e.Response == nil {
		return ""
	}
	v, _ := e.Response[key].(string)
	return v
}

// fieldErrors formats RFC 7807-style validation errors as "field: msg; field: msg".
func (e *APIError) fieldErrors() string {
	if e.Response == nil {
		return ""
	}
	raw, ok := e.Response["errors"].(map[string]any)
	if !ok || len(raw) == 0 {
		return ""
	}
	keys := make([]string, 0, len(raw))
	for k := range raw {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s: %s", k, formatErrorValue(raw[k])))
	}
	return strings.Join(parts, "; ")
}

// formatErrorValue renders one field's errors. The API sends each as
// {"message": ..., "code": ...}, shown as "message (code)".
func formatErrorValue(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case map[string]any:
		message, _ := t["message"].(string)
		code, _ := t["code"].(string)
		switch {
		case message != "" && code != "":
			return fmt.Sprintf("%s (%s)", message, code)
		case message != "":
			return message
		case code != "":
			return code
		}
		return fmt.Sprintf("%v", t)
	case []any:
		strs := make([]string, 0, len(t))
		for _, item := range t {
			strs = append(strs, formatErrorValue(item))
		}
		return strings.Join(strs, ", ")
	default:
		return fmt.Sprintf("%v", t)
	}
}
