package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lzientek/octopush-middleware/db"
	"github.com/lzientek/octopush-middleware/models"
)

type SendTemplateController struct{}

func (u SendTemplateController) GetAll(c *gin.Context) {
	var sendedTemplates []models.SendTemplate
	err := db.GetDB().Find(&sendedTemplates, models.SendTemplate{UserID: c.MustGet("userID").(string)})

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": sendedTemplates})
}

func (u SendTemplateController) Create(c *gin.Context) {
	var template models.SendTemplate

	if err := c.ShouldBindJSON(&template); err == nil {
		err := db.GetDB().Create(&template)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{"data": template})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
