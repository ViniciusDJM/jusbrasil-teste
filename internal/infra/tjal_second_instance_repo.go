package infra

import (
	"bytes"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/exp/slices"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
)

type TJALSecondRepository struct {
	datasource RequestDatasource
	commonRepository
}

func NewTJALSecondRepository(datasource RequestDatasource) (newRepo TJALSecondRepository) {
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

func (repo TJALSecondRepository) executeParseSecond(
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

func (repo TJALSecondRepository) parseModalRadio(
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

func (repo TJALSecondRepository) FindSecondInstance() (result entities.JudicialProcess, err error) {
	var body []byte
	if body, err = repo.datasource.SearchSecondInstance(datasources.SearchFilter{}); err != nil {
		return
	}

	htmlDocument, parseErr := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if parseErr != nil {
		err = parseErr
		return
	}

	modalRadioButtonSelector := htmlDocument.Find("#modalIncidentes")
	processCode := repo.parseModalRadio(modalRadioButtonSelector)

	filter := datasources.SearchFilter{ProcessCode: processCode}
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

func (repo TJALSecondRepository) parseProcessParts(
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
			case slices.Contains([]string{"apelante"}, label):
				if !slices.Contains(process.Appellant, toInsert) {
					process.Appellant = append(process.Appellant, toInsert)
				}
			case slices.Contains([]string{"apelado", "apelada"}, label):
				if !slices.Contains(process.Appellee, toInsert) {
					process.Appellee = append(process.Appellee, toInsert)
				}
			}
		}
	}
	return
}
