// story: e01s01
// scenario: SC-e01s01-P0-01
package main

import (
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestFooterContainsVersion(t *testing.T) {
	versionBytes, err := os.ReadFile("VERSION")
	if err != nil {
		t.Fatalf("failed to read VERSION file: %v", err)
	}
	expectedVersion := strings.TrimSpace(string(versionBytes))

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, expectedVersion) {
		t.Fatalf("expected footer to contain version %s, got: %s", expectedVersion, body)
	}
}

// scenario: SC-e01s01-P1-04 (regression: deploy engine injects PORT, app must bind to it)
func TestListenAddrUsesInjectedPort(t *testing.T) {
	t.Setenv("PORT", "10004")
	if got := listenAddr(); got != ":10004" {
		t.Fatalf("expected :10004, got %s", got)
	}
}

func TestListenAddrDefaultsWhenPortUnset(t *testing.T) {
	t.Setenv("PORT", "")
	if got := listenAddr(); got != ":8080" {
		t.Fatalf("expected :8080, got %s", got)
	}
}
