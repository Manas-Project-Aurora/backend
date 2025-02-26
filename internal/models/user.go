package models

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"not null"`
	PasswordHash string `gorm:"not null"`
	Telegram     string `gorm:"not null"`
	IsAdmin      bool   `gorm:"default:false"`
}
