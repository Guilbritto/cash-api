package category

import "github.com/Guilbritto/cash-api/internal/domain/entities"

type CategoryRepository interface {
	Save(category *entities.Category) (entities.Category, error)
}
