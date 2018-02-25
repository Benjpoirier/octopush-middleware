package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lzientek/octopush-middleware/db"
	"github.com/lzientek/octopush-middleware/models"
)

func ApiMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("ApiKey")
		secret := c.GetHeader("ApiSecret")

		var user models.User
		err := db.GetDB().First(&user, models.User{APISecret: secret, APIKey: key}).Error

		if err == nil {
			c.Set("user", user)
			c.Next()
			return
		}
		c.AbortWithStatusJSON(401, gin.H{
			"code":    401,
			"message": "Unauthorized",
		})
	}
}
