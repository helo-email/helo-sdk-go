package helo

import (
	"fmt"
	"sort"
	"strings"
)

// APIError represents an error response from the API.
type APIError struct {
	StatusCode int
	Message    string
	Body       []byte
	Response   map[string]any
	Cause      error
}

// Error implements the error interface. It includes the HTTP status,
// the API's detail or title (RFC 7807 problem-details), and any field-level
// validation errors when the response body provides them.
func (e *APIError) Error() string {
	if e.StatusCode == 0 {
		return e.Message
	}

	var b strings.Builder
	fmt.Fprintf(&b, "%s (status %d)", e.Message, e.StatusCode)

	if detail := e.Detail(); detail != "" && detail != e.Message {
		fmt.Fprintf(&b, ": %s", detail)
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

// Detail returns the human-readable detail from the response body if present,
// falling back to title and then to the generic message.
func (e *APIError) Detail() string {
	if e.Response != nil {
		if v, ok := e.Response["detail"].(string); ok && v != "" {
			return v
		}
		if v, ok := e.Response["title"].(string); ok && v != "" {
			return v
		}
		if v, ok := e.Response["message"].(string); ok && v != "" {
			return v
		}
	}
	return e.Message
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

func formatErrorValue(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case []any:
		strs := make([]string, 0, len(t))
		for _, item := range t {
			strs = append(strs, fmt.Sprintf("%v", item))
		}
		return strings.Join(strs, ", ")
	default:
		return fmt.Sprintf("%v", t)
	}
}
