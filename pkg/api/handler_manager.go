package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/domain/interfaces"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/domain/services"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
)

type HandlerManager struct {
	serv services.ProcessService
}

func NewController(repo interfaces.ProcessRepository) HandlerManager {
	return HandlerManager{
		services.NewProcessService(repo),
	}
}

// AlagoasHandler godoc
//
//	@Summary		Search process number
//	@Description	Receive a list of params and then search for the process
//	@Tags			Judicial Process
//	@ID				first-instance
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.JudicialProcessDTO
//	@Router			/v1/alagoas/:processNumber [get]
func (ctrl HandlerManager) AlagoasHandler(ctx *fiber.Ctx) error {
	processNumber := ctx.Query("processNumber")
	if processNumber == "" {
		return ctx.Status(fiber.StatusPreconditionRequired).JSON(fiber.Map{"error": "you must pass processNumber query params"})
	}
	result, err := ctrl.serv.LoadProcess(entities.NewCNJ(processNumber))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(convertEntityToResponse(result.First))
}
