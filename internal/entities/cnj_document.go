package entities

type CNJ struct {
	number string
}

func NewCNJ(number string) CNJ {
	return CNJ{number: number}
}

func (cnj CNJ) String() string {
	return cnj.number
}

func (cnj CNJ) CourtNumber() string {
	return cnj.number[5:]
}

func (cnj CNJ) YearNumber() string {
	return cnj.number[6:]
}
