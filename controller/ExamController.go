package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/services"
)

func ExamController(c *gin.Context) {

	var req models.EventDetailRequest
	c.ShouldBind(&req)
	token := c.Request.Header["Auth-Bearer-Token"]
	var ID_User int
	var verify_status string
	var err_verif error
	if token != nil {
		ID_User, verify_status, err_verif = middleware.VerifyToken(token[0])
		if verify_status == "invalid-token" || verify_status == "expired" {
			ID_User = 0
		}
	} else {
		ID_User, verify_status = 0, "no-token"
	}
	data, status, err := services.ExamService(&req, ID_User)
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
		msg := "You aren't assigned to this problem set repository"
		middleware.SendJSON401(c, &status, &msg)
	} else {
		msg := "There is an internal server error while sending request!"
		status = "InternalError"
		middleware.SendJSON500(c, &status, &msg)
	}
}
