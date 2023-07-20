package api

import (
	"time"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/pkg/api/models"
)

// convertEntityToResponse converts the entities.JudicialProcess to models.JudicialProcessDTO.
// It maps the fields from the entity to the response DTO.
func convertEntityToResponse(entity entities.JudicialProcess) (response *models.JudicialProcessDTO) {
	response = new(models.JudicialProcessDTO)
	response.ActionValue = entity.ActionValue
	response.Area = entity.Area
	response.Class = entity.Class
	response.DistributionDate = entity.DistributionDate
	response.Judge = entity.Judge
	response.MovementsList = parseMovimentList(entity.MovementsList)
	response.ProcessParts.Author = parseProcessParts(entity.ProcessParts.Author)
	response.ProcessParts.Appellant = parseProcessParts(entity.ProcessParts.Appellant)
	response.ProcessParts.Appellee = parseProcessParts(entity.ProcessParts.Appellee)
	response.ProcessParts.Defendant = parseProcessParts(entity.ProcessParts.Defendant)
	response.ProcessParts.Victim = entity.ProcessParts.Victim
	response.ProcessParts.Third = entity.ProcessParts.Third
	response.ProcessParts.Witness = entity.ProcessParts.Witness
	response.Subject = entity.Subject
	return
}

// parseProcessParts converts the list of entities.ProcessPeople to a list of models.ProcessPeopleDTO.
// It maps the fields from the entity to the response DTO, setting the "IsLawyer" field based on the "Kind" field.
func parseProcessParts(peopleList []entities.ProcessPeople) (response []models.ProcessPeopleDTO) {
	for _, person := range peopleList {
		response = append(response, models.ProcessPeopleDTO{
			Name:     person.Name,
			IsLawyer: person.Kind == entities.PersonLawyer,
		})
	}
	return
}

// parseMovimentList converts the list of entities.Movement to a list of models.MovementDTO.
// It maps the fields from the entity to the response DTO, formatting the date using the "time.DateOnly" format.
func parseMovimentList(movementList []entities.Movement) (response []models.MovementDTO) {
	for _, movement := range movementList {
		response = append(response, models.MovementDTO{
			Date:        movement.Date.Format(time.DateOnly),
			Description: movement.Description,
		})
	}
	return
}
