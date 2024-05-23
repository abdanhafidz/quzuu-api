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
	ID_User, _, err_verif := middleware.AuthUser(c)
	data, status, err := services.EventListService(&req, ID_User)
	if err != nil && status != "no-record" && err_verif != nil {
		err = errors.Join(err, err_verif)
		panic(err)
	}
	if status == "ok" {
		middleware.SendJSON200(c, &data)
	} else if status == "no-record" {
		req.PerPage = 20
		req.PageNumber = 1
		req.Filter = ""
		data, status, err = services.EventListService(&req, ID_User)
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
