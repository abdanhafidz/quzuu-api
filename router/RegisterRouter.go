package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/controller"
)

func RegisterRoutes(route *gin.Engine) {
	RegisterRouter := route.Group("/api")
	{
		RegisterRouter.POST("/register", controller.RegisterController)
	}
}
