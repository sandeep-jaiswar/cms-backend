package models

import "time"

type Discount struct {
	ID             uint    `gorm:"primaryKey"`
	Code           string  `gorm:"unique;not null"`
	DiscountType   string  `gorm:"not null"` // e.g., 'percentage', 'flat'
	DiscountValue  float64 `gorm:"not null"`
	ExpirationDate time.Time
	Orders         []Order `gorm:"many2many:order_discounts;"`
}
