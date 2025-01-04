package models

import "time"

type User struct {
	Id           uint   `gorm:"primaryKey" json:"-"`
	Username     string `gorm:"unique;not null" json:"username"`
	Email        string `gorm:"unique;not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"password"`
	Role         string `gorm:"default:'Editor'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
