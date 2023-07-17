package infra

import (
	_ "embed"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/mocks"
)

var (
	//go:embed test/fixtures/tjal_second_instance_list.html
	secondInstanceSearchBody []byte

	//go:embed test/fixtures/tjal_second_instance.html
	secondInstanceShowBody []byte
)

func TestTJALRepository_SecondInstance(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	datasourceMock := mocks.NewMockRequestDatasource(mockCtrl)
	{
		datasourceMock.EXPECT().
			SearchSecondInstance(gomock.Any()).
			DoAndReturn(func(filter datasources.SearchFilter) ([]byte, error) {
				if filter.ProcessCode == "P00006BXP0000" {
					return secondInstanceShowBody, nil
				}
				return secondInstanceSearchBody, nil
			}).
			Times(2)
	}

	repo := NewTJALSecondRepository(datasourceMock)
	result, err := repo.FindSecondInstance()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, firstInstanceExpected) {
		t.Errorf(
			"Parsed process is not the same as expected\nReceived: `%+v`\nExpected: `%+v`",
			result, firstInstanceExpected,
		)
	}
}
