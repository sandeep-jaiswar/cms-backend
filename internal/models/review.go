package models

import "time"

type Review struct {
    ID         uint      `gorm:"primaryKey"`
    ProductID  uint      `gorm:"index"`
    Product    Product   `gorm:"foreignKey:ProductID"`
    CustomerID uint      `gorm:"index"`
    Customer   Customer  `gorm:"foreignKey:CustomerID"`
    Rating     int       `gorm:"not null;check:rating >= 1 AND rating <= 5"`
    ReviewText string
    CreatedAt  time.Time `gorm:"autoCreateTime"`
}
