package infra

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
)

// TJFirstRepository is a struct representing the repository for searching the first instance of a judicial process in the TJ system.
type TJFirstRepository struct {
	datasource RequestDatasource
	commonRepository
}

// NewTJFirstRepository creates a new instance of the TJFirstRepository.
func NewTJFirstRepository(datasource RequestDatasource) (newRepo TJFirstRepository) {
	newRepo.datasource = datasource
	newRepo.commonRepository.selectors.movement = movementSelectorConfig{
		date:        ".dataMovimentacao",
		description: ".descricaoMovimentacao",
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

// FindFirstInstance searches for the first instance of a judicial process using the provided CNJ (National Council of Justice) number.
func (repo TJFirstRepository) FindFirstInstance(cnj entities.CNJ) (result entities.JudicialProcess, err error) {
	var (
		body   []byte
		filter = datasources.SearchFilter{
			ProcessNumber:      cnj.String(),
			UnifiedYearNumber:  cnj.YearNumber(),
			UnifiedCourtNumber: cnj.CourtNumber(),
		}
	)
	if body, err = repo.datasource.SearchFirstInstance(filter); err != nil {
		return
	}

	htmlDocument, parseErr := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if parseErr != nil {
		err = parseErr
		return
	}

	// Extract main process data
	mainProcessDataSelector := htmlDocument.Find("#containerDadosPrincipaisProcesso")
	if mainProcessDataSelector != nil {
		result.Class = repo.findProcessClass(mainProcessDataSelector)
		result.Subject = repo.findProcessSubject(mainProcessDataSelector)
		result.Judge = repo.findProcessJudge(mainProcessDataSelector)
	}

	// Extract additional details data
	moreDetailsDataSelector := htmlDocument.Find("#maisDetalhes")
	if moreDetailsDataSelector != nil {
		repo.extractMoreDetails(moreDetailsDataSelector, &result)
		result.Area = repo.findProcessArea(moreDetailsDataSelector)
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

	// Extract movements data
	{
		movementsDataSelector := htmlDocument.Find("#tabelaUltimasMovimentacoes")
		if movementsDataSelector != nil {
			result.MovementsList = repo.parseMovementTable(movementsDataSelector)
		}

		movementsDataSelector = htmlDocument.Find("#tabelaTodasMovimentacoes")
		if movementsDataSelector != nil {
			result.MovementsList = append(result.MovementsList, repo.parseMovementTable(movementsDataSelector)...)
		}
	}

	return
}
