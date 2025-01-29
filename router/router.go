package router

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/config"
	"github.com/quzuu-be/controller"
	"github.com/quzuu-be/middleware"
)

func StartService() {
	router := gin.Default()
	routerGroup := router.Group("/api")
	{
		routerGroup.GET("/event-details", middleware.AuthUser, controller.EventDetailController)
		routerGroup.GET("/event-list", middleware.AuthUser, controller.EventListController)
		routerGroup.POST("/register-event", middleware.AuthUser, controller.EventRegisterController)
		routerGroup.POST("/exam", middleware.AuthUser, controller.ExamController)
		routerGroup.GET("/problemset-list", middleware.AuthUser, controller.ProblemSetController)
		routerGroup.POST("/login", controller.LoginController)
		routerGroup.POST("/register", controller.RegisterController)
		routerGroup.GET("/", controller.HomeController)
	}
	router.Run(config.TCP_ADDRESS)
}
