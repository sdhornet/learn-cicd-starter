package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKeyValid(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey secret-key")

	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("GetAPIKey returned error: %v", err)
	}
	if got != "secret-key" {
		t.Fatalf("GetAPIKey = %q, want %q", got, "secret-key")
	}
}

func TestGetAPIKeyMissingHeader(t *testing.T) {
	got, err := GetAPIKey(http.Header{})
	if got != "" {
		t.Fatalf("GetAPIKey = %q, want empty string", got)
	}
	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("GetAPIKey error = %v, want %v", err, ErrNoAuthHeaderIncluded)
	}
}
