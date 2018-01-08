package models

import "time"

type SendTemplate struct {
	ID            string     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	SmsTemplateID string     `gorm:"type:uuid;"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:""`
}
