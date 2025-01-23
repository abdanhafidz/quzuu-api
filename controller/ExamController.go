package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/services"
)

func ExamController(c *gin.Context) {
	var req models.ExamRequest
	c.ShouldBind(&req)

	var account models.AccountData
	cParam, _ := c.Get("accountData")
	account = cParam.(models.AccountData)

	data, status, err := services.ExamService(req.IdEvent, account.IdUser, req.IdProblemSet)
	err = errors.Join(err, account.ErrVerif)
	if err != nil && status != "no-record" {
		panic(err)
	}
	if status == "ok" {
		middleware.SendJSON200(c, &data)
	} else if status == "no-record" {
		msg := "There is no Data with that Credential"
		middleware.SendJSON500(c, &status, &msg)
	} else if status == "unauthorized" {
		msg := "You aren't assigned to this problem set repository"
		middleware.SendJSON401(c, &status, &msg)
	} else if status == "unregistered" {
		msg := "You haven't registered to this event yet!"
		middleware.SendJSON401(c, &status, &msg)
	} else {
		msg := "There is an internal server error while sending request!"
		status = "InternalError"
		middleware.SendJSON500(c, &status, &msg)
	}
}
