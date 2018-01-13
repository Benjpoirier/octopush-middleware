package models

import "time"

type SendTemplate struct {
	ID            string     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	SmsTemplateID string     `json:"sms_template_id" gorm:"type:uuid;"`
	SmsRecipients string     `json:"sms_recipients" binding:"required"`
	SmsSender     string     `json:"sms_sender" gorm:"size:11"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"-"`
}

type ApiSendTemplate struct {
	SendTemplate
	Data interface{} `json:"Data"`
}
