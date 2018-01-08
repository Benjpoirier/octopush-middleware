package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lzientek/octopush-middleware/db"
	"github.com/lzientek/octopush-middleware/models"
)

type SmsTemplateController struct{}

func (u SmsTemplateController) GetAll(c *gin.Context) {
	var smsTemplates []models.SmsTemplate
	err := db.GetDB().Find(&smsTemplates, models.SmsTemplate{UserID: c.MustGet("userID").(string)})

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": smsTemplates})
}

func (u SmsTemplateController) Create(c *gin.Context) {
	var template models.SmsTemplate

	if err := c.ShouldBindJSON(&template); err == nil {
		template.UserID = c.MustGet("userID").(string)
		err := db.GetDB().Create(&template).Error

		if err == nil {
			c.JSON(http.StatusOK, gin.H{"data": template})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (u SmsTemplateController) Update(c *gin.Context) {
	var template models.SmsTemplate

	if err := c.ShouldBindJSON(&template); err == nil {
		err := db.GetDB().Model(&template).Where("id = ? AND user_id = ?", c.Param("id"), c.MustGet("userID").(string)).Update(template).Error
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"data": template})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
