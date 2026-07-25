// story: e01s01
// scenario: SC-e01s01-P0-01
package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFooterContainsVersion(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, "0.1.0") {
		t.Fatalf("expected footer to contain version 0.1.0, got: %s", body)
	}
}
