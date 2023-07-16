package infra

import (
	_ "embed"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/mocks"
)

//go:embed test/fixtures/tjal_first_instance.html
var firstInstanceBody []byte

var firstInstanceExpected = entities.JudicialProcess{
	Class:            "Procedimento Comum Cível",
	Area:             "Cível",
	Subject:          "Dano Material",
	DistributionDate: time.Date(2018, time.May, 2, 19, 1, 0, 0, time.UTC),
	Judge:            "José Cícero Alves da Silva",
	ActionValue:      281178.42,
	ProcessParts: entities.ProcessParts{
		Author: []entities.ProcessPeople{
			{Name: "José Carlos Cerqueira Souza Filho", Kind: 0},
			{Name: "Vinicius Faria de Cerqueira", Kind: 1},
		},
		Defendant: []entities.ProcessPeople{
			{Name: "Cony Engenharia Ltda.", Kind: 0},
			{Name: "Carlos Henrique de Mendonça Brandão", Kind: 1},
			{Name: "Guilherme Freire Furtado", Kind: 1},
			{Name: "Maria Eugênia Barreiros de Mello", Kind: 1},
			{Name: "Vítor Reis de Araujo Carvalho", Kind: 1},
		},
	},
	MovementsList: []entities.Movement{
		{
			Date:        time.Date(2023, time.May, 5, 0, 0, 0, 0, time.UTC),
			Description: "Execução de Sentença Iniciada\nSeq.: 01 - Cumprimento de sentença",
		},
		{
			Date:        time.Date(2023, time.May, 5, 0, 0, 0, 0, time.UTC),
			Description: "Ato Publicado\nRelação: 0282/2023 Data da Publicação: 08/05/2023 Número do Diário: 3296",
		},
		{
			Date: time.Date(2023, time.May, 4, 0, 0, 0, 0, time.UTC),
			Description: "Disponibilização no Diário da Justiça Eletrônico\n" +
				"Relação: 0282/2023 Teor do ato: Autos n°: 0710802-55.2018.8.02.0001 " +
				"Ação: Procedimento Comum Cível Autor: José Carlos Cerqueira Souza Filho e outro " +
				"Réu: Cony Engenharia Ltda. e outro ATO ORDINATÓRIO Em cumprimento ao Provimento nº 15/2019, " +
				"da Corregedoria-Geral da Justiça do Estado de Alagoas, em virtude do retorno dos autos da instância " +
				"superior, manifestem-se as partes, em 15 (quinze) dias, requerendo o que de direito. Maceió, " +
				"04 de maio de 2023 Marcelo Rodrigo Falcão Vieira Analista(escrivão substituto) Advogados(s): " +
				"Nelson Wilians Fratoni Rodrigues (OAB 9395A/AL), Carlos Henrique de Mendonça Brandão (OAB 6770/AL), " +
				"Vinicius Faria de Cerqueira (OAB 9008/AL), Maria Eugênia Barreiros de Mello (OAB 14717/AL), " +
				"Guilherme Freire Furtado (OAB 14781/AL), Vítor Reis de Araujo Carvalho (OAB 14928/AL)",
		},
		{
			Date:        time.Date(2023, time.May, 4, 0, 0, 0, 0, time.UTC),
			Description: "Recebido pela Contadoria UNIFICADA\n",
		},
		{
			Date:        time.Date(2023, time.May, 4, 0, 0, 0, 0, time.UTC),
			Description: "Ato Ordinatório - Artigo 162, §4º, CPC\nAto Ordinatório- Remessa à contadoria",
		},
	},
}

func TestTJALRepository_FirstInstance(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	datasourceMock := mocks.NewMockRequestDatasource(mockCtrl)
	datasourceMock.EXPECT().DoRequest().Return(firstInstanceBody, nil).Times(1)

	repo := TJALRepository{datasource: datasourceMock}
	result, err := repo.FindFirstInstance()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(result, firstInstanceExpected) {
		t.Errorf(
			"Parsed process is not the same as expected\nReceived: `%v`\nExpected: `%v`",
			result, firstInstanceExpected,
		)
	}
}
