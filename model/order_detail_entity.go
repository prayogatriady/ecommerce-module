package model

import (
	"time"

	"gorm.io/gorm"
)

// OrderDetail represents the column of the order_details table
type OrderDetail struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	OrderId   int64    `gorm:"column:order_id"`
	ProductId int64    `gorm:"column:product_id"`
	Price     int64    `gorm:"column:price"`
	Quantity  int64    `gorm:"column:quantity"`
	Order     *Order   `gorm:"foreignKey:order_id;references:id"`   // many to one relationship (belongs to)
	Product   *Product `gorm:"foreignKey:product_id;references:id"` // many to one relationship (belongs to)
}

// Convert order_details table ame to t_order_details
func (o *OrderDetail) TableName() string {
	return "t_order_details"
}
