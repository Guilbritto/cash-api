package category

import (
	"github.com/Guilbritto/cash-api/internal/domain/entities"
	"github.com/Guilbritto/cash-api/internal/domain/repository"
	"github.com/Guilbritto/cash-api/internal/dto"
)

type CategoryService struct {
	CategoryRepository repository.CategoryRepository
}

func (s *CategoryService) Create(category dto.CreateCategoryRequest, userId string) (entities.Category, error) {
	newCategory, err := entities.NewCategory(category.Name, userId)
	if err != nil {
		return entities.Category{}, err
	}

	createdCategory, err := s.CategoryRepository.Save(newCategory)
	if err != nil {
		return entities.Category{}, err
	}

	return createdCategory, nil
}
