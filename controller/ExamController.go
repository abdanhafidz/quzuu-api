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
	ID_User, _, err_verif := middleware.AuthUser(c)
	IDEvent := req.IDEvent
	IDProblemSet := req.IDProblemSet
	data, status, err := services.ExamService(IDEvent, ID_User, IDProblemSet)
	if err != nil && status != "no-record" && err_verif != nil {
		err = errors.Join(err, err_verif)
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
