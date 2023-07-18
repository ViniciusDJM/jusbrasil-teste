package services

import (
	"errors"
	"sync"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/domain/interfaces"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
)

type ProcessService struct {
	repo interfaces.ProcessRepository
}

func NewProcessService(repo interfaces.ProcessRepository) ProcessService {
	return ProcessService{repo}
}

func (serv ProcessService) LoadProcess(cnj entities.CNJ) (result struct{ First, Second entities.JudicialProcess }, err error) {
	var (
		errorList [2]error
		wg        sync.WaitGroup
	)
	wg.Add(2)

	go func() {
		defer wg.Done()
		result.First, errorList[0] = serv.repo.FindFirstInstance(cnj)
	}()

	go func() {
		defer wg.Done()
		result.Second, errorList[1] = serv.repo.FindSecondInstance(cnj)
	}()

	wg.Wait()
	err = errors.Join(errorList[0], errorList[1])
	return
}
