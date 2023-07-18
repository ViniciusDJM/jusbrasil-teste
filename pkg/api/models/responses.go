package models

import "time"

type ProcessPeopleDTO struct {
	Name     string `json:"name"`
	IsLawyer bool   `json:"is_lawyer"`
}
type ProcessPartsDTO struct {
	Appellant, Appellee, Author, Defendant []ProcessPeopleDTO
}
type MovementDTO struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

type JudicialProcessDTO struct {
	Class            string          `json:"class"`
	Area             string          `json:"area"`
	Subject          string          `json:"subject"`
	DistributionDate time.Time       `json:"distribution_date"`
	Judge            string          `json:"judge"`
	ActionValue      float64         `json:"action_value"`
	ProcessParts     ProcessPartsDTO `json:"process_parts"`
	MovementsList    []MovementDTO   `json:"movements_list"`
}
