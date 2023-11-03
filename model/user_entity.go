package model

import (
	"time"

	"gorm.io/gorm"
)

// User represents the column of the users table
type User struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Username  string    `gorm:"column:username"`
	FullName  string    `gorm:"column:full_name"`
	Password  string    `gorm:"column:password"`
	RoleId    int64     `gorm:"column:role_id"`
	Balance   int64     `gorm:"column:balance"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone"`
	IsActive  bool      `gorm:"column:is_active"`
	Order     []Order   `gorm:"foreignKey:user_id;references:id"` // one to many relationship
	Product   []Product `gorm:"foreignKey:user_id;references:id"` // one to many relationship
	Role      *Role     `gorm:"foreignKey:role_id;references:id"` // many to one relationship (belongs to)
}

// Convert users table ame to m_users
func (u *User) TableName() string {
	return "m_users"
}
