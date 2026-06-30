package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		// name        string
		// authHeader  string
		// expectedKey string
		// expectedErr error
	}{
		{
			name:        "valid API key",
			authHeader:  "ApiKey abc123",
			expectedKey: "abc123",
			expectedErr: nil,
		},
		{
			name:        "missing authorization header",
			authHeader:  "",
			expectedKey: "",
			expectedErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:        "incorrect authorization scheme",
			authHeader:  "Bearer abc123",
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
		{
			name:        "missing API key",
			authHeader:  "ApiKey",
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
		{
			name:        "authorization header without scheme",
			authHeader:  "abc123",
			expectedKey: "",
			expectedErr: errors.New("malformed authorization header"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			headers := http.Header{}

			if test.authHeader != "" {
				headers.Set("Authorization", test.authHeader)
			}

			actualKey, actualErr := GetAPIKey(headers)

			if actualKey != test.expectedKey {
				t.Errorf(
					"expected API key %q, got %q",
					test.expectedKey,
					actualKey,
				)
			}

			if test.expectedErr == nil {
				if actualErr != nil {
					t.Errorf("expected no error, got %v", actualErr)
				}
				return
			}

			if actualErr == nil {
				t.Fatalf("expected error %q, got nil", test.expectedErr.Error())
			}

			if actualErr.Error() != test.expectedErr.Error() {
				t.Errorf(
					"expected error %q, got %q",
					test.expectedErr.Error(),
					actualErr.Error(),
				)
			}
		})
	}
}
