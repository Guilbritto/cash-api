package categories

import (
	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/Guilbritto/cash-api/internal/mappers"
	"github.com/Guilbritto/cash-api/internal/models"
)

func (s *Service) Create(category dto.CreateCategoryRequest, userId string) (*dto.CategoryResponse, error) {
	newCategory, err := models.NewCategory(category.Name, userId)
	if err != nil {
		return &dto.CategoryResponse{}, err
	}

	createdCategory, err := s.CategoryRepository.Save(newCategory)
	if err != nil {
		return &dto.CategoryResponse{}, err
	}

	return mappers.CategoryToResponse(createdCategory), nil
}
