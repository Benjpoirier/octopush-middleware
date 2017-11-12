package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lzientek/octopush-middleware/models"
	"gopkg.in/mgo.v2/bson"
)

type SmsTemplateController struct{}

var smsTemplateDao = new(models.SmsTemplateDao)

func (u SmsTemplateController) GetAll(c *gin.Context) {
	templates, err := smsTemplateDao.GetAllByUser("")

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": templates})
	return
}

func (u SmsTemplateController) Create(c *gin.Context) {
	var template models.SmsTemplate
	template.ID = bson.NewObjectId()

	if err := c.ShouldBindJSON(&template); err == nil {
		err := smsTemplateDao.Create(&template)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{"data": template})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
