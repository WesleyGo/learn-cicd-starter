package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-api-key")

	apiKey, err := GetAPIKey(headers)
	if err == nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedAPIKey := "my-api-key"
	if apiKey != expectedAPIKey {
		t.Errorf("expected API key %q, got %q", expectedAPIKey, apiKey)
	}
}
