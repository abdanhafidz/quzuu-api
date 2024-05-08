package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/controller"
)

func EventDetailRoutes(route *gin.Engine) {
	eventDetailRouter := route.Group("/api")
	{
		eventDetailRouter.GET("/event-details", controller.EventDetailController)
	}
}
