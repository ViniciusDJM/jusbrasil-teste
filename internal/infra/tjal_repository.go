package infra

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/pkg/utils"
)

type TJALRepository struct {
	datasource RequestDatasource
	commonRepository
}

func NewTJALRepository(datasource RequestDatasource) TJALRepository {
	return TJALRepository{datasource: datasource}
}

func (repo TJALRepository) parseMoneyValue(element *goquery.Selection) float64 {
	if element == nil {
		return 0
	}

	return utils.ParseMonetaryValue(element.Text(), "R$")
}

func (repo TJALRepository) FindFirstInstance() (result entities.JudicialProcess, err error) {
	var body []byte
	if body, err = repo.datasource.SearchFirstInstance(); err != nil {
		return
	}

	htmlDocument, parseErr := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if parseErr != nil {
		err = parseErr
		return
	}

	mainProcessDataSelector := htmlDocument.Find("#containerDadosPrincipaisProcesso")
	if mainProcessDataSelector != nil {
		result.Class = repo.findProcessClass(mainProcessDataSelector)
		result.Subject = repo.findProcessSubject(mainProcessDataSelector)
		result.Judge = repo.findProcessJudge(mainProcessDataSelector)
	}

	moreDetailsDataSelector := htmlDocument.Find("#maisDetalhes")
	if moreDetailsDataSelector != nil {
		result.DistributionDate, _ = repo.parseDistributionDate(
			moreDetailsDataSelector.Find("#dataHoraDistribuicaoProcesso"),
		)
		result.Area = moreDetailsDataSelector.Find("#areaProcesso").Text()
		result.ActionValue = repo.parseMoneyValue(
			moreDetailsDataSelector.Find("#valorAcaoProcesso"),
		)
	}

	processPartsDataSelector := htmlDocument.Find("#tablePartesPrincipais")
	if processPartsDataSelector != nil {
		result.ProcessParts = repo.parseProcessParts(processPartsDataSelector)
	}

	movementsDataSelector := htmlDocument.Find("#tabelaUltimasMovimentacoes")
	if movementsDataSelector != nil {
		result.MovementsList = repo.parseMovementTable(movementsDataSelector)
	}

	return
}

func (repo TJALRepository) parseTableProcessParts(
	selector *goquery.Selection,
) (result map[string]string) {
	result = make(map[string]string, selector.Length())
	selector.Each(func(index int, selection *goquery.Selection) {
		var (
			label string
			name  string
		)
		labelSelector := selection.Find(".label")
		nameSelector := selection.Find(".nomeParteEAdvogado")
		if labelSelector != nil {
			label = strings.TrimSpace(labelSelector.Text())
		}
		if nameSelector != nil {
			name = NodeToStringSlice(nameSelector.Get(0))
		}
		result[label] = name
	})

	return
}

func (repo TJALRepository) parseProcessParts(
	selector *goquery.Selection,
) (process entities.ProcessParts) {
	tableBody := selector.ChildrenFiltered("tbody")
	if tableBody == nil {
		return
	}

	tableParts := tableBody.ChildrenFiltered("tr")
	partsMap := repo.parseTableProcessParts(tableParts)
	for label, content := range partsMap {
		nameList := strings.Split(content, "\n")
		for _, name := range nameList {
			kind := entities.PersonNormal
			if _, cutName, isLawyer := strings.Cut(name, "Advogado:"); isLawyer {
				kind = entities.PersonLawyer
				name = cutName
			}
			if _, cutName, isLawyer := strings.Cut(name, "Advogada:"); isLawyer {
				kind = entities.PersonLawyer
				name = cutName
			}
			switch strings.ToLower(label) {
			case "autor":
				process.Author = append(process.Author, entities.ProcessPeople{
					Name: strings.TrimSpace(name),
					Kind: kind,
				})
			case "ré":
				process.Defendant = append(process.Defendant, entities.ProcessPeople{
					Name: strings.TrimSpace(name),
					Kind: kind,
				})
			}
		}
	}
	return
}

func (repo TJALRepository) parseMovementTable(
	selector *goquery.Selection,
) (result []entities.Movement) {
	var currentMovement entities.Movement

	tableParts := selector.ChildrenFiltered("tr")
	tableParts.Each(func(i int, selection *goquery.Selection) {
		dateSelector := selection.Find(".dataMovimentacao")
		descriptionSelector := selection.Find(".descricaoMovimentacao")

		if dateSelector != nil {
			currentMovement.Date, _ = utils.ParseBRTDate(dateSelector.Text())
		}

		if descriptionSelector != nil && descriptionSelector.Length() > 0 {
			currentMovement.Description = NodeToStringSlice(descriptionSelector.Get(0))
		}

		result = append(result, currentMovement)
	})

	return
}
