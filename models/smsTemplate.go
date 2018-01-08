package models

import "time"

type SmsTemplate struct {
	ID        string     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID    string     `gorm:"type:uuid;" json:"userId"`
	Title     string     `json:"title" binding:"required"`
	Content   string     `json:"content" binding:"required"`
	Language  string     `json:"language"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:""`
}
