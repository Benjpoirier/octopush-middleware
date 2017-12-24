package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lzientek/octopush-middleware/db"
)

type UserController struct{}

func (u UserController) GetAll(c *gin.Context) {
	var users []User
	err := db.GetDb().Find(&users)

	if err != nil {
		c.JSON(500, gin.H{"message": "Error to retrieve user", "error": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"data": users})
	return
}

func (u UserController) Create(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err == nil {
		err := userDao.Create(&user)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{"data": user})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
