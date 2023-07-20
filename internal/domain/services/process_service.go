package services

import (
	"errors"
	"sync"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/domain/interfaces"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
)

// RepositoryFactoryCallback is a type of function that returns an instance of ProcessRepository.
type (
	RepositoryFactoryCallback func(courtNumber string) interfaces.ProcessRepository
	ProcessService            struct {
		repoFactory RepositoryFactoryCallback
	}
)

// NewProcessService creates and returns a new instance of ProcessService with the provided repository factory.
func NewProcessService(factory RepositoryFactoryCallback) ProcessService {
	return ProcessService{factory}
}

// selectProcessRepo selects a process repository based on the court number (courtNumber) contained in the given CNJ.
func (serv ProcessService) selectProcessRepo(cnj entities.CNJ) interfaces.ProcessRepository {
	return serv.repoFactory(cnj.CourtNumber())
}

// LoadProcess loads the data of the first and second judicial processes associated with the given CNJ.
// It uses goroutines and synchronization with WaitGroup to load the processes concurrently.
// Returns the resulting structure with the processes and any errors that occurred during loading.
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
