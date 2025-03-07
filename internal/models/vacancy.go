package models

import "time"

type VacncyStatus string

const (
	Pending  = "Pending"
	Active   = "Active"
	Archived = "Archived"
)

type Vacancy struct {
	ID             uint         `gorm:"primaryKey" json:"id"`
	OrganizationID uint         `gorm:"not null" json:"organization_id"`
	Org            Organization `gorm:"foreignKey:OrganizationID" json:"-"`
	Title          string       `gorm:"not null;size:100" json:"title"`
	Description    string       `gorm:"not null" json:"description"`
	Type           string       `gorm:"size:50" json:"type"`
	SalaryFrom     int          `gorm:"default:0" json:"salary_from"`
	SalaryTo       int          `gorm:"default:0" json:"salary_to"`
	SalaryType     string       `gorm:"size:50" json:"salary_type"`
	Currency       string       `gorm:"size:50" json:"currecny"`
	Address        string       `gorm:"not null;size:255" json:"address"`
	UserID         uint         `gorm:"not null" json:"user_id"`
	Status         VacncyStatus `gorm:"default:Pending" json:"status"`
	CreatedAt      time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}
