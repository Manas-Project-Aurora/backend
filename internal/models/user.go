package models

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"unique;not null;size:50" json:"username"`
	PasswordHash string `gorm:"not null" json:"password_hash"`
	Telegram     string `gorm:"size:50" json:"telegram_username"`
	IsAdmin      bool   `gorm:"default:false" json:"is_admin"`
	IsActive     bool   `gorm:"default:false" json:"is_active"`
}
