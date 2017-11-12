package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lzientek/octopush-middleware/controllers"
	"github.com/lzientek/octopush-middleware/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("api/v1")
	{
		users := v1.Group("/users")
		{
			userController := new(controllers.UserController)

			users.GET("/", userController.GetAll)
			users.POST("/", userController.Create)
		}

		templates := v1.Group("/templates")
		{
			smsTemplateController := new(controllers.SmsTemplateController)
			templates.GET("/", smsTemplateController.GetAll)
			templates.POST("/", smsTemplateController.Create)
		}
	}

	return router
}
