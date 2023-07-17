package services

import (
	"errors"
	"log"
	"sync"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/domain/interfaces"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
)

type ProcessService struct {
	repo interfaces.ProcessRepository
}

func NewProcessService() ProcessService {
	return ProcessService{}
}

func (serv ProcessService) LoadProcess() {
	var (
		instanceReturns [2]entities.JudicialProcess
		errorList       [2]error
		wg              sync.WaitGroup
	)
	wg.Add(2)

	go func() {
		defer wg.Done()
		instanceReturns[0], errorList[0] = serv.repo.FindFirstInstance()
	}()

	go func() {
		defer wg.Done()
		instanceReturns[1], errorList[1] = serv.repo.FindSecondInstance()
	}()

	wg.Wait()
	err := errors.Join(errorList[0], errorList[1])
	log.Println(instanceReturns, err)
}
