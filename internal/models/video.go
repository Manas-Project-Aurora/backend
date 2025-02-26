package models

import "time"

type Video struct {
	ID          uint `gorm:"primaryKey"`
	CompanyID   uint
	Image       string
	Date        time.Time `gorm:"not null"`
	YouTubeLink string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	UserID      uint
}
