package infra

import (
	_ "embed"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/mocks"
)

var (
	//go:embed test/fixtures/tjal_second_instance_list.html
	secondInstanceSearchBody []byte

	//go:embed test/fixtures/tjal_second_instance.html
	secondInstanceShowBody []byte
)
var secondInstanceExpected = entities.JudicialProcess{
	Class:            "Apelação Cível",
	Area:             "Cível",
	Subject:          "Obrigações",
	DistributionDate: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
	Judge:            "",
	ActionValue:      281178.42,
	ProcessParts: entities.ProcessParts{
		Appellee: []entities.ProcessPeople{
			{Name: "José Carlos Cerqueira Souza Filho", Kind: 0},
			{Name: "Vinicius Faria de Cerqueira", Kind: 1},
			{Name: "Livia Nascimento da Rocha", Kind: 0},
		},
		Appellant: []entities.ProcessPeople{
			{Name: "Cony Engenharia Ltda.", Kind: 0},
			{Name: "Carlos Henrique de Mendonça Brandão", Kind: 1},
			{Name: "Guilherme Freire Furtado", Kind: 1},
			{Name: "Maria Eugênia Barreiros de Mello", Kind: 1},
			{Name: "Vítor Reis de Araujo Carvalho", Kind: 1},
			{Name: "Banco do Brasil S A", Kind: 0},
			{Name: "Nelson Wilians Fratoni Rodrigues", Kind: 1},
		},
	},
	MovementsList: []entities.Movement{
		{
			Date:        time.Date(2023, time.April, 26, 0, 0, 0, 0, time.UTC),
			Description: "Certidão de Envio ao 1º Grau\nFaço remessa dos presentes autos à Origem.",
		},
		{
			Date:        time.Date(2023, time.April, 26, 0, 0, 0, 0, time.UTC),
			Description: "Baixa Definitiva",
		},
		{
			Date:        time.Date(2023, time.April, 26, 0, 0, 0, 0, time.UTC),
			Description: "Certidão Emitida\nTERMO DE BAIXA Faço baixar estes autos ao Exmo(a). Juiz(a) de Direito da 4ª Vara Cível da Capital, em cumprimento ao despacho de página 872. Maceió, 26 de abril de 2023. Eleonora Paes Cerqueira de França Diretora Adjunta Especial de Assuntos Judiciários Cícera Cristina Lima de Araújo Bandeira Analista Judiciário",
		},
		{
			Date:        time.Date(2023, time.April, 12, 0, 0, 0, 0, time.UTC),
			Description: "Publicado",
		},
		{
			Date:        time.Date(2023, time.April, 12, 0, 0, 0, 0, time.UTC),
			Description: "Certidão Emitida\nCertifico que foi disponibilizado(a) no Diário da Justiça Eletrônico do Tribunal de Justiça de Alagoas, nesta data, o(a) Despacho/Decisão retro, nos termos do art 4º, § 3º, da Lei nº 11.419/2006. Maceió, 12 de abril de 2023 Eleonora Paes Cerqueira de França Diretora Adjunta Especial de Assuntos Judiciários",
		},
		{
			Date:        time.Date(2023, time.April, 12, 0, 0, 0, 0, time.UTC),
			Description: "Publicado",
		},
		{
			Date:        time.Date(2023, time.March, 23, 0, 0, 0, 0, time.UTC),
			Description: "Certidão Emitida\nFaço estes autos conclusos ao Excelentíssimo Senhor Vice Presidente do Tribunal de Justiça de Alagoas. Maceió, 23 de março de 2023 Eleonora Paes Cerqueira de França Diretora Adjunta Especial de Assuntos Judiciários Andréia Maria Oliveira da Silva Analista Judiciário",
		},
		{
			Date:        time.Date(2023, time.March, 23, 0, 0, 0, 0, time.UTC),
			Description: "Decisão dos Tribunais Superiores\n...conheço do agravo para negar provimento ao recurso especial",
		},
	},
}

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
	result, err := repo.FindSecondInstance(entities.NewCNJ("0710802-55.2018.8.02.0001"))
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, secondInstanceExpected) {
		t.Errorf(
			"Parsed process is not the same as expected\nReceived: `%+v`\nExpected: `%+v`",
			result, secondInstanceExpected,
		)
	}
}
