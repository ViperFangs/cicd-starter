package auth

import (
	"errors"
	"net/http"
	"testing"
)

type testCase struct {
	name           string
	headers        http.Header
	expectedAPIKey string
	expectedErr    error
}

func TestGetAPIKey(t *testing.T) {
	tests := []testCase{
		{
			name:           "valid header",
			headers:        http.Header{"Authorization": []string{"ApiKey api-key"}},
			expectedAPIKey: "api-key",
			expectedErr:    nil,
		},
		{
			name:           "missing authorization header",
			headers:        http.Header{},
			expectedAPIKey: "",
			expectedErr:    ErrNoAuthHeaderIncluded,
		},
	}

	// Iterate over the test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.headers)
			if apiKey != tc.expectedAPIKey {
				t.Errorf("Expected API key %s, got %s", tc.expectedAPIKey, apiKey)
			}
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
