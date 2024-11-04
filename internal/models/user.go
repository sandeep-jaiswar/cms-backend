package models

import "time"

type User struct {
    ID           uint      `gorm:"primaryKey"`
    Username     string    `gorm:"unique;not null"`
    Email        string    `gorm:"unique;not null"`
    PasswordHash string    `gorm:"not null"`
    RoleID       uint      `gorm:"index"`
    Role         Role      `gorm:"foreignKey:RoleID"`
    CreatedAt    time.Time `gorm:"autoCreateTime"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
