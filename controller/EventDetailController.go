package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/services"
)

func EventDetailController(c *gin.Context) {

	var req models.EventDetailRequest
	c.ShouldBind(&req)
	ID_User, _, err_verif := middleware.AuthUser(c)
	data, status, err := services.EventDetailService(req.IDEvent, ID_User)
	if err != nil && status != "no-record" && err_verif != nil {
		err = errors.Join(err, err_verif)
		panic(err)
	}
	if status == "ok" {
		middleware.SendJSON200(c, &data)
	} else if status == "no-record" {
		msg := "There is no Event Data with that event ID"
		middleware.SendJSON500(c, &status, &msg)
	} else if status == "unauthorized" {
		msg := "You aren't assigned to this event"
		middleware.SendJSON401(c, &status, &msg)
	} else {
		msg := "There is an internal server error while sending request!"
		status = "InternalError"
		middleware.SendJSON500(c, &status, &msg)
	}
}
