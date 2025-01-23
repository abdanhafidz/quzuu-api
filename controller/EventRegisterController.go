package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/services"
)

func EventRegisterController(c *gin.Context) {
	var req models.RegisterEventRequest
	c.ShouldBind(&req)

	var account models.AccountData
	cParam, _ := c.Get("accountData")
	account = cParam.(models.AccountData)

	data, status, err := services.EventRegisterService(req.IdEvent, account.IdUser, req.EventCode)
	err = errors.Join(err, account.ErrVerif)
	// fmt.Println(status)
	if err != nil && status != "invalid-event-code" {
		panic(err)
	}
	if status == "ok" {
		middleware.SendJSON200(c, &data)
	} else if status == "registered" {
		status := "UserRegistered"
		msg := "User Has been Registered"
		middleware.SendJSON400(c, &status, &msg)
	} else if status == "invalid-event-code" {
		status := "InvalidEventCode"
		msg := "Invalid Event SID/Code"
		middleware.SendJSON400(c, &status, &msg)
	} else {
		msg := "There is an internal server error while sending request!"
		middleware.SendJSON500(c, &status, &msg)
	}

}
