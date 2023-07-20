package models

import "time"

type ProcessPeopleDTO struct {
	Name     string `json:"name"`
	IsLawyer bool   `json:"is_lawyer"`
}

// ProcessPartsDTO represents different parties involved in a judicial process, such as appellants, appellees, authors, defendants, victims, third parties, and witnesses. Each category contains a list of ProcessPeopleDTO objects.
type ProcessPartsDTO struct {
	Appellant []ProcessPeopleDTO `json:"appellant,omitempty"`
	Appellee  []ProcessPeopleDTO `json:"appellee,omitempty"`
	Author    []ProcessPeopleDTO `json:"author,omitempty"`
	Defendant []ProcessPeopleDTO `json:"defendant,omitempty"`
	Victim    []string           `json:"victim,omitempty"`
	Third     []string           `json:"third,omitempty"`
	Witness   []string           `json:"Witness,omitempty"`
}

// MovementDTO represents a movement in the judicial process, containing a date and a description.
type MovementDTO struct {
	Date        string `json:"date"`
	Description string `json:"description"`
}

// JudicialProcessDTO represents the main data of a judicial process, including its class, area, subject, distribution date, judge, action value, ProcessPartsDTO, and a list of MovementDTO.
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

// ProcessDataResponse represents the response for a judicial process data, with optional fields for first and second instance JudicialProcessDTO. It's used to send the data for both first and second instances of a judicial process in JSON format.
type ProcessDataResponse struct {
	FirstInstance  *JudicialProcessDTO `json:"first_instance,omitempty"`
	SecondInstance *JudicialProcessDTO `json:"second_instance,omitempty"`
}
