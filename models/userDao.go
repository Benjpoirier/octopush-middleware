package models

import (
	"github.com/lzientek/octopush-middleware/db"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Email    string        `form:"email" json:"email" binding:"required"`
	Password string        `form:"password" json:"password" binding:"required"`
}
type UserDao struct{}

func (h UserDao) GetAll() ([]User, error) {
	con := db.Init().C("users")

	defer db.CloseSession()

	var result []User

	err := con.Find(nil).All(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (h UserDao) Create(user *User) error {
	con := db.Init().C("users")
	defer db.CloseSession()

	user.ID = bson.NewObjectId()

	err := con.Insert(user)

	return err
}
