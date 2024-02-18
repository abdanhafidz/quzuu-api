package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/controller"
)

func HomeRoutes(route *gin.Engine) {
	homeRouter := route.Group("/api")
	{
		homeRouter.GET("/", controller.HomeController)
	}
}
