package interfaces

import "github.com/ViniciusDJM/jusbrasil-teste/internal/entities"

type ProcessRepository interface {
	FindFirstInstance(entities.CNJ) (result entities.JudicialProcess, err error)
	FindSecondInstance(entities.CNJ) (result entities.JudicialProcess, err error)
}
