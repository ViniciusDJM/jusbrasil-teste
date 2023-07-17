package datasources

import "testing"

func TestAddQueryParamsToURL(t *testing.T) {
	testCases := []struct {
		URL         string
		QueryParams map[string]any
		ExpectedURL string
	}{
		{
			URL: "https://example.com",
			QueryParams: map[string]any{
				"key1": "value1",
				"key2": []string{"value2", "value3"},
			},
			ExpectedURL: "https://example.com?key1=value1&key2=value2&key2=value3",
		},
		{
			URL: "https://example.com",
			QueryParams: map[string]any{
				"key1": "value1",
			},
			ExpectedURL: "https://example.com?key1=value1",
		},
		{
			URL:         "https://example.com",
			QueryParams: map[string]any{},
			ExpectedURL: "https://example.com",
		},
	}

	for _, tc := range testCases {
		result := addQueryParamsToURL(tc.URL, tc.QueryParams)
		if result != tc.ExpectedURL {
			t.Errorf(
				"URL mismatch for input: %s\nExpected: %s\nGot: %s",
				tc.URL, tc.ExpectedURL, result,
			)
		}
	}
}
