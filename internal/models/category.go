package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	Id           string        `db:"id" json:"id" validate:"required"`
	Name         string        `db:"name" json:"name" validate:"required"`
	UserId       string        `db:"user_id" json:"user_id" validate:"required"`
	Transactions []Transaction `db:"transactions" json:"transactions" validate:"required"`
	CreatedAt    time.Time     `db:"created_at" json:"created_at" validate:"required"`
	UpdatedAt    time.Time     `db:"updated_at" json:"updated_at" validate:"required"`
}

func NewCategory(name string, userId string) (*Category, error) {
	category := &Category{
		Id:        uuid.New().String(),
		Name:      name,
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return category, nil
}
