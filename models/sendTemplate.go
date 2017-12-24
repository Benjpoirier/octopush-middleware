package models

import (
	"github.com/jinzhu/gorm"
)

type SendTemplate struct {
	gorm.Model
	ID            string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	SmsTemplateID string `gorm:"index"`
}
