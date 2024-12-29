package models

import "time"

type User struct {
	Id           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"default:'Editor'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}


