package model

import (
	"time"

	"gorm.io/gorm"
)

// Role represents the column of the roles table
type Role struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	RoleName    string `gorm:"column:role_name"`
	Description string `gorm:"column:description"`
	IsActive    bool   `gorm:"column:is_active"`
	User        []User `gorm:"foreignKey:role_id;references:id"` // one to many relationship
}

// Convert roles table ame to m_roles
func (r *Role) TableName() string {
	return "m_roles"
}
