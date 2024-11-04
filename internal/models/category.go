package models

type Category struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique;not null"`
	Slug        string `gorm:"unique;not null"`
	Description string
	Products    []Product `gorm:"many2many:product_categories;"`
}
