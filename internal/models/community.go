package models

type Community struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"unique;not null"`
	Type  string `gorm:"unique;not null"`
	URL   string `gorm:"unique;not null"`
}
