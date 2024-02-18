package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/controller"
)

func EventListRoutes(route *gin.Engine) {
	eventListRouter := route.Group("/api")
	{
		eventListRouter.GET("/event-list", controller.EventListController)
	}
}
