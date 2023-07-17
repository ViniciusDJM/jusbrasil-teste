package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/interfaces"
)

type HandlerManager struct {
	tjalRepo interfaces.TJALRepository
}

func NewController() HandlerManager {
	return HandlerManager{
		tjalRepo: infra.NewTJALFirstRepository(datasources.TJAlagoasDatasource{}),
	}
}

// FirstInstanceHandler godoc
//
//	@Summary		Search process number
//	@Description	Receive a list of params and then search for the process
//	@Tags			Judicial Process
//	@ID				first-instance
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	entities.JudicialProcess
//	@Router			/v1/first-instance [get]
func (ctrl HandlerManager) FirstInstanceHandler(ctx *fiber.Ctx) error {
	result, err := ctrl.tjalRepo.FindFirstInstance(entities.CNJ{})
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}
