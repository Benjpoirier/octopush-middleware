package server

import (
	"github.com/gin-gonic/gin"
	"github.com/lzientek/octopush-middleware/controllers"
	"github.com/lzientek/octopush-middleware/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("api/v1")
	{
		users := v1.Group("/users")
		{
			users.Use(middleware.AdminMiddleware())
			userController := new(controllers.UserController)
			users.GET("/", userController.GetAll)
			users.POST("/", userController.Create)
			users.PUT("/:id", userController.Update)
			users.GET("/:id", userController.Show)
		}

		templates := v1.Group("/templates")
		{
			templates.Use(middleware.ApiMiddleware())

			smsTemplateController := new(controllers.SmsTemplateController)
			templates.GET("/", smsTemplateController.GetAll)
			templates.POST("/", smsTemplateController.Create)
			templates.PUT("/:id", smsTemplateController.Update)
			templates.GET("/:id", smsTemplateController.Show)
		}

		sends := v1.Group("/send")
		{
			sends.Use(middleware.ApiMiddleware())
			sendTemplateController := new(controllers.SendTemplateController)

			sends.POST("/:smsTemplateId", sendTemplateController.Create, middleware.ApiMiddleware())
		}
	}

	return router
}
