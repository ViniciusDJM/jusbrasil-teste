package entities

import (
	"strconv"
	"unicode"
)

type CNJ struct {
	number     string
	year       string
	segment    Segmento
	court      string
	sequential string
}

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

func NewCNJ(number string) CNJ {

	dados := make([]string, 0)
	numeros := make([]rune, 0)

	for _, char := range number {
		if unicode.IsDigit(char) {
			numeros = append(numeros, char)
		} else if len(numeros) > 0 {
			dados = append(dados, string(numeros))
			numeros = numeros[:0]
		}
	}

	seg, _ := strconv.Atoi(dados[2])

	processo := CNJ{
		number:     number,
		year:       dados[1],
		segment:    Segmento(seg),
		court:      dados[3],
		sequential: dados[4],
	}
	return processo
}

func (cnj CNJ) String() string {
	return cnj.number
}

func (cnj CNJ) CourtNumber() string {
	return cnj.court
}

func (cnj CNJ) YearNumber() string {
	return cnj.year
}
