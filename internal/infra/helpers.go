package infra

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
)

//go:generate mockgen -destination=../mocks/datasources_mock.go -package=mocks github.com/ViniciusDJM/jusbrasil-teste/internal/infra RequestDatasource

type RequestDatasource interface {
	SearchFirstInstance(filter datasources.SearchFilter) ([]byte, error)
	SearchSecondInstance(filter datasources.SearchFilter) ([]byte, error)
}

func NodeToStringSlice(originalNode *html.Node) (result string) {
	var builder strings.Builder
	for node := originalNode; node != nil; node = node.NextSibling {
		switch node.DataAtom {
		case atom.Br:
			builder.WriteRune('\n')
		case 0:
			builder.WriteString(strings.TrimSpace(node.Data))
		default:
			builder.WriteString(NodeToStringSlice(node.FirstChild))
		}
	}

	return builder.String()
}
