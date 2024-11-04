package models

import "time"

type AuditLog struct {
    ID        uint      `gorm:"primaryKey"`
    Action    string    `gorm:"not null"`
    Entity    string    `gorm:"not null"`  // e.g., 'product', 'order', 'user'
    EntityID  uint      `gorm:"not null"`
    UserID    uint      `gorm:"index"`
    User      User      `gorm:"foreignKey:UserID"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}
