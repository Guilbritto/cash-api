package database

import (
	"github.com/Guilbritto/cash-api/internal/domain/entities"
	gorm "gorm.io/gorm"
)

type CategoryRepository struct {
	Db *gorm.DB
}

func (c *CategoryRepository) Save(category *entities.Category) (entities.Category, error) {
	tx := c.Db.Create(category)

	return *category, tx.Error
}

func (c *CategoryRepository) GetById(categoryId string) (*entities.Category, error) {
	var category entities.Category

	result := c.Db.Where("id= ?", categoryId).First(&category)

	if result.Error != nil {
		return &entities.Category{}, result.Error
	}

	return &category, nil
}
