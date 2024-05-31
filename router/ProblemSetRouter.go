package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/controller"
)

func ProblemSetRoutes(route *gin.Engine) {
	problemSetRouter := route.Group("/api")
	{
		problemSetRouter.GET("/problemset-list", controller.ProblemSetController)
		problemSetRouter.GET("/questions", controller.QuestionsController)
	}
}
