package handlers

import (
	"github.com/Guilbritto/cash-api/internal/usecases"
)

type Handlers struct {
	UseCases usecases.UseCases
}

func New(useCases *usecases.UseCases) *Handlers {
	return &Handlers{
		UseCases: *useCases,
	}
}
