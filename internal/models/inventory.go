package models

import "time"

type Inventory struct {
	ID             uint    `gorm:"primaryKey"`
	ProductID      uint    `gorm:"index"`
	Product        Product `gorm:"foreignKey:ProductID"`
	QuantityChange int     `gorm:"not null"`
	Reason         string
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}
