package infra

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/pkg/utils"
)

type (
	movementSelectorConfig struct {
		date        string
		description string
	}
	mainSelectorConfig struct {
		class, subject, judge, area, distributionDate, actionValue string
	}
	processPartsSelectorConfig struct {
		label string
		name  string
	}
	commonRepository struct {
		selectors struct {
			movement     movementSelectorConfig
			mainData     mainSelectorConfig
			processParts processPartsSelectorConfig
		}
	}
)

func (repo commonRepository) findProcessClass(selector *goquery.Selection) string {
	if classSelector := selector.Find(repo.selectors.mainData.class); classSelector != nil {
		return strings.TrimSpace(classSelector.Text())
	}
	return ""
}

// Find the process subject from the provided HTML selection.
func (repo commonRepository) findProcessSubject(selector *goquery.Selection) string {
	if subjectSelector := selector.Find(repo.selectors.mainData.subject); subjectSelector != nil {
		return strings.TrimSpace(subjectSelector.Text())
	}
	return ""
}

// Find the process judge from the provided HTML selection.
func (repo commonRepository) findProcessJudge(selector *goquery.Selection) string {
	if judgeSelector := selector.Find(repo.selectors.mainData.judge); judgeSelector != nil {
		return strings.TrimSpace(judgeSelector.Text())
	}
	return ""
}

// Find the process area from the provided HTML selection.
func (repo commonRepository) findProcessArea(selector *goquery.Selection) string {
	if areaSelector := selector.Find(repo.selectors.mainData.area); areaSelector != nil {
		return strings.TrimSpace(areaSelector.Text())
	}
	return ""
}

// Parse the distribution date from the provided HTML element.
func (repo commonRepository) parseDistributionDate(element *goquery.Selection) (time.Time, error) {
	if element == nil {
		return time.Time{}, nil
	}

	return utils.ParseBRTDateTime(element.Text())
}

// Extract more details from the main data selection and update the result JudicialProcess entity.
func (repo commonRepository) extractMoreDetails(
	selector *goquery.Selection,
	result *entities.JudicialProcess,
) {
	result.DistributionDate, _ = repo.parseDistributionDate(
		selector.Find(repo.selectors.mainData.distributionDate),
	)

	if moneySelector := selector.Find(repo.selectors.mainData.actionValue); moneySelector != nil {
		result.ActionValue = utils.ParseMonetaryValue(moneySelector.Text(), "R$")
	}
}

// Parse the movement table and return the extracted movements.
func (repo commonRepository) parseMovementTable(
	selector *goquery.Selection,
) (result []entities.Movement) {
	var currentMovement entities.Movement

	tableParts := selector.ChildrenFiltered("tr")
	tableParts.Each(func(i int, selection *goquery.Selection) {
		dateSelector := selection.Find(repo.selectors.movement.date)
		descriptionSelector := selection.Find(repo.selectors.movement.description)

		if dateSelector != nil {
			currentMovement.Date, _ = utils.ParseBRTDate(strings.TrimSpace(dateSelector.Text()))
		}

		if descriptionSelector != nil && descriptionSelector.Length() > 0 {
			currentMovement.Description = strings.TrimSpace(
				NodeToStringSlice(descriptionSelector.Get(0)),
			)
		}

		result = append(result, currentMovement)
	})

	return
}

// Parse the process parts table and return the extracted data as a map.
func (repo commonRepository) parseTableProcessParts(
	selector *goquery.Selection,
) (result map[string]string) {
	result = make(map[string]string, selector.Length())
	selector.Each(func(index int, selection *goquery.Selection) {
		var (
			label string
			name  string
		)
		labelSelector := selection.Find(repo.selectors.processParts.label)
		nameSelector := selection.Find(repo.selectors.processParts.name)
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

// Parse the process parts from the provided HTML selection and update the process entity.
func (repo commonRepository) parseProcessParts(
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
			insertProcessPart(label, toInsert, process)
		}
	}
}
