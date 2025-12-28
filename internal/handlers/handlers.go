package handlers

import "github.com/Guilbritto/cash-api/internal/useCases"

type Handlers struct {
	UseCases useCases.UseCases
}

func New(useCases *useCases.UseCases) *Handlers {
	return &Handlers{
		UseCases: *useCases,
	}
}
