package models

import (
	"time"

	"github.com/google/uuid"
)

type BudgetPlanning struct {
	Id           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	BudgetId     uuid.UUID `gorm:"type:uuid;not null" json:"budget_id"`
	BudgetItemId uuid.UUID `gorm:"type:uuid;not null" json:"budget_item_id"`

	// Relacionamentos
	Budget     Budget     `gorm:"foreignKey:BudgetID" json:"budget,omitempty"`
	BudgetItem BudgetItem `gorm:"foreignKey:BudgetItemID" json:"budget_item,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}
