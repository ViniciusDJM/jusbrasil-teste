package entities

import (
	"testing"
)

// TestNewCNJ is a unit test for the NewCNJ function in the entities package.
func TestNewCNJ(t *testing.T) {
	tests := []struct {
		input    string
		expected CNJ
	}{
		// Test cases definition with inputs and expected results.
		{"1234567-89.2021.8.01.0001", CNJ{number: "1234567-89.2021.8.01.0001", year: "2021", segment: JusticaFederalJuizados, court: "01", noticeOrderNumber: "1234567", sourceUnitID: "0001", verifiers: [2]byte{'8', '9'}}},
		{"9876543-21.2022.3.02.1234", CNJ{number: "9876543-21.2022.3.02.1234", year: "2022", segment: JusticaTrabalho, court: "02", noticeOrderNumber: "9876543", sourceUnitID: "1234", verifiers: [2]byte{'2', '1'}}},
		{"5555555-01.2023.7.04.5678", CNJ{number: "5555555-01.2023.7.04.5678", year: "2023", segment: JusticaEstadualJuizados, court: "04", noticeOrderNumber: "5555555", sourceUnitID: "5678", verifiers: [2]byte{'0', '1'}}},
	}

	for _, test := range tests {
		// Call the NewCNJ function with the specific test input.
		processo := NewCNJ(test.input)

		// Compare each field of the returned structure with the expected results.
		if processo.year != test.expected.year {
			t.Errorf("Incorrect year. Expected: %s, Got: %s", test.expected.year, processo.year)
		}

		if processo.segment != test.expected.segment {
			t.Errorf("Incorrect segment. Expected: %v, Got: %v", test.expected.segment, processo.segment)
		}

		if processo.court != test.expected.court {
			t.Errorf("Incorrect court. Expected: %s, Got: %s", test.expected.court, processo.court)
		}

		if processo.noticeOrderNumber != test.expected.noticeOrderNumber {
			t.Errorf("Incorrect notice order number. Expected: %s, Got: %s", test.expected.noticeOrderNumber, processo.noticeOrderNumber)
		}

		if processo.sourceUnitID != test.expected.sourceUnitID {
			t.Errorf("Incorrect source unit ID. Expected: %s, Got: %s", test.expected.sourceUnitID, processo.sourceUnitID)
		}

		if processo.verifiers != test.expected.verifiers {
			t.Errorf("Incorrect verifiers. Expected: %v, Got: %v", test.expected.verifiers, processo.verifiers)
		}
	}
}
