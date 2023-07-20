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
	name             string
	input            string
	processCode      string
	searchMockedBody []byte
	showMockedBody   []byte
	expected         entities.JudicialProcess
}

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
			expected:         fixtures.AlagoasSecondInstance,
		},
	}
	for _, tc := range testCaseSecondInstances {
		t.Run(tc.name, tc.Run)
	}
}
