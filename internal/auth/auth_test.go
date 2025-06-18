package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

// func TestAPIKey(t *testing.T) {
// 	apiKey := "alex!9657"

// 	header := http.Header{
// 		"Authorization": {fmt.Sprintf("ApiKey %s", apiKey)},
// 	}

// 	want := apiKey
// 	got, err := GetAPIKey(header)
// 	if err != nil {
// 		t.Error("error getting API key", err)
// 	}
// 	if want != got {
// 		t.Error("incorrect API key", err)
// 	}

// }

// func TestAPIKeyIncorrect(t *testing.T) {
// 	apiKey := "alex!9657"

// 	header := http.Header{
// 		"Authorization": {fmt.Sprintf("ApiKey %s", apiKey)},
// 	}

// 	want := "alex&9657"
// 	got, err := GetAPIKey(header)
// 	if err != nil {
// 		t.Error("error getting API key", err)
// 	}
// 	if want != got {
// 		return
// 	}

// }

// func TestAPIKeyNoHeader(t *testing.T) {
// 	apiKey := "alex!9657"

// 	header := http.Header{
// 		"Content-Type": {"application-json"},
// 	}

// 	want := apiKey
// 	got, err := GetAPIKey(header)
// 	if err != nil {
// 		return
// 	}
// 	if want != got {
// 		t.Error("incorrect API key")
// 	}

// }

// func TestAPIKeyMalformedHeader(t *testing.T) {
// 	apiKey := "alex!9657"

// 	header := http.Header{
// 		"Authorization": {fmt.Sprintf("ApiKey%s", apiKey)},
// 	}

// 	want := apiKey
// 	got, err := GetAPIKey(header)
// 	if err != nil {
// 		return
// 	}
// 	if want != got {
// 		t.Error("incorrect API key")
// 	}

// }

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		{
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			expectErr: "no authorization header",
		},
		{
			key:       "Authorization",
			value:     "-",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "Bearer xxxxxx",
			expectErr: "malformed authorization header",
		},
		{
			key:       "Authorization",
			value:     "ApiKey xxxxxx",
			expect:    "xxxxxx",
			expectErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey Case #%v:", i), func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}

			if output != test.expect {
				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
				return
			}
		})
	}
}
