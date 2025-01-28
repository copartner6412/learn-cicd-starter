package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKeyFailsForInvalidInput(t *testing.T) {
	testCases := map[string]http.Header{
		"nil_header":  nil,
		"without_key": {},
		"without_authorization_key": {
			"another": []string{"another"},
		},
		"nil_value": {
			"Authorization": nil,
		},
		"no_value": {
			"Authorization": []string{},
		},
		"value_one_empty": {
			"Authorization": []string{
				"",
			},
		},
		"value_one_part": {
			"Authorization": []string{
				"ApiKey",
			},
		},
		"value_wrong_first_part": {
			"Authorization": []string{
				"apiKey token",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			if _, err := auth.GetAPIKey(testCase); err == nil {
				t.Error(err)
			}
		})

	}
}
