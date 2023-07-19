package bootstrap

import (
	"github.com/ViniciusDJM/jusbrasil-teste/internal/domain/interfaces"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
	"github.com/ViniciusDJM/jusbrasil-teste/pkg/api"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(router fiber.Router) {
	ctrl := api.NewController(injectRepo())
	router.Get("/alagoas", ctrl.AlagoasHandler)
	router.Post("/alagoas", ctrl.AlagoasBodyHandler)
}

func injectRepo() interfaces.ProcessRepository {
	return composedRepo{
		TJALFirstRepository:  infra.NewTJALFirstRepository(datasources.TJAlagoasDatasource{}),
		TJALSecondRepository: infra.NewTJALSecondRepository(datasources.TJAlagoasDatasource{}),
	}
}
