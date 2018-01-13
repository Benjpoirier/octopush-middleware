package models

import "time"

type SmsTemplate struct {
	ID        string     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID    string     `gorm:"type:uuid;" json:"user_id"`
	Title     string     `json:"title" binding:"required"`
	Content   string     `json:"content" gorm:"size:2000" binding:"required"`
	Language  string     `json:"language" gorm:"size:5"`
	SmsSender string     `json:"sms_sender" gorm:"size:11"`
	SmsType   string     `json:"sms_type" gorm:"size:3"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
