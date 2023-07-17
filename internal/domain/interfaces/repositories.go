package interfaces

import "github.com/ViniciusDJM/jusbrasil-teste/internal/entities"

type ProcessRepository interface {
	FindFirstInstance() (result entities.JudicialProcess, err error)
	FindSecondInstance() (result entities.JudicialProcess, err error)
}
