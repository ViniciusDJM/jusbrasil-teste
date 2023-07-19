package bootstrap

import (
	"github.com/ViniciusDJM/jusbrasil-teste/internal/domain/interfaces"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
	"github.com/ViniciusDJM/jusbrasil-teste/pkg/api"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(router fiber.Router) {
	ctrl := api.NewController(injectRepo(), injectRepo())
	router.Get("/search", ctrl.SearchHandler)
	router.Post("/search", ctrl.SearchBodyHandler)
}

func injectRepo() interfaces.ProcessRepository {
	return composedRepo{
		TJALFirstRepository:  infra.NewTJALFirstRepository(datasources.TJAlagoasDatasource{}),
		TJALSecondRepository: infra.NewTJALSecondRepository(datasources.TJAlagoasDatasource{}),
	}
}
