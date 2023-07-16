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

func ParseBRTDateTime(element string) (time.Time, error) {
	dateHourPart, _, _ := strings.Cut(element, " - ")
	delimiters := [...]rune{'/', '/', ' ', ':'}

	// 02/05/2018 às 19:01
	var (
		currentNumberSize uint8
		delimiterIndex    uint8
		builder           strings.Builder
	)
	for _, character := range dateHourPart {
		if character >= '0' && character <= '9' {
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

			builder.WriteRune(character)
			currentNumberSize += 1
		}
	}

	return time.Parse(FormatDateTimeBR, builder.String())
}
