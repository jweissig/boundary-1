package schema

import "testing"

const NilVersion = nilVersion

// TestClearEditions is a test helper to reset the editions map.
func TestClearEditions(t *testing.T) {
	t.Helper()
	editions = make(dialects)
}
