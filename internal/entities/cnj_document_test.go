package entities

import (
	"testing"
)

func TestNewCNJ(t *testing.T) {
	tests := []struct {
		input    string
		expected CNJ
	}{
		{"1234567-89.2021.8.01.0001", CNJ{year: "2021", segment: JusticaEstadualJuizados, court: "8", sequential: "0001"}},
		{"9876543-21.2022.3.02.1234", CNJ{year: "2022", segment: JusticaTrabalho, court: "3", sequential: "1234"}},
		{"5555555-01.2023.7.04.5678", CNJ{year: "2023", segment: JusticaEstadual, court: "7", sequential: "5678"}},
		// Adicione mais casos de teste conforme necessário
	}

	for _, test := range tests {
		processo := NewCNJ(test.input)

		if processo.year != test.expected.year {
			t.Errorf("Ano incorreto. Esperado: %s, Obtido: %s", test.expected.year, processo.year)
		}

		if processo.segment != test.expected.segment {
			t.Errorf("Segmento incorreto. Esperado: %v, Obtido: %v", test.expected.segment, processo.segment)
		}

		if processo.court != test.expected.court {
			t.Errorf("Tribunal incorreto. Esperado: %s, Obtido: %s", test.expected.court, processo.court)
		}

		if processo.sequential != test.expected.sequential {
			t.Errorf("Número Sequencial incorreto. Esperado: %s, Obtido: %s", test.expected.sequential, processo.sequential)
		}
	}
}
