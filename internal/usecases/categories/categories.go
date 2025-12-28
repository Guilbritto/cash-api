package categories

import (
	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/Guilbritto/cash-api/internal/models"
	"github.com/Guilbritto/cash-api/internal/repositories"
)

type CategoriesUseCase interface {
	Create(category dto.CreateCategoryRequest, userId string) (models.Category, error)
	GetAll(userId string) ([]models.Category, error)
}
type Service struct {
	CategoryRepository repositories.CategoryRepository
}
