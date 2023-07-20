package entities

import "time"

// PersonKind represents the type of person involved in a judicial process.
type PersonKind uint8

// Constants representing the types of people involved in a judicial process.
const (
	PersonNormal PersonKind = iota // Regular individual involved in the process.
	PersonLawyer                   // Lawyer involved in the process.
)

// ProcessPeople represents a person involved in a judicial process.
type ProcessPeople struct {
	Name string     // Name of the person.
	Kind PersonKind // Kind of person (regular individual or lawyer).
}

// ProcessParts represents the different roles (parts) people can have in a judicial process.
type ProcessParts struct {
	Appellant, Appellee, Author, Defendant []ProcessPeople // People involved as appellant, appellee, author, or defendant.
	Victim, Third, Witness                 []string        // Names of individuals involved as victims, third parties, or witnesses.
}

// Movement represents a movement (action or event) within the judicial process.
type Movement struct {
	Date        time.Time // Date when the movement occurred.
	Description string    // Description of the movement.
}

// JudicialProcess represents a complete judicial process with its details and movements.
type JudicialProcess struct {
	Class            string       // Classification of the process.
	Area             string       // Area to which the process belongs.
	Subject          string       // Subject of the process.
	DistributionDate time.Time    // Date when the process was distributed.
	Judge            string       // Name of the judge overseeing the process.
	ActionValue      float64      // Value associated with the action of the process.
	ProcessParts     ProcessParts // Roles and individuals involved in the process.
	MovementsList    []Movement   // List of movements/actions that occurred within the process.
}
