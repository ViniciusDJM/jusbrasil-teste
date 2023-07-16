package utils

import (
	"strings"
)

func ParseMonetaryValue(input, prefix string) (result float64) {
	sanitizedElement := strings.TrimSpace(strings.TrimPrefix(input, prefix))

	var divideBy float64
	for _, character := range sanitizedElement {
		if character == ',' {
			divideBy = 10
		}

		if character >= '0' && character <= '9' {
			switch {
			case divideBy > 0:
				result += float64(character-'0') / divideBy
				divideBy *= 10
			default:
				result = (result * 10) + float64(character-'0')
			}
		}
	}

	result = float64(uint64(result*100)) / 100
	return
}
