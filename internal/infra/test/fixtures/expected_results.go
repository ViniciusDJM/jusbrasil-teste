package fixtures

import (
	"time"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
)

var AlagoasFirstInstance = entities.JudicialProcess{
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
			{Name: "Livia Nascimento da Rocha", Kind: 0},
		},
		Defendant: []entities.ProcessPeople{
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
			Description: "Recebido pela Contadoria UNIFICADA",
		},
		{
			Date:        time.Date(2023, time.May, 4, 0, 0, 0, 0, time.UTC),
			Description: "Ato Ordinatório - Artigo 162, §4º, CPC\nAto Ordinatório- Remessa à contadoria",
		},
		{
			Date:        time.Date(2023, time.May, 5, 0, 0, 0, 0, time.UTC),
			Description: "Execução de Sentença Iniciada\nSeq.: 01 - Cumprimento de sentença",
		},
		{
			Date:        time.Date(2023, time.May, 5, 0, 0, 0, 0, time.UTC),
			Description: "Ato Publicado\nRelação: 0282/2023 Data da Publicação: 08/05/2023 Número do Diário: 3296",
		},
		{
			Date:        time.Date(2023, time.May, 4, 0, 0, 0, 0, time.UTC),
			Description: "Disponibilização no Diário da Justiça Eletrônico\nRelação: 0282/2023 Teor do ato: Autos n°: 0710802-55.2018.8.02.0001 Ação: Procedimento Comum Cível Autor: José Carlos Cerqueira Souza Filho e outro Réu: Cony Engenharia Ltda. e outro ATO ORDINATÓRIO Em cumprimento ao Provimento nº 15/2019, da Corregedoria-Geral da Justiça do Estado de Alagoas, em virtude do retorno dos autos da instância superior, manifestem-se as partes, em 15 (quinze) dias, requerendo o que de direito. Maceió, 04 de maio de 2023 Marcelo Rodrigo Falcão Vieira Analista(escrivão substituto) Advogados(s): Nelson Wilians Fratoni Rodrigues (OAB 9395A/AL), Carlos Henrique de Mendonça Brandão (OAB 6770/AL), Vinicius Faria de Cerqueira (OAB 9008/AL), Maria Eugênia Barreiros de Mello (OAB 14717/AL), Guilherme Freire Furtado (OAB 14781/AL), Vítor Reis de Araujo Carvalho (OAB 14928/AL)",
		},
	},
}

var AlagoasSecondInstance = entities.JudicialProcess{
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

var CearaFirstInstance = entities.JudicialProcess{
	Class:            "Ação Penal - Procedimento Ordinário",
	Area:             "Criminal",
	Subject:          "Crimes de Trânsito",
	DistributionDate: time.Date(2018, time.May, 2, 9, 13, 0, 0, time.UTC),
	Judge:            "",
	ActionValue:      0,
	ProcessParts: entities.ProcessParts{
		Victim: []string{
			"G. de O. C.",
			"A. S. F.",
		},
		Author: []entities.ProcessPeople{
			{Name: "Ministério Público do Estado do Ceará"},
		},
		Third: []string{
			"Departamento de Tecnologia da Informação e Comunicação - DETIC (Polícia Civil)",
		},
		Witness: []string{
			"M. L. S. I.",
		},
	},
	MovementsList: []entities.Movement{
		{
			Date:        time.Date(2021, time.September, 16, 0, 0, 0, 0, time.UTC),
			Description: "Juntada de Aviso de Recebimento (AR)",
		},
		{
			Date:        time.Date(2021, time.September, 10, 0, 0, 0, 0, time.UTC),
			Description: "Arquivado Definitivamente",
		},
		{
			Date:        time.Date(2022, time.August, 16, 0, 0, 0, 0, time.UTC),
			Description: "Juntada de Ofício\nNº Protocolo: WEB1.22.02299977-0\nTipo da Petição: Ofício\nData: 16/08/2022 12:49",
		},
	},
}

var CearaSecondInstance = entities.JudicialProcess{
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
