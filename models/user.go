package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID        string     `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Email     string     `gorm:"unique_index" json:"email"`
	Password  string     `json:"-"`
	APISecret string     `json:"api_secret"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

func (h User) Login(db *gorm.DB, email string, password string) (error, User) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error

	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	}

	return err, user
}
func (u *User) BeforeUpdate() (err error) {
	if u.Password != "" {
		crypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 0)
		u.Password = string(crypted[:])
		return err
	}
	return
}

func (u *User) BeforeCreate() (err error) {
	if u.Password != "" {
		crypted, err := bcrypt.GenerateFromPassword([]byte(u.Password), 0)
		u.Password = string(crypted[:])
		return err
	}
	return
}
