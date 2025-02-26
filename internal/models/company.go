package models

type Company struct {
	ID          uint `gorm:"primaryKey"`
	Website     string
	Logo        string
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Telegram    string `gorm:"not null"`
	Email       string `gorm:"not null"`
	PhoneNumber string `gorm:"not null"`
	Skype       string
}
