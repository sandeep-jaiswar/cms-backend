package models

import "time"

type Product struct {
    ID          uint      `gorm:"primaryKey"`
    Name        string    `gorm:"not null"`
    Slug        string    `gorm:"unique;not null"`
    Description string
    Price       float64   `gorm:"not null"`
    Stock       int       `gorm:"default:0"`
    SKU         string    `gorm:"unique"`
    Categories  []Category `gorm:"many2many:product_categories;"`
    Tags        []Tag      `gorm:"many2many:product_tags;"`
    Media       []Media   `gorm:"many2many:product_media;"`
    CreatedAt   time.Time `gorm:"autoCreateTime"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
