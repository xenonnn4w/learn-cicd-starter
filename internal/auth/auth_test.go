package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	goodHeaders := http.Header{}
	goodHeaders.Set("Authorization", "ApiKey 1234567890")
	badHeaders := http.Header{}
	badHeaders.Set("Authorization", "notapi 1234567890")

	apiKey, err := GetAPIKey(goodHeaders)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if apiKey != "1234567890" {
		t.Fatalf("expected api key to be 1234567890, got %s", apiKey)
	}

	apiKey, err = GetAPIKey(badHeaders)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if apiKey != "" {
		t.Fatalf("expected api key to be empty, got %s", apiKey)
	}
}
