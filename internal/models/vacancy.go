package models

type Vacancy struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	CompanyID      uint   `gorm:"not null" json:"company_id"`
	JobTitle       string `gorm:"not null;size:100" json:"job_title"`
	JobDescription string `gorm:"not null" json:"job_description"`
	Type           string `gorm:"size:50" json:"type"`
	Salary         int    `gorm:"default:0" json:"salary"`
	SalaryType     string `gorm:"size:50" json:"salary_type"`
	Address        string `gorm:"not null;size:255" json:"address"`
	UserID         uint   `gorm:"not null" json:"user_id"`
	IsPublished    bool   `gorm:"default:false" json:"is_published"`
}
