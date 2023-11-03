package model

import (
	"time"

	"gorm.io/gorm"
)

// Voucher represents the column of the vouchers table
type Voucher struct {
	ID           int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Description  string    `gorm:"column:description"`
	MinPrice     int64     `gorm:"column:min_price"`
	QuantityAll  int64     `gorm:"column:quantity_all"`
	QuantityUser int64     `gorm:"column:quantity_user"`
	ExpiredAt    time.Time `gorm:"column:expired_at"`
	IsActive     bool      `gorm:"column:is_active"`
	Order        []Order   `gorm:"foreignKey:voucher_id;references:id"` // one to many relationship
}

// Convert vouchers table ame to m_vouchers
func (v *Voucher) TableName() string {
	return "m_vouchers"
}
