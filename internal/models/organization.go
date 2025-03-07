package models

import "time"

type Organization struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	WebsiteUrl       string    `gorm:"size:255" json:"Website_url"`
	LogoUrl          string    `gorm:"size:255" json:"logo_url"`
	Name             string    `gorm:"not null;size:100" json:"title"`
	Description      string    `gorm:"not null" json:"description"`
	TelegramUsername string    `gorm:"not null;size:50" json:"telegram"`
	Email            string    `gorm:"not null;size:100;unique" json:"email"`
	PhoneNumber      string    `gorm:"not null;size:20" json:"phone_number"`
	Status           bool      `gorm:"default:false" json:"status"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
