package bootstrap

import (
	"github.com/ViniciusDJM/jusbrasil-teste/internal/domain/interfaces"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra"
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra/datasources"
	"github.com/ViniciusDJM/jusbrasil-teste/pkg/api"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(router fiber.Router) {
	ctrl := api.NewController(injectRepo)
	router.Get("/search", ctrl.SearchHandler)
	router.Post("/search", ctrl.SearchBodyHandler)
}

func injectRepo(courtNumber string) interfaces.ProcessRepository {
	var dSource infra.RequestDatasource = datasources.TJAlagoasDatasource{}
	if courtNumber == "06" {
		dSource = datasources.TJCearaDatasource{}
	}
	return composedRepo{
		TJFirstRepository:  infra.NewTJFirstRepository(dSource),
		TJSecondRepository: infra.NewTJSecondRepository(dSource),
	}
}
