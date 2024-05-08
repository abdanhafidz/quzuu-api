package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/controller"
)

func ProblemSetRoutes(route *gin.Engine) {
	eventDetailRouter := route.Group("/api")
	{
		eventDetailRouter.GET("/problemset-list", controller.ProblemSetController)
	}
}
