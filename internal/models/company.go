package models

type Company struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Website     string `gorm:"size:255" json:"website"`
	Logo        string `gorm:"size:255" json:"logo_url"`
	Title       string `gorm:"not null;size:100" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Telegram    string `gorm:"not null;size:50" json:"telegram"`
	Email       string `gorm:"not null;size:100;unique" json:"email"`
	PhoneNumber string `gorm:"not null;size:20" json:"phone_number"`
	Skype       string `gorm:"size:50" json:"skype"`
}
