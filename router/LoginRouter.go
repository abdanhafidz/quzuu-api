package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/controller"
)

func LoginRoutes(route *gin.Engine) {
	homeRouter := route.Group("/api")
	{
		homeRouter.POST("/login", controller.LoginController)
	}
}
