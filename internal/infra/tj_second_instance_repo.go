package infra

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
)

// TJSecondRepository is a struct representing the repository for searching the second instance of a judicial process in the TJ system.
type TJSecondRepository struct {
	datasource RequestDatasource
	commonRepository
}

// NewTJSecondRepository creates a new instance of the TJSecondRepository.
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

// executeParseSecond is a helper function that executes the parsing of the second instance data from the provided HTML document.
func (repo TJSecondRepository) executeParseSecond(
	htmlDocument *goquery.Document,
) (result entities.JudicialProcess, err error) {
	// Extract main process data
	mainProcessDataSelector := htmlDocument.Find(".unj-entity-header__summary > div:nth-child(1)")
	if mainProcessDataSelector != nil {
		result.Class = repo.findProcessClass(mainProcessDataSelector)
		result.Subject = repo.findProcessSubject(mainProcessDataSelector)
		result.Judge = repo.findProcessJudge(mainProcessDataSelector)
		result.Area = repo.findProcessArea(mainProcessDataSelector)
	}

	// Extract additional details data
	moreDetailsDataSelector := htmlDocument.Find("#maisDetalhes")
	if moreDetailsDataSelector != nil {
		repo.extractMoreDetails(moreDetailsDataSelector, &result)
	}

	// Extract movements data
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

	// Extract process parts data
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

// parseModalRadio is a helper function that parses the selected code from a radio button in a modal dialog.
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

// FindSecondInstance searches for the second instance of a judicial process using the provided CNJ (National Council of Justice) number.
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

	// Parse the selected code from a radio button in a modal dialog
	modalRadioButtonSelector := htmlDocument.Find("#modalIncidentes")
	processCode := repo.parseModalRadio(modalRadioButtonSelector)

	// Create a new filter to search the second instance using the extracted process code
	filter = datasources.SearchFilter{ProcessCode: processCode}
	if body, err = repo.datasource.SearchSecondInstance(filter); err != nil {
		return
	}

	if htmlDocument, parseErr = goquery.NewDocumentFromReader(bytes.NewReader(body)); parseErr != nil {
		err = parseErr
		return
	}

	// Execute the parsing of the second instance data
	result, err = repo.executeParseSecond(htmlDocument)
	return
}
