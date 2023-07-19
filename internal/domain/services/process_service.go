package services

import (
	"errors"
	"sync"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/domain/interfaces"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
)

type ProcessService struct {
	alRepo, ceRepo interfaces.ProcessRepository
}

func NewProcessService(alRepo, ceRepo interfaces.ProcessRepository) ProcessService {
	return ProcessService{alRepo, ceRepo}
}

func (serv ProcessService) selectProcessRepo(cnj entities.CNJ) interfaces.ProcessRepository {
	if cnj.CourtNumber() == "02" {
		return serv.alRepo
	}
	return serv.ceRepo
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
