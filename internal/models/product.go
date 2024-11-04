package models

import "time"

type Product struct {
	ID           uint       `gorm:"primaryKey"`
	Name         string     `gorm:"not null"`                      // The name of the product.
	Slug         string     `gorm:"unique;not null"`               // A URL-friendly identifier.
	Description  string     `gorm:"type:text"`                     // A detailed description of the product.
	Price        float64    `gorm:"not null"`                      // The price of the product.
	Currency     string     `gorm:"not null;default:'INR"`         // The currency in which the price is specified.
	Stock        int        `gorm:"default:0"`                     // The number of items in stock.
	SKU          string     `gorm:"unique"`                        // Stock Keeping Unit, a unique identifier for the product.
	Brand        string     `gorm:"default:''"`                    // The brand of the product.
	Availability string     `gorm:"default:'OutOfStock'"`          // The availability status (e.g., InStock, OutOfStock).
	Image        string     `gorm:"type:text"`                     // URL of the product image.
	Weight       float64    `gorm:"default:0"`                     // The weight of the product, useful for shipping.
	Dimensions   string     `gorm:"type:text"`                     // The dimensions of the product (e.g., "L 20 x W 10 x H 5 cm").
	Categories   []Category `gorm:"many2many:product_categories;"` // Categories to which the product belongs.
	Tags         []Tag      `gorm:"many2many:product_tags;"`       // Tags associated with the product.
	Media        []Media    `gorm:"many2many:product_media;"`      // Additional media (videos, etc.) associated with the product.
	CreatedAt    time.Time  `gorm:"autoCreateTime"`                // Timestamp when the product was created.
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`                // Timestamp when the product was last updated.
}
