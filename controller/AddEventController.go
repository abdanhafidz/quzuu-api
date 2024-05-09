package controller

import (
	"github.com/gin-gonic/gin"
)

func AddEventController(c *gin.Context) {
	// var req models.EventListRequest
	// c.ShouldBind(&req)
	// data, status, err := services.EventListService(&req)
	// if err != nil && status != "no-record" {
	// 	panic(err)
	// }
	// if status == "ok" {
	// 	middleware.SendJSON200(c, &data)
	// } else if status == "no-record" {
	// 	req.PerPage = 20
	// 	req.PageNumber = 1
	// 	req.Filter = ""
	// 	data, status, err = services.EventListService(&req)
	// 	if status == "ok" && err == nil {
	// 		middleware.SendJSON200(c, &data)
	// 	} else {
	// 		panic(err)
	// 	}
	// } else {
	// 	msg := "There is an internal server error while sending request!"
	// 	middleware.SendJSON500(c, &msg)
	// }
}
