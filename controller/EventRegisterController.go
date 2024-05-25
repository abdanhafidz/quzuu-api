package controller

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/services"
)

func EventRegisterController(c *gin.Context) {
	var Req models.RegisterEventRequest
	c.ShouldBind(&Req)
	ID_User, statusAuth, err_verif := middleware.AuthUser(c)
	// fmt.Println(statusAuth)
	if statusAuth == "invalid-token" || statusAuth == "no-token" {
		status := "Unauthorized"
		msg := "Make sure that you've been Logged in before / you're not Authorized to access this endpoint"
		middleware.SendJSON401(c, &status, &msg)
		return
	}
	id_event := Req.IDEvent
	event_code := Req.EventCode
	data, status, err := services.EventRegisterService(id_event, ID_User, event_code)
	err = errors.Join(err, err_verif)
	fmt.Println(status)
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
