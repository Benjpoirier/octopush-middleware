package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lzientek/octopush-middleware/db"
	"github.com/lzientek/octopush-middleware/lib/rand"
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
	var user models.CreateUser
	if err := c.ShouldBindJSON(&user); err == nil {
		var dbUser = models.User{
			Email:           user.Email,
			ThirdPartAPIKey: user.ThirdPartAPIKey,
			APIKey:          rand.String(32),
			APISecret:       rand.String(64),
		}

		err := db.GetDB().Create(&dbUser).Error

		if err == nil {
			c.JSON(http.StatusOK, gin.H{"data": dbUser})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (u UserController) Show(c *gin.Context) {
	var user models.User
	err := db.GetDB().First(&user, models.User{ID: c.Param("id")}).Error

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func (u UserController) Update(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err == nil {
		err := db.GetDB().Model(&user).Where("id = ?", c.Param("id")).Update(user).Error

		if err == nil {
			c.JSON(http.StatusOK, gin.H{"data": user})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
