package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectError bool
	}{
		{
			name: "valid api key",
			headers: http.Header{
				"Authorization": []string{"ApiKey secret-token-123"},
			},
			expectedKey: "secret-token-123",
			expectError: false,
		},
		{
			name:        "missing authorization header",
			headers:     http.Header{},
			expectedKey: "",
			expectError: true,
		},
		{
			name: "malformed authorization header without prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer some-token"},
			},
			expectedKey: "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualKey, err := GetAPIKey(tt.headers)

			if (err != nil) != tt.expectError {
				t.Fatalf("expected error: %v, got: %v", tt.expectError, err)
			}

			if actualKey != tt.expectedKey {
				t.Fatalf("expected key: %s, got: %s", tt.expectedKey, actualKey)
			}
		})
	}
}
