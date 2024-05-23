package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/controller"
)

func EventRegisterRoutes(route *gin.Engine) {
	registerEventRouter := route.Group("/api")
	{
		registerEventRouter.POST("/register-event", controller.EventRegisterController)
	}
}
