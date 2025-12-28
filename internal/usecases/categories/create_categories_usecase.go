package categories

import (
	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/Guilbritto/cash-api/internal/models"
)

func (s *Service) Create(category dto.CreateCategoryRequest, userId string) (models.Category, error) {
	newCategory, err := models.NewCategory(category.Name, userId)
	if err != nil {
		return models.Category{}, err
	}

	createdCategory, err := s.CategoryRepository.Save(newCategory)
	if err != nil {
		return models.Category{}, err
	}

	return createdCategory, nil
}
