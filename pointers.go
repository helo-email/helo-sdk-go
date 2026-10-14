package helo

// Bool returns a pointer to v, for the optional boolean fields of request
// types. They are pointers so that an explicit false is sent rather than
// omitted, e.g. Enabled: Bool(false).
func Bool(v bool) *bool {
	return &v
}
