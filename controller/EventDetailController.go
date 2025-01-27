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
	c.ShouldBindJSON(&req)

	var account models.AccountData
	cParam, _ := c.Get("accountData")
	account = cParam.(models.AccountData)

	data, status, err := services.EventDetailService(req.IdEvent, account.IdUser)
	if err != nil && status != "no-record" && account.ErrVerif != nil {
		err = errors.Join(err, account.ErrVerif)
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
