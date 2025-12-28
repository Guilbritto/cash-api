package dto

import (
	"github.com/google/uuid"
)

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryResponse struct {
	Id   uuid.UUID `json:"id" validate:"required"`
	Name string    `json:"name" validate:"required"`
}
