package repositories

import "github.com/Guilbritto/cash-api/internal/models"

type CategoryRepository interface {
	Save(category *models.Category) (models.Category, error)
	GetById(categoryId string) (*models.Category, error)
}
