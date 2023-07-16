package entities

import "time"

type PersonKind uint8

const (
	PersonNormal PersonKind = iota
	PersonLawyer
)

type (
	ProcessPeople struct {
		Name string
		Kind PersonKind
	}
	ProcessParts struct {
		Author    []ProcessPeople
		Defendant []ProcessPeople
	}
	Movement struct {
		Date        time.Time
		Description string
	}

	JudicialProcess struct {
		Class            string
		Area             string
		Subject          string
		DistributionDate time.Time
		Judge            string
		ActionValue      float64
		ProcessParts     ProcessParts
		MovementsList    []Movement
	}
)
