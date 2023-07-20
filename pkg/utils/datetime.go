package utils

import (
	"strings"
	"time"
)

const (
	FormatDateBR     = "02/01/2006"
	FormatTimeBR     = "15:04"
	FormatDateTimeBR = FormatDateBR + " " + FormatTimeBR
)

func ParseBRTDate(element string) (time.Time, error) {
	return time.Parse(FormatDateBR, element)
}

// ParseBRTDateTime parses a date-time string in Brazilian format and returns a time.Time value.
func ParseBRTDateTime(element string) (time.Time, error) {
	// Extract the date and time part from the input string
	dateHourPart, _, _ := strings.Cut(element, " - ")
	delimiters := [...]rune{'/', '/', ' ', ':'}

	var (
		currentNumberSize uint8
		delimiterIndex    uint8
		builder           strings.Builder
	)
	// Iterate through the characters in the dateHourPart
	for _, character := range dateHourPart {
		// Check if the character is a digit (0-9)
		if character >= '0' && character <= '9' {
			// Handle delimiters based on the position of the digit
			if delimiterIndex <= 1 && currentNumberSize >= 2 {
				builder.WriteRune(delimiters[delimiterIndex])
				currentNumberSize = 0
				delimiterIndex += 1
			}

			if delimiterIndex > 1 {
				if currentNumberSize == 4 {
					builder.WriteRune(delimiters[delimiterIndex])
					currentNumberSize = 0
					delimiterIndex += 1
				} else if delimiterIndex == 3 && currentNumberSize == 2 {
					builder.WriteRune(delimiters[delimiterIndex])
				}
			}

			// Append the digit to the builder
			builder.WriteRune(character)
			currentNumberSize += 1
		}
	}

	// Parse the formatted date-time string using the defined format
	return time.Parse(FormatDateTimeBR, builder.String())
}
