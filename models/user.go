package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	ID           string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Email        string `gorm:"unique_index" json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	APISecret    string `json:"api_secret"`
	SmsTemplates []SmsTemplate
}

func (h User) Login(db *gorm.DB, email string, password string) (error, User) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error

	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	}

	return err, user
}
