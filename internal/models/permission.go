package models

type Permission struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"unique;not null"`
    Description string
    Roles       []Role `gorm:"many2many:role_permissions;"`
}
