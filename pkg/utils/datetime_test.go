package utils

import (
	"testing"
	"time"
)

// TestParseDistributionDate is a test function that tests the ParseBRTDateTime function.
func TestParseDistributionDate(t *testing.T) {
	testCases := []struct {
		input  string
		result time.Time
	}{
		{
			input:  "15/07/2023 às 12:34",
			result: time.Date(2023, time.July, 15, 12, 34, 0, 0, time.UTC),
		},
		{
			input:  "05/11/2022  09:15",
			result: time.Date(2022, time.November, 5, 9, 15, 0, 0, time.UTC),
		},
		{
			input:  "30/01/2024 _&%¨#* 18:45",
			result: time.Date(2024, time.January, 30, 18, 45, 0, 0, time.UTC),
		},
		{
			input:  "10/09/2025 :;:://::;: 23:59",
			result: time.Date(2025, time.September, 10, 23, 59, 0, 0, time.UTC),
		},
		{
			input:  "100920252359",
			result: time.Date(2025, time.September, 10, 23, 59, 0, 0, time.UTC),
		},
	}

	// Loop through each test case and run the test
	for _, tCase := range testCases {
		t.Run(tCase.input, func(t *testing.T) {
			// Call the ParseBRTDateTime function to parse the input string
			result, err := ParseBRTDateTime(tCase.input)
			if err != nil {
				t.Fatal(err)
			}

			// Check if the parsed time matches the expected result
			if result != tCase.result {
				t.Error("Time are not equal")
			}
		})
	}
}
