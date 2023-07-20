package entities

import (
	"strconv"
	"strings"
)

// CNJ is the process number that follows the pattern: NNNNNNN-DD.AAAA.J.TR.OOOO
type CNJ struct {
	number string
	// year AAAA
	year string
	// segment J
	segment Segmento
	// court TR
	court string
	// noticeOrderNumber NNNNNNN
	noticeOrderNumber string
	// sourcerUnitID OOOO
	sourceUnitID string
	// verifiers DD
	verifiers [2]byte
}

// Segmento is an enumeration representing different segments of the judiciary.
type Segmento int

const (
	JusticaFederal          Segmento = 1
	JusticaEstadual         Segmento = 2
	JusticaTrabalho         Segmento = 3
	JusticaEleitoral        Segmento = 4
	JusticaMilitarEstadual  Segmento = 5
	JusticaMilitarFederal   Segmento = 6
	JusticaEstadualJuizados Segmento = 7
	JusticaFederalJuizados  Segmento = 8
)

// NewCNJ creates and returns a new CNJ instance from the given number string.
func NewCNJ(number string) CNJ {
	// The function parses the input number string and initializes the CNJ structure accordingly.
	// It splits the input into fragments and assigns them to the corresponding fields of the CNJ structure.

	var fragmentArray = [6]string{"0000000", "00", "0000", "0", "00", "0000"}
	var builder strings.Builder
	var fragmentIndex uint8
	for _, character := range number + "$" {
		if character < '0' || character > '9' {
			continue
		}
		builder.WriteRune(character)
		if builder.Len() == len(fragmentArray[fragmentIndex]) {
			fragmentArray[fragmentIndex] = builder.String()
			builder.Reset()
			fragmentIndex += 1
		}
	}

	segment, _ := strconv.Atoi(fragmentArray[3])
	processo := CNJ{
		number:            number,
		year:              fragmentArray[2],
		segment:           Segmento(segment),
		court:             fragmentArray[4],
		noticeOrderNumber: fragmentArray[0],
		sourceUnitID:      fragmentArray[5],
		verifiers:         [2]byte{fragmentArray[1][0], fragmentArray[1][1]},
	}
	return processo
}

// String returns the CNJ number as a string.
func (cnj CNJ) String() string {
	return cnj.number
}

// CourtNumber returns the court number from the CNJ.
func (cnj CNJ) CourtNumber() string {
	return cnj.court
}

// YearNumber returns the year from the CNJ.
func (cnj CNJ) YearNumber() string {
	return cnj.year
}
