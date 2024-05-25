package services

import (
	"errors"
	"time"

	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/repositories"
	"gorm.io/gorm"
)

type Response struct {
	Data           *models.Events
	RegisterStatus int
}

func EventRoleCheck(id_event int, id_account int) (string, error) {
	eventDetail, rowsDetail := repositories.GetEventDetail(id_event)
	_, errDetail := middleware.RecordCheck(rowsDetail)
	dataAssign, statusAssign, errAssign := CheckEventAssign(id_account, id_event)
	err := errors.Join(errDetail, errAssign)
	// fmt.Println(eventDetail)
	// fmt.Println(statusAssign)
	if eventDetail.Public == "Y" {
		if statusAssign == "no-record" {
			return "unregistered", err
		} else {
			return "registered", err
		}
	} else {
		if statusAssign != "no-record" && dataAssign != nil {
			return "registered", err
		} else {
			return "unauthorized", err
		}
	}
}
func EventListService(req *models.EventListRequest, id_account int) (data interface{}, status string, err error) {
	offset := req.PerPage * (req.PageNumber - 1)
	limit := req.PerPage
	filter := "%" + req.Filter + "%"
	data, eventList := repositories.GetEventList(offset, limit, filter, id_account)
	status, err = middleware.RecordCheck(eventList)
	return data, status, err
}
func CheckEventAssign(id_account int, id_event int) (res interface{}, status string, err error) {
	data, eventAssign := repositories.GetEventAssign(id_account, id_event)
	status, err = middleware.RecordCheck(eventAssign)
	return data, status, err
}
func EventDetailService(id_event int, id_account int) (DetailResponse Response, status string, err error) {
	DataDetail, eventDetail := repositories.GetEventDetail(id_event)
	statusDetail, errDetail := middleware.RecordCheck(eventDetail)
	DetailResponse.Data = DataDetail
	statusAssign, errAssign := EventRoleCheck(id_event, id_account)
	err = errors.Join(errDetail, errAssign)
	// fmt.Println(id_account, id_event)
	if statusAssign != "unauthorized" {
		if statusAssign == "registered" {
			DetailResponse.RegisterStatus = 1
		} else {
			DetailResponse.RegisterStatus = 0
		}
		return DetailResponse, statusDetail, errAssign

	} else {
		return Response{
			Data:           &models.Events{},
			RegisterStatus: 0,
		}, statusAssign, err
	}

}
func GetEventStatus(id_event int, id_account int) (res string, status string, err error) {
	eventData, statusDetail, err := EventDetailService(id_event, id_account)
	if statusDetail == "ok" {
		startTime := eventData.Data.StartEvent
		endTime := eventData.Data.EndEvent
		curTime := time.Now()
		if curTime.After(startTime) && curTime.Before(endTime) {
			res = "OnGoing"
		} else if curTime.Before(startTime) {
			res = "NotStart"
		} else if curTime.After(endTime) {
			res = "Finish"
		}
		return res, statusDetail, err
	} else {
		return "error", statusDetail, err
	}

}

func EventRegisterService(id_event int, id_account int, eventCode string) (data interface{}, status string, err error) {
	var AssignUserToEvent *gorm.DB
	statusAssign, errAssign := EventRoleCheck(id_event, id_account)
	if statusAssign == "registered" {
		return data, statusAssign, errAssign
	} else if statusAssign == "unregistered" {
		// It means the event is public event
		data, AssignUserToEvent = repositories.CreateEventAssign(id_event, id_account)
		status, err = middleware.RecordCheck(AssignUserToEvent)
	} else if statusAssign == "unauthorized" {
		dataEvent, _ := repositories.GetEventDetailByCode(eventCode)
		// fmt.Println("data :", dataEvent)
		if dataEvent.IDEvent == uint(id_event) {
			data, AssignUserToEvent = repositories.CreateEventAssign(id_event, id_account)
			status, err = middleware.RecordCheck(AssignUserToEvent)
		} else {
			status = "invalid-event-code"
		}
	}
	return data, status, err
}
