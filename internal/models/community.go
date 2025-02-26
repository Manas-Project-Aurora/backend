package models

type Community struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `gorm:"unique;not null;size:100" json:"title"`
	Type  string `gorm:"unique;not null;size:100" json:"type"`
	URL   string `gorm:"unique;not null;size:255" json:"url"`
}
