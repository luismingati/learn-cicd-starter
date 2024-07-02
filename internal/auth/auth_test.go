package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	type response struct {
		value string
		err   error
	}
	type test struct {
		input http.Header
		want  response
	}

	tests := []test{
		{
			input: http.Header{
				"Authorization": []string{"ApiKey 1234"},
			},
			want: response{
				value: "1234",
				err:   nil,
			},
		},
		{
			input: http.Header{},
			want: response{
				value: "",
				err:   ErrNoAuthHeaderIncluded,
			},
		},
		{
			input: http.Header{
				"Authorization": []string{"Bearer 1234"},
			},
			want: response{
				value: "",
				err:   ErrMalformedAuthHeader,
			},
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if got != tc.want.value {
			t.Fatalf("GetAPIKey(%v) = %v; want %v", tc.input, got, tc.want.value)
		}
		if err != tc.want.err {
			t.Fatalf("GetAPIKey(%v) = %v; want %v", tc.input, err, tc.want.err)
		}
	}
}
