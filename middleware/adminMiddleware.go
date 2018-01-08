package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lzientek/octopush-middleware/config"
)

func AdminMiddleware() gin.HandlerFunc {
	config := config.GetConfig()

	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if strings.HasPrefix(auth, "Admin") {
			if strings.Split(auth, " ")[1] == config.GetString("app.admin_password") {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(401, gin.H{
			"code":    401,
			"message": "Unauthorized",
		})
	}
}
