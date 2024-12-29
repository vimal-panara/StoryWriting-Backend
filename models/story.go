package models

import "time"

type Story struct {
	Id        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Content   string `gorm:"type:text;not null"`
	CreatedBy uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
