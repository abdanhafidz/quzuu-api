package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/controller"
)

func AddEventRoutes(route *gin.Engine) {
	AddeventRouter := route.Group("/api/admin")
	{
		AddeventRouter.GET("/add-event", controller.AddEventController)
	}
}
