package infra

import (
	_ "embed"
	"reflect"
	"testing"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/test/fixtures"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/mocks"
	"github.com/golang/mock/gomock"
)

var (
	//go:embed test/fixtures/tjal_second_instance_list.testhtml
	alSecondInstanceSearchBody []byte

	//go:embed test/fixtures/tjal_second_instance.testhtml
	alSecondInstanceShowBody []byte
)

var (
	//go:embed test/fixtures/tjce_second_instance_list.testhtml
	ceSecondInstanceSearchBody []byte

	//go:embed test/fixtures/tjce_second_instance.testhtml
	ceSecondInstanceShowBody []byte
)

type TestCaseSecondInstance struct {
	name             string                   // Test case name
	input            string                   // Input CNJ (National Council of Justice) number
	processCode      string                   // Process code to simulate a second instance
	searchMockedBody []byte                   // Mocked response body for the search request
	showMockedBody   []byte                   // Mocked response body for the show request
	expected         entities.JudicialProcess // Expected result of the test case
}

// Run is a function that runs a single second instance test case
func (tcase TestCaseSecondInstance) Run(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	datasourceMock := mocks.NewMockRequestDatasource(mockCtrl)
	{
		datasourceMock.EXPECT().
			SearchSecondInstance(gomock.Any()).
			DoAndReturn(func(filter datasources.SearchFilter) ([]byte, error) {
				if filter.ProcessCode == tcase.processCode {
					return tcase.showMockedBody, nil
				}
				return tcase.searchMockedBody, nil
			}).
			Times(2)
	}

	repo := NewTJSecondRepository(datasourceMock)
	result, err := repo.FindSecondInstance(entities.NewCNJ(tcase.input))
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

// TestTJRepository_SecondInstance runs the test cases for searching the second instance of a judicial process
func TestTJRepository_SecondInstance(t *testing.T) {
	testCaseSecondInstances := [...]TestCaseSecondInstance{
		{
			name:             "Alagoas",
			input:            "0710802-55.2018.8.02.0001",
			processCode:      "P00006BXP0000",
			searchMockedBody: alSecondInstanceSearchBody,
			showMockedBody:   alSecondInstanceShowBody,
			expected:         fixtures.AlagoasSecondInstance,
		},
		{
			name:             "Ceara",
			input:            "0070337-91.2008.8.06.0001",
			processCode:      "P000020AM0000",
			searchMockedBody: ceSecondInstanceSearchBody,
			showMockedBody:   ceSecondInstanceShowBody,
			expected:         fixtures.CearaSecondInstance,
		},
	}
	for _, tc := range testCaseSecondInstances {
		t.Run(tc.name, tc.Run)
	}
}
