package interfaces

import "github.com/ViniciusDJM/jusbrasil-teste/internal/entities"

//go:generate mockgen -destination=../mocks/repositories_mock.go -package=mocks github.com/ViniciusDJM/jusbrasil-teste/internal/interfaces TJALRepository

type TJALRepository interface {
	FindFirstInstance(cnj entities.CNJ) (result entities.JudicialProcess, err error)
}
