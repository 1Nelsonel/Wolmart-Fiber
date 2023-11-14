package models

import "time"

type FAQ struct {
	ID        uint `gorm:"primaryKey"`
	Question string `gorm:"not null"`
	Answer string `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}