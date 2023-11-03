package model

import (
	"time"

	"gorm.io/gorm"
)

// Order represents the column of the order table
type Order struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	OrderedAt   time.Time     `gorm:"column:ordered_at;autoCreateTime"`
	TotalPrice  int64         `gorm:"column:total_price"`
	UserId      int64         `gorm:"column:user_id"`
	VoucherId   int64         `gorm:"column:voucher_id"`
	OrderDetail []OrderDetail `gorm:"foreignKey:order_id;references:id"` // one to many relationship
	User        *User         `gorm:"foreignKey:user_id;references:id"`  // many to one relationship (belongs to)
}

// Convert orders table ame to t_orders
func (o *Order) TableName() string {
	return "t_orders"
}
