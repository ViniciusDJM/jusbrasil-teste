package models

import "time"

type ProcessPeopleDTO struct {
	Name     string `json:"name"`
	IsLawyer bool   `json:"is_lawyer"`
}
type ProcessPartsDTO struct {
	Appellant []ProcessPeopleDTO `json:"appellant,omitempty"`
	Appellee  []ProcessPeopleDTO `json:"appellee,omitempty"`
	Author    []ProcessPeopleDTO `json:"author,omitempty"`
	Defendant []ProcessPeopleDTO `json:"defendant,omitempty"`
	Victim    []string           `json:"victim,omitempty"`
	Third     []string           `json:"third,omitempty"`
	Witness   []string           `json:"Witness,omitempty"`
}
type MovementDTO struct {
	Date        string `json:"date"`
	Description string `json:"description"`
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

type ProcessDataResponse struct {
	FirstInstance  *JudicialProcessDTO `json:"first_instance,omitempty"`
	SecondInstance *JudicialProcessDTO `json:"second_instance,omitempty"`
}
