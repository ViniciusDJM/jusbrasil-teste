package infra

import (
	_ "embed"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/test/fixtures"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/mocks"
)

//go:embed test/fixtures/tjal_first_instance.html
var alFirstInstanceSearchBody []byte

//go:embed test/fixtures/tjce_first_instance.html
var ceFirstInstanceSearchBody []byte

type TestCase struct {
	name       string
	input      string
	mockedBody []byte
	expected   entities.JudicialProcess
}

func (tcase TestCase) Run(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	datasourceMock := mocks.NewMockRequestDatasource(mockCtrl)
	datasourceMock.EXPECT().
		SearchFirstInstance(gomock.Any()).
		Return(tcase.mockedBody, nil).
		Times(1)

	repo := NewTJFirstRepository(datasourceMock)
	result, err := repo.FindFirstInstance(entities.NewCNJ(tcase.input))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, tcase.expected) {
		t.Errorf(
			"Parsed process is not the same as expected\nReceived: `%+v`\nExpected: `%+v`",
			result, tcase.expected,
		)
	}
}

func TestTJALRepository_FirstInstance(t *testing.T) {
	testCases := [...]TestCase{
		{
			name:       "Alagoas",
			input:      "0710802-55.2018.8.02.0001",
			mockedBody: alFirstInstanceSearchBody,
			expected:   fixtures.AlagoasFirstInstance,
		},
		{
			name:       "Ceara",
			input:      "0070337-91.2008.8.06.0001",
			mockedBody: ceFirstInstanceSearchBody,
			expected:   fixtures.CearaFirstInstance,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, tc.Run)
	}
}
