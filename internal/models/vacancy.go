package models

type Vacancy struct {
	ID             uint   `gorm:"primaryKey"`
	CompanyID      uint   `gorm:"not null"`
	JobTitle       string `gorm:"not null"`
	JobDescription string `gorm:"not null"`
	Type           string
	Salary         int    `gorm:"not null"`
	SalaryType     string `gorm:"not null"`
	Address        string `gorm:"not null"`
	UserID         uint   `gorm:"not null"`
	IsPublished    bool   `gorm:"default:false"`
}
