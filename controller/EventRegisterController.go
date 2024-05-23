package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/services"
)

func EventRegisterController(c *gin.Context) {
	var Req models.RegisterEventRequest
	c.ShouldBind(&Req)
	id_event := Req.IDEvent
	id_account := Req.IDAccount
	event_code := Req.EventCode
	data, status, err := services.EventRegisterService(id_event, id_account, event_code)
	if err != nil && status != "duplicate" {
		panic(err)
	}
	if status == "ok" {
		middleware.SendJSON200(c, &data)
	} else if status == "duplicate" {
		msg := "Email / Username has been used"
		middleware.SendJSON401(c, &status, &msg)
	} else {
		msg := "There is an internal server error while sending request!"
		middleware.SendJSON500(c, &status, &msg)
	}
}
