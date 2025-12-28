package mappers

import (
	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/Guilbritto/cash-api/internal/models"
)

func CategoriesToResponse(categories []models.Category) *[]dto.CategoryResponse {
	response := make([]dto.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		response = append(response, dto.CategoryResponse{
			Id:   category.Id,
			Name: category.Name,
		})
	}
	return &response
}

func CategoryToResponse(category models.Category) *dto.CategoryResponse {
	response := dto.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
	return &response
}
