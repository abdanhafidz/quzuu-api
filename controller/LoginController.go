package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/services"
)

func LoginController(c *gin.Context) {
	var loginReq models.LoginRequest
	c.ShouldBind(&loginReq)
	token, authStatus, err := services.LoginService(&loginReq)
	if err != nil {
		panic(err)
	}
	data := gin.H{
		"token": token,
	}
	if authStatus == "ok" {
		middleware.SendJSON200(c, &data)
	} else {
		status := "InvalidUser"
		middleware.SendJSON401(c, &status, &authStatus)
	}
}
