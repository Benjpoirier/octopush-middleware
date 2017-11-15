package models

import (
	"errors"

	"github.com/lzientek/octopush-middleware/db"
	"gopkg.in/mgo.v2/bson"
)

type SmsTemplate struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	UserID  bson.ObjectId `bson:"user_id" json:"userId"`
	Title   string        `form:"title" json:"title" binding:"required"`
	Content string        `bson:"content" json:"content" binding:"required"`
}

type SmsTemplateDao struct{}

func (h SmsTemplateDao) GetAllByUser(id string) ([]SmsTemplate, error) {

	con := db.Init().C("smsTemplate")

	defer db.CloseSession()

	var result []SmsTemplate

	err := con.Find(bson.M{"user_id": bson.ObjectIdHex(id)}).All(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (h SmsTemplateDao) Create(smsTemplate *SmsTemplate) error {

	con := db.Init().C("smsTemplate")

	defer db.CloseSession()
	smsTemplate.ID = bson.NewObjectId()
	err := con.Insert(smsTemplate)

	return err
}

func (h SmsTemplateDao) Update(id string, userID string, smsTemplate *SmsTemplate) error {

	con := db.Init().C("smsTemplate")

	defer db.CloseSession()

	var template SmsTemplate
	err := con.FindId(bson.ObjectIdHex(id)).One(&template)
	if err != nil {
		return err
	} else if template.UserID.Hex() != userID {
		return errors.New("Invalid User")
	}

	smsTemplate.ID = bson.ObjectIdHex(id)
	smsTemplate.UserID = template.UserID
	err = con.UpdateId(smsTemplate.ID, &smsTemplate)

	return err
}
