package utils

import (
	"testing"
)

func TestParseMonetaryValue(t *testing.T) {
	testCases := []struct {
		input  string
		result float64
	}{
		{
			input:  "R$ 281.178,42",
			result: 281178.42,
		},
		{
			input:  "R$ 100,00",
			result: 100.00,
		},
		{
			input:  "R$ 1.000.000,00",
			result: 1000000.00,
		},
		{
			input:  "R$ 5,50",
			result: 5.50,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.input, func(t *testing.T) {
			result := ParseMonetaryValue(tCase.input, "R$")

			if result != tCase.result {
				t.Error("Values are not equal")
			}
		})
	}
}
