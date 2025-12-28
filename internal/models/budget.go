package models

import (
	"time"

	"github.com/google/uuid"
)

type Budget struct {
	Id        uuid.UUID `db:"id" gorm:"type:uuid;primaryKey" json:"id"`
	UserId    string    `db:"user_id" gorm:"column:user_id;not null" json:"user_id"`
	month     time.Time `db:"month" gorm:"type:date;not null" json:"month"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	BudggetMembers  []BudgetMember   `json:"budget_members" gorm:"foreignKey:BudgetmemberId;regference:Id"`
	BudgetPlannings []BudgetPlanning `gorm:"foreignKey:BudgetID" json:"plannings,omitempty"`
}
