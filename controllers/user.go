package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lzientek/octopush-middleware/db"
	"github.com/lzientek/octopush-middleware/models"
)

type UserController struct{}

func (u UserController) GetAll(c *gin.Context) {
	var users []models.User
	err := db.GetDB().Find(&users).Error

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": users})
	return
}

func (u UserController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err == nil {
		err := db.GetDB().Create(&user).Error

		if err == nil {
			c.JSON(http.StatusOK, gin.H{"data": user})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
