package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/controller"
)

func ExamRoutes(route *gin.Engine) {
	examRouter := route.Group("/api")
	{
		examRouter.POST("/exam", controller.ExamController)
	}
}
