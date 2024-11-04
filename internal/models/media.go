package models

import "time"

type Media struct {
	ID         uint   `gorm:"primaryKey"`
	URL        string `gorm:"not null"`
	FileType   string
	FileSize   int
	UploadedBy uint
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
