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

func (repo commonRepository) findProcessSubject(selector *goquery.Selection) string {
	if subjectSelector := selector.Find(repo.selectors.mainData.subject); subjectSelector != nil {
		return strings.TrimSpace(subjectSelector.Text())
	}
	return ""
}

func (repo commonRepository) findProcessJudge(selector *goquery.Selection) string {
	if judgeSelector := selector.Find(repo.selectors.mainData.judge); judgeSelector != nil {
		return strings.TrimSpace(judgeSelector.Text())
	}
	return ""
}

func (repo commonRepository) findProcessArea(selector *goquery.Selection) string {
	if areaSelector := selector.Find(repo.selectors.mainData.area); areaSelector != nil {
		return strings.TrimSpace(areaSelector.Text())
	}
	return ""
}

func (repo commonRepository) parseDistributionDate(element *goquery.Selection) (time.Time, error) {
	if element == nil {
		return time.Time{}, nil
	}

	return utils.ParseBRTDateTime(element.Text())
}

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
