package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID           string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Email        string `gorm:"unique_index" json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	APISecret    string `json:"api_secret"`
	SmsTemplates []SmsTemplate
}
