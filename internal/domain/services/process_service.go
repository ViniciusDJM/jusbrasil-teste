package services

import (
	"errors"
	"sync"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/domain/interfaces"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
)

type (
	RepositoryFactoryCallback func(courtNumber string) interfaces.ProcessRepository
	ProcessService            struct {
		repoFactory RepositoryFactoryCallback
	}
)

func NewProcessService(factory RepositoryFactoryCallback) ProcessService {
	return ProcessService{factory}
}

func (serv ProcessService) selectProcessRepo(cnj entities.CNJ) interfaces.ProcessRepository {
	return serv.repoFactory(cnj.CourtNumber())
}
func (serv ProcessService) LoadProcess(cnj entities.CNJ) (result struct{ First, Second entities.JudicialProcess }, err error) {
	var (
		errorList [2]error
		wg        sync.WaitGroup
		repo      = serv.selectProcessRepo(cnj)
	)
	wg.Add(2)

	go func() {
		defer wg.Done()
		result.First, errorList[0] = repo.FindFirstInstance(cnj)
	}()

	go func() {
		defer wg.Done()
		result.Second, errorList[1] = repo.FindSecondInstance(cnj)
	}()

	wg.Wait()
	err = errors.Join(errorList[0], errorList[1])
	return
}
