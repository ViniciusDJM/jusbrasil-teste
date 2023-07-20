package infra

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
)

type TJSecondRepository struct {
	datasource RequestDatasource
	commonRepository
}

func NewTJSecondRepository(datasource RequestDatasource) (newRepo TJSecondRepository) {
	newRepo.datasource = datasource
	newRepo.commonRepository.selectors.movement = movementSelectorConfig{
		date:        ".dataMovimentacaoProcesso",
		description: ".descricaoMovimentacaoProcesso",
	}
	newRepo.commonRepository.selectors.mainData = mainSelectorConfig{
		class:            "#classeProcesso",
		subject:          "#assuntoProcesso",
		judge:            "#juizProcesso",
		area:             "#areaProcesso",
		distributionDate: "#dataHoraDistribuicaoProcesso",
		actionValue:      "#valorAcaoProcesso",
	}
	newRepo.commonRepository.selectors.processParts = processPartsSelectorConfig{
		label: ".label",
		name:  ".nomeParteEAdvogado",
	}

	return
}

func (repo TJSecondRepository) executeParseSecond(
	htmlDocument *goquery.Document,
) (result entities.JudicialProcess, err error) {
	mainProcessDataSelector := htmlDocument.Find(".unj-entity-header__summary > div:nth-child(1)")
	if mainProcessDataSelector != nil {
		result.Class = repo.findProcessClass(mainProcessDataSelector)
		result.Subject = repo.findProcessSubject(mainProcessDataSelector)
		result.Judge = repo.findProcessJudge(mainProcessDataSelector)
		result.Area = repo.findProcessArea(mainProcessDataSelector)
	}

	moreDetailsDataSelector := htmlDocument.Find("#maisDetalhes")
	if moreDetailsDataSelector != nil {
		repo.extractMoreDetails(moreDetailsDataSelector, &result)
	}

	{
		movementsDataSelector := htmlDocument.Find("#tabelaUltimasMovimentacoes")
		if movementsDataSelector != nil {
			result.MovementsList = repo.parseMovementTable(movementsDataSelector)
		}

		movementsDataSelector = htmlDocument.Find("#tabelaTodasMovimentacoes")
		if movementsDataSelector != nil {
			result.MovementsList = append(
				result.MovementsList,
				repo.parseMovementTable(movementsDataSelector)...,
			)
		}
	}

	{
		processPartsDataSelector := htmlDocument.Find("#tablePartesPrincipais")
		if processPartsDataSelector != nil {
			repo.parseProcessParts(processPartsDataSelector, &result.ProcessParts)
		}

		processPartsDataSelector = htmlDocument.Find("#tableTodasPartes")
		if processPartsDataSelector != nil {
			repo.parseProcessParts(processPartsDataSelector, &result.ProcessParts)
		}
	}

	return
}

func (repo TJSecondRepository) parseModalRadio(
	selector *goquery.Selection,
) (selectedCode string) {
	choiceDataSelector := selector.Find(".modal__process-choice")
	if choiceDataSelector == nil {
		return ""
	}

	inputSelector := choiceDataSelector.Find("#processoSelecionado")
	selectedCode, _ = inputSelector.Attr("value")
	return
}

func (repo TJSecondRepository) FindSecondInstance(
	cnj entities.CNJ,
) (result entities.JudicialProcess, err error) {
	var (
		body   []byte
		filter = datasources.SearchFilter{
			ProcessNumber:      cnj.String(),
			UnifiedYearNumber:  cnj.YearNumber(),
			UnifiedCourtNumber: cnj.CourtNumber(),
		}
	)
	if body, err = repo.datasource.SearchSecondInstance(filter); err != nil {
		return
	}

	htmlDocument, parseErr := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if parseErr != nil {
		err = parseErr
		return
	}

	modalRadioButtonSelector := htmlDocument.Find("#modalIncidentes")
	processCode := repo.parseModalRadio(modalRadioButtonSelector)

	filter = datasources.SearchFilter{ProcessCode: processCode}
	if body, err = repo.datasource.SearchSecondInstance(filter); err != nil {
		return
	}

	if htmlDocument, parseErr = goquery.NewDocumentFromReader(bytes.NewReader(body)); parseErr != nil {
		err = parseErr
		return
	}

	result, err = repo.executeParseSecond(htmlDocument)
	return
}
