package router

import (
	"github.com/gin-gonic/gin"
)

func StartService() {
	router := gin.Default()
	HomeRoutes(router)
	LoginRoutes(router)
	RegisterRoutes(router)
	EventListRoutes(router)
	EventDetailRoutes(router)
	ProblemSetRoutes(router)
	EventRegisterRoutes(router)
	router.Run()
}
