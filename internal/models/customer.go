package models

import "time"

type Customer struct {
    ID          uint      `gorm:"primaryKey"`
    UserID      uint      `gorm:"index"`
    User        User      `gorm:"foreignKey:UserID"`
    FirstName   string
    LastName    string
    Phone       string
    Address     string
    City        string
    State       string
    PostalCode  string
    CreatedAt   time.Time `gorm:"autoCreateTime"`
}
