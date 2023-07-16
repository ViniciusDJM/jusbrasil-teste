package infra

import (
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/ViniciusDJM/jusbrasil-teste/pkg/utils"
)

type commonRepository struct {
}

func (repo commonRepository) findProcessClass(selector *goquery.Selection) string {
	classSelector := selector.Find("#classeProcesso")
	if classSelector == nil {
		return ""
	}
	return classSelector.Text()
}

func (repo commonRepository) findProcessSubject(selector *goquery.Selection) string {
	subjectSelector := selector.Find("#assuntoProcesso")
	if subjectSelector == nil {
		return ""
	}
	return subjectSelector.Text()
}

func (repo commonRepository) findProcessJudge(selector *goquery.Selection) string {
	judgeSelector := selector.Find("#juizProcesso")
	if judgeSelector == nil {
		return ""
	}
	return judgeSelector.Text()
}

func (repo commonRepository) parseDistributionDate(element *goquery.Selection) (time.Time, error) {
	if element == nil {
		return time.Time{}, nil
	}

	return utils.ParseBRTDateTime(element.Text())
}
