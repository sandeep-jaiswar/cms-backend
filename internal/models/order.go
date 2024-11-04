package models

import "time"

type Order struct {
	ID         uint        `gorm:"primaryKey"`
	CustomerID uint        `gorm:"index"`
	Customer   Customer    `gorm:"foreignKey:CustomerID"`
	Total      float64     `gorm:"not null"`
	Status     string      `gorm:"default:'pending'"`
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
	Discounts  []Discount  `gorm:"many2many:order_discounts;"`
	CreatedAt  time.Time   `gorm:"autoCreateTime"`
}
