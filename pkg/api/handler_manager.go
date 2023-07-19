package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/ViniciusDJM/jusbrasil-teste/internal/domain/interfaces"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/domain/services"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/entities"
	"github.com/ViniciusDJM/jusbrasil-teste/pkg/api/models"
)

type HandlerManager struct {
	serv services.ProcessService
}

func NewController(alRepo, ceRepo interfaces.ProcessRepository) HandlerManager {
	return HandlerManager{
		services.NewProcessService(alRepo, ceRepo),
	}
}

// SearchHandler godoc
//
//	@Summary		Search process number
//	@Description	Receive a list of params and then search for the process
//	@Tags			Judicial Process
//	@ID				search-process
//	@Accept			json
//	@Produce		json
//	@Param			processNumber	query		string	false	"processNumber that will be searched"
//	@Success		200				{object}	models.ProcessDataResponse
//	@Router			/v1/search [get]
func (ctrl HandlerManager) SearchHandler(ctx *fiber.Ctx) error {
	formInput := models.SearchProcessForm{Number: ctx.Query("processNumber")}
	_ = ctx.BodyParser(&formInput)
	if formInput.Number == "" {
		return ctx.Status(fiber.StatusPreconditionRequired).JSON(fiber.Map{"error": "you must pass processNumber query params"})
	}
	result, err := ctrl.serv.LoadProcess(entities.NewCNJ(formInput.Number))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	response := models.ProcessDataResponse{FirstInstance: convertEntityToResponse(result.First), SecondInstance: convertEntityToResponse(result.Second)}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// SearchBodyHandler godoc
//
//	@Summary		Search process number
//	@Description	Receive a list of params and then search for the process
//	@Tags			Judicial Process
//	@ID				search-process-post
//	@Accept			json
//	@Produce		json
//	@Param			data	body		models.SearchProcessForm	false	"data that will be searched"
//	@Success		200		{object}	models.ProcessDataResponse
//	@Router			/v1/search [post]
func (ctrl HandlerManager) SearchBodyHandler(ctx *fiber.Ctx) error {
	return ctrl.SearchHandler(ctx)
}
