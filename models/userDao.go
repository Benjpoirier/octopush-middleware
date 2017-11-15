package models

import (
	"github.com/lzientek/octopush-middleware/db"
	bcrypt "golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Email     string        `bson:"email" json:"email" binding:"required"`
	Password  string        `bson:"password" json:"password" binding:"required"`
	ApiSecret string        `bson:"api_secret" json:"api_secret"`
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
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err == nil {
		user.Password = string(hash[:])
		err = con.Insert(user)
		user.Password = ""
	}

	return err
}

func (h UserDao) Login(email string, password string) (error, User) {
	con := db.Init().C("users")
	defer db.CloseSession()

	var resUser User

	err := con.Find(bson.M{"email": email}).One(&resUser)

	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(resUser.Password), []byte(password))
	}

	return err, resUser
}
