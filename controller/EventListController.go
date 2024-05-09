package controller

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/services"
)

func EventListController(c *gin.Context) {
	var req models.EventListRequest
	c.ShouldBind(&req)
	token := c.Request.Header["Auth-Bearer-Token"]
	var ID_User int
	var verify_status string
	var err_verif error
	if(token != nil){
	ID_User, verify_status, err_verif = middleware.VerifyToken(token[0])
	if verify_status == "invalid-token" || verify_status == "expired" {
		ID_User = 0
	}
	}else{
		ID_User, verify_status = 0,"no-token"
	}
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
			fmt.Println(verify_status)
		}
	} else {
		msg := "There is an internal server error while sending request!"
		middleware.SendJSON500(c, &msg)
	}
}
