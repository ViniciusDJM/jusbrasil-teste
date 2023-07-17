package infra

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/exp/slices"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
)

type TJALFirstRepository struct {
	datasource RequestDatasource
	commonRepository
}

func NewTJALFirstRepository(datasource RequestDatasource) (newRepo TJALFirstRepository) {
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

func (repo TJALFirstRepository) FindFirstInstance() (result entities.JudicialProcess, err error) {
	var body []byte
	if body, err = repo.datasource.SearchFirstInstance(datasources.SearchFilter{}); err != nil {
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
		repo.extractMoreDetails(moreDetailsDataSelector, &result)
		result.Area = repo.findProcessArea(moreDetailsDataSelector)
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

	return
}

func (repo TJALFirstRepository) parseProcessParts(
	selector *goquery.Selection,
	process *entities.ProcessParts,
) {
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

			label = strings.Replace(strings.ToLower(label), ":", "", 1)
			toInsert := entities.ProcessPeople{Name: strings.TrimSpace(name), Kind: kind}
			switch {
			case slices.Contains([]string{"autor", "autora"}, label):
				if !slices.Contains(process.Author, toInsert) {
					process.Author = append(process.Author, toInsert)
				}
			case slices.Contains([]string{"ré", "réu"}, label):
				if !slices.Contains(process.Defendant, toInsert) {
					process.Defendant = append(process.Defendant, toInsert)
				}
			}
		}
	}
	return
}
