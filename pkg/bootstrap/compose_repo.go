package bootstrap

import (
	"github.com/ViniciusDJM/jusbrasil-teste/internal/infra"
)

type composedRepo struct {
	infra.TJALSecondRepository
	infra.TJALFirstRepository
}
