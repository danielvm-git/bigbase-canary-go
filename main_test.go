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

func TestHandlerSetsContentType(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	ct := rec.Header().Get("Content-Type")
	if ct != "text/html; charset=utf-8" {
		t.Fatalf("expected Content-Type text/html; charset=utf-8, got %s", ct)
	}
}

func TestHandlerMissingVersionFile(t *testing.T) {
	dir := t.TempDir()
	orig, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("failed to chdir to temp dir: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(orig); err != nil {
			t.Fatalf("failed to restore working directory: %v", err)
		}
	})

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, "vunknown") {
		t.Fatalf("expected vunknown in body when VERSION missing, got: %s", body)
	}
	if rec.Code != 200 {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}
