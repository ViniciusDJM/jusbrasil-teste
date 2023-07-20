package infra

import (
	"strings"

	"golang.org/x/exp/slices"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
)

//go:generate mockgen -destination=../mocks/datasources_mock.go -package=mocks github.com/ViniciusDJM/jusbrasil-teste/internal/infra RequestDatasource

type RequestDatasource interface {
	SearchFirstInstance(filter datasources.SearchFilter) ([]byte, error)
	SearchSecondInstance(filter datasources.SearchFilter) ([]byte, error)
}

// NodeToStringSlice converts an HTML node and its children to a string.
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

// insertProcessPart is a helper function to insert process participants based on the label.
func insertProcessPart(label string, toInsert entities.ProcessPeople, process *entities.ProcessParts) {
	switch {
	case slices.Contains([]string{"autor", "autora"}, label):
		if !slices.Contains(process.Author, toInsert) {
			process.Author = append(process.Author, toInsert)
		}
	case slices.Contains([]string{"ré", "réu"}, label):
		if !slices.Contains(process.Defendant, toInsert) {
			process.Defendant = append(process.Defendant, toInsert)
		}
	case slices.Contains([]string{"apelante"}, label):
		if !slices.Contains(process.Appellant, toInsert) {
			process.Appellant = append(process.Appellant, toInsert)
		}
	case slices.Contains([]string{"apelado", "apelada"}, label):
		if !slices.Contains(process.Appellee, toInsert) {
			process.Appellee = append(process.Appellee, toInsert)
		}
	case slices.Contains([]string{"vítima"}, label):
		if !slices.Contains(process.Victim, toInsert.Name) {
			process.Victim = append(process.Victim, toInsert.Name)
		}
	case slices.Contains([]string{"terceiro"}, label):
		if !slices.Contains(process.Third, toInsert.Name) {
			process.Third = append(process.Third, toInsert.Name)
		}
	case slices.Contains([]string{"testemunha"}, label):
		if !slices.Contains(process.Witness, toInsert.Name) {
			process.Witness = append(process.Witness, toInsert.Name)
		}
	}
}
