package helo

import (
	"encoding/json"
	"fmt"
	"strings"
)

// marshalWithNulls encodes v, a request struct's plain type, then sets each
// field named in nulls to null. nullable lists the fields the API accepts null
// for, so a misspelled name fails instead of being sent.
func marshalWithNulls(v any, nulls []string, nullable ...string) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil || len(nulls) == 0 {
		return data, err
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return nil, err
	}
	for _, name := range nulls {
		if !containsString(nullable, name) {
			return nil, fmt.Errorf("NullFields: %q can't be null (only %s can)", name, strings.Join(nullable, ", "))
		}
		fields[name] = json.RawMessage("null")
	}
	return json.Marshal(fields)
}

func containsString(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
