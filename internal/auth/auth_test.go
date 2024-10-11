package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		want       string
		wantErr    error
	}{
		{
			name:    "No Authorization Header",
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		// {
		// 	name:    "Malformed Authorization Header",
		// 	headers: http.Header{"Authorization": []string{"Bearer token"}},
		// 	want:    "",
		// 	wantErr: errors.New("malformed authorization header"),
		// },
		{
			name:    "Valid Authorization Header",
			headers: http.Header{"Authorization": []string{"ApiKey valid_api_key"}},
			want:    "valid_api_key",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Errorf("GetAPIKey() = %v, %v; want %v, %v", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

