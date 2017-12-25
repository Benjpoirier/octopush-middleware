package server

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/lzientek/octopush-middleware/controllers"
	"github.com/lzientek/octopush-middleware/db"
	"github.com/lzientek/octopush-middleware/models"
)

var authMiddleware = &jwt.GinJWTMiddleware{
	Realm:      "user zone",
	Key:        []byte("89ef52f2-f3db-490c-91d6-bfbcb914f3a2"),
	Timeout:    time.Hour,
	MaxRefresh: time.Hour,
	Authenticator: func(email string, password string, c *gin.Context) (string, bool) {
		err, user := new(models.User).Login(db.GetDB(), email, password)
		if err == nil {
			return user.ID, true
		}

		return "", false
	},

	TokenLookup: "header:Authorization",
}

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("api/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)

		users := v1.Group("/users")
		{
			userController := new(controllers.UserController)

			users.POST("/", userController.Create)
		}

		templates := v1.Group("/templates")
		{
			templates.Use(authMiddleware.MiddlewareFunc())

			smsTemplateController := new(controllers.SmsTemplateController)
			templates.GET("/", smsTemplateController.GetAll)
			templates.POST("/", smsTemplateController.Create)
			templates.PUT("/:id", smsTemplateController.Update)
			templates.GET("/:smsTemplateId/sent", smsTemplateController.Update)
		}
	}

	return router
}
