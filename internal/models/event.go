package models

import "time"

type Event struct {
	ID              uint `gorm:"primaryKey"`
	CompanyID       uint
	Title           string `gorm:"not null"`
	Location        string
	Date            time.Time `gorm:"not null"`
	Description     string    `gorm:"not null"`
	Website         string
	RegistrationURL string
	UserID          uint
	IsPublished     bool `gorm:"default:false"`
}
