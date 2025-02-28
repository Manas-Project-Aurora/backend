package models

import "time"

type Video struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CompanyID   uint      `gorm:"not null" json:"company_id"`
	Image       string    `gorm:"size:255" json:"image"`
	Date        time.Time `gorm:"not null" json:"date"`
	YoutubeLink string    `gorm:"not null;size:255" json:"youtube_link"`
	Description string    `gorm:"not null" json:"description"`
	UserID      uint      `gorm:"not null" json:"user_id"`
}
