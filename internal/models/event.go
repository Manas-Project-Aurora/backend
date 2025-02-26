package models

import "time"

type Event struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	CompanyID       uint      `gorm:"not null" json:"company_id"`
	Title           string    `gorm:"not null;size:100" json:"title"`
	Location        string    `gorm:"size:255" json:"location"`
	Date            time.Time `gorm:"not null" json:"date"`
	Description     string    `gorm:"not null" json:"description"`
	Website         string    `gorm:"size:255" json:"website"`
	RegistrationURL string    `gorm:"size:255" json:"registration_url"`
	UserID          uint      `gorm:"not null" json:"user_id"`
}
