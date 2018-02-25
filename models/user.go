package models

import (
	"time"
)

type CreateUser struct {
	Email           string `json:"email" binding:"required"`
	ThirdPartAPIKey string `json:"third_part_api_key"  binding:"required"`
}

type User struct {
	ID              string     `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Email           string     `gorm:"unique_index" json:"email"`
	APIKey          string     `json:"api_key" gorm:"unique_index"`
	APISecret       string     `json:"api_secret"  gorm:"index"`
	ThirdPartAPIKey string     `json:"third_part_api_key"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"-"`
}
