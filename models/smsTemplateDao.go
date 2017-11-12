package models

import (
	"github.com/lzientek/octopush-middleware/db"
	"gopkg.in/mgo.v2/bson"
)

type SmsTemplate struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Title   string        `form:"title" json:"title" binding:"required"`
	Content string        `form:"content" json:"content" binding:"required"`
}
type SmsTemplateDao struct{}

func (h SmsTemplateDao) GetAllByUser(id string) ([]SmsTemplate, error) {

	con := db.Init().C("smsTemplate")

	defer db.CloseSession()

	var result []SmsTemplate

	err := con.Find(nil).All(&result)

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
