package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/services"
)

func EventListController(c *gin.Context) {
	var req models.EventListRequest
	c.ShouldBind(&req)

	var account models.AccountData
	cParam, _ := c.Get("accountData")
	account = cParam.(models.AccountData)

	data, status, err := services.EventListService(&req, account.IdUser)
	if err != nil && status != "no-record" && account.ErrVerif != nil {
		err = errors.Join(err, account.ErrVerif)
		panic(err)
	}
	if status == "ok" {
		middleware.SendJSON200(c, &data)
	} else if status == "no-record" {
		req.PerPage = 20
		req.PageNumber = 1
		req.Filter = ""
		data, status, err = services.EventListService(&req, account.IdUser)
		if status == "ok" && err == nil {
			middleware.SendJSON200(c, data)
		} else {
			msg := "There is an internal server error while sending request!"
			middleware.SendJSON500(c, &status, &msg)
		}
	} else {
		msg := "There is an internal server error while sending request!"
		status = "InternalError"
		middleware.SendJSON500(c, &status, &msg)
	}
}
