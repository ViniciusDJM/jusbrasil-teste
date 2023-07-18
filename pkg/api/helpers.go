package api

import (
	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/pkg/api/models"
)

func convertEntityToResponse(entity entities.JudicialProcess) (response models.JudicialProcessDTO) {
	response.ActionValue = entity.ActionValue
	response.Area = entity.Area
	response.Class = entity.Class
	response.DistributionDate = entity.DistributionDate
	response.Judge = entity.Judge
	// response.MovementsList = entity.MovementsList
	response.ProcessParts.Author = parseProcessParts(entity.ProcessParts.Author)
	response.ProcessParts.Appellant = parseProcessParts(entity.ProcessParts.Appellant)
	response.ProcessParts.Appellee = parseProcessParts(entity.ProcessParts.Appellee)
	response.ProcessParts.Defendant = parseProcessParts(entity.ProcessParts.Defendant)
	response.Subject = entity.Subject
	return
}

func parseProcessParts(peopleList []entities.ProcessPeople) (response []models.ProcessPeopleDTO) {
	for _, person := range peopleList {
		response = append(response, models.ProcessPeopleDTO{
			Name:     person.Name,
			IsLawyer: person.Kind == entities.PersonLawyer,
		})
	}
	return
}
