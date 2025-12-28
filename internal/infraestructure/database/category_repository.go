package database

import (
	"github.com/Guilbritto/cash-api/internal/models"
	gorm "gorm.io/gorm"
)

type CategoryRepository struct {
	Db *gorm.DB
}

func (c *CategoryRepository) Save(category *models.Category) (models.Category, error) {
	tx := c.Db.Create(category)

	return *category, tx.Error
}

func (c *CategoryRepository) GetById(categoryId string) (*models.Category, error) {
	var category models.Category

	result := c.Db.Where("id= ?", categoryId).First(&category)

	if result.Error != nil {
		return &models.Category{}, result.Error
	}

	return &category, nil
}

func (c *CategoryRepository) GetAll(userId string) (*[]models.Category, error) {
	var categories []models.Category

	result := c.Db.Find(&categories).Where("user_id=? OR user_id=null", userId)
	if result.Error != nil {
		return &[]models.Category{}, result.Error
	}

	return &categories, nil

}
