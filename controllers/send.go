package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoisie/mustache"
	"github.com/lzientek/octopush-middleware/db"
	"github.com/lzientek/octopush-middleware/lib/octopush"
	"github.com/lzientek/octopush-middleware/models"
)

type SendTemplateController struct{}

func (u SendTemplateController) GetAll(c *gin.Context) {
	var sentTemplates []models.SendTemplate
	err := db.GetDB().Find(&sentTemplates, models.SendTemplate{SmsTemplateID: c.Param("smsTemplateId")}).Error

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve sendings", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": sentTemplates})
}

func (u SendTemplateController) Create(c *gin.Context) {
	var send models.ApiSendTemplate

	if err := c.ShouldBindJSON(&send); err == nil {
		var template models.SmsTemplate
		err := db.GetDB().Find(&template, c.Param("smsTemplateId")).Error
		if err != nil {
			user, _ := c.Get("user")
			Send(&template, send, user.(models.User))
		}

		if err == nil {
			c.JSON(http.StatusOK, gin.H{"data": template})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func Send(template *models.SmsTemplate, send models.ApiSendTemplate, user models.User) (octopush.OctopushResult, error) {
	sender := template.SmsSender
	if send.SmsSender != "" {
		sender = send.SmsSender
	}
	msg := mustache.Render(template.Content, send.Data)
	sms := octopush.OctopushSms{
		SmsRecipients: send.SmsRecipients,
		SmsSender:     sender,
		SmsText:       msg,
		Userlogin:     user.Email,
		APIKey:        user.ThirdPartAPIKey,
	}
	result, err := sms.Send()
	err = db.GetDB().Create(&send).Error
	return result, err
}
