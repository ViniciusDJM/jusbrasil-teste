package datasources

import (
	"errors"
	"io"
	"strings"
)

// Errors that can occur during data source handling.
var (
	ErrInvalidContentType = errors.New("the given content type of response is not supported")
	ErrInvalidStatus      = errors.New("the response status is not the same as expected")
)

// SearchFilter represents a set of parameters used for searching data.
type SearchFilter struct {
	PageNumber         int    // Page number of the search.
	ProcessCode        string // Process code used in the search.
	ProcessNumber      string // Process number used in the search.
	UnifiedYearNumber  string // Unified year number used in the search.
	UnifiedCourtNumber string // Unified court number used in the search.
}

// Writer is an interface that extends io.Writer with additional methods.
type Writer interface {
	WriteRune(r rune) (int, error)
	WriteString(string) (int, error)
	io.Writer
}

// writeParamSlice writes the key-value pairs for a slice parameter to the writer.
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

// addQueryParamsToURL adds query parameters to a given URL and returns the modified URL.
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
