package models

import (
	"time"

	"github.com/google/uuid"
)

type BudgetItem struct {
	Id          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Description string    `gorm:"not null" json:"description"`
	Amount      float64   `gorm:"not null" json:"amount"`
	DueDate     time.Time `gorm:"type:date;not null" json:"due_date"`
	CategoryId  uuid.UUID `gorm:"type:uuid;not null" json:"category_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
