package models

import (
	"time"

	"github.com/google/uuid"
)

type BudgetMemberRole string

const (
	BudgetRoleOwner  BudgetMemberRole = "owner"
	BudgetRoleMember BudgetMemberRole = "member"
)

type BudgetMember struct {
	Id       uuid.UUID        `gorm:"type:uuid;primaryKey" json:"id"`
	BudgetId uuid.UUID        `gorm:"type:uuid;not null" json:"budget_id"`
	UserId   string           `gorm:"not null" json:"user_id"`
	Role     BudgetMemberRole `gorm:"type:varchar(20);not null" json:"role"`

	CreatedAt time.Time `json:"created_at"`
}
