package model

import (
	"time"

	"gorm.io/gorm"
)

// Product represents the column of the products table
type Product struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	ProductName string        `gorm:"column:product_name"`
	Description string        `gorm:"column:description"`
	Quantity    int64         `gorm:"column:quantity"`
	Price       int64         `gorm:"column:price"`
	UserId      int64         `gorm:"column:user_id"`
	IsActive    bool          `gorm:"column:is_active"`
	OrderDetail []OrderDetail `gorm:"foreignKey:product_id;references:id"` // one to many relationship
	User        *User         `gorm:"foreignKey:user_id;references:id"`    // many to one relationship (belongs to)
}

// Convert products table ame to m_products
func (p *Product) TableName() string {
	return "m_products"
}
