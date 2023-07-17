package datasources

import (
	"errors"
	"io"
	"strings"
)

var (
	ErrInvalidContentType = errors.New("the given content type of response is not supported")
	ErrInvalidStatus      = errors.New("the response status is not the same as expected")
)

type SearchFilter struct {
	PageNumber         int
	ProcessCode        string
	ProcessNumber      string
	UnifiedYearNumber  string
	UnifiedCourtNumber string
}

type Writer interface {
	WriteRune(r rune) (int, error)
	WriteString(string) (int, error)
	io.Writer
}

func writeParamSlice(writer Writer, key string, slice []string) {
	for index, element := range slice {
		if index > 0 {
			writer.WriteRune('&')
		}
		writer.WriteString(key)
		writer.WriteRune('=')
		writer.WriteString(element)
	}
}

func addQueryParamsToURL(url string, queryParams map[string]any) string {
	var builder strings.Builder
	builder.WriteString(url)

	var index uint16
	for key, param := range queryParams {
		if index == 0 {
			builder.WriteRune('?')
		} else if index > 0 {
			builder.WriteRune('&')
		}

		switch param.(type) {
		case string:
			builder.WriteString(key)
			if paramStr := param.(string); paramStr != "" {
				builder.WriteRune('=')
				builder.WriteString(paramStr)
			}
		case []string:
			writeParamSlice(&builder, key, param.([]string))
		}
		index += 1
	}

	return builder.String()
}
