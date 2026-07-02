package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email        string    `json:"email" gorm:"size:255;not null;uniqueIndex" validate:"required,email"`
	FullName     string    `json:"full_name" gorm:"size:60;not null" validate:"required"`
	PasswordHash string    `json:"-" gorm:"not null"`
	IsActive     bool      `gorm:"default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
