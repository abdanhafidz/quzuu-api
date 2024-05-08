package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/services"
)

func ProblemSetController(c *gin.Context) {
	var req models.EventDetailRequest
	c.ShouldBind(&req)
	data, status, err := services.EventDetailService(&req)
	if err != nil && status != "no-record" {
		panic(err)
	}
	if status == "ok" {
		middleware.SendJSON200(c, &data)
	} else if status == "no-record" {
		msg := "There is no Event Data with that event ID"
		middleware.SendJSON500(c, &msg)
	} else if status == "unauthorized" {
		msg := "You aren't assigned to this event"
		middleware.SendJSON401(c, &msg)
	} else {
		msg := "There is an internal server error while sending request!"
		middleware.SendJSON500(c, &msg)
	}
}
