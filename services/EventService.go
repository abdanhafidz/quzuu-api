package services

import (
	"errors"
	"reflect"
	"time"

	"github.com/quzuu-be/lib"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/repositories"
)

func CheckEventAssign(id_event int, id_account int) ServiceResult[models.EventAssign] {
	repositories.EventAssignRepository.Wrapper = models.EventAssign{
		IDAccount: uint(id_account),
		IDEvent:   uint(id_event),
	}
	var eventAssign = repositories.EventAssignRepository.Find()
	return ServiceResult[models.EventAssign]{
		Result: eventAssign.Result,
		Exception: lib.Exception{
			DataNotFound: eventAssign.NoRecord,
		},
		Error: eventAssign.RowsError,
	}
}

func EventRoleCheck(id_event int, id_account int) (res ServiceResult[bool]) {
	repositories.EventRepository.Wrapper.IDEvent = uint(id_event)
	eventDetail := repositories.EventRepository.Find()
	eventAssign := CheckEventAssign(id_account, id_event)
	res.Error = errors.Join(eventDetail.RowsError, eventAssign.Error)
	if eventDetail.Result.Public == "Y" && eventAssign.Exception.DataNotFound {
		res.Result = true
		res.Exception.UserNotRegisteredToEvent = true
	} else if eventDetail.Result.Public == "N" && eventAssign.Exception.DataNotFound {
		res.Result = false
		res.Exception.Unauthorized = true
	} else if !eventAssign.Exception.DataNotFound && !reflect.ValueOf(eventAssign.Result).IsNil() {
		res.Result = true
	}
	return res
}
func EventListService(req *models.EventListRequest, id_account int) ServiceResult[[]models.Events] {
	offset := req.PerPage * (req.PageNumber - 1)
	limit := req.PerPage
	filter := "%" + req.Filter + "%"
	eventList := repositories.EventPaginateRepository.FindAllPaginate(offset, limit, filter, id_account)
	return ServiceResult[[]models.Events]{
		Result: eventList.Result,
		Exception: lib.Exception{
			DataNotFound: eventList.NoRecord,
		},
		Error: eventList.RowsError,
	}
}

func EventDetailService(id_event int, id_account int) (res ServiceResult[models.EventResponse]) {
	var DetailResponse models.EventResponse
	repositories.EventRepository.Wrapper.IDEvent = uint(id_event)
	eventDetail := repositories.EventRepository.Find()
	authorizeEvent := EventRoleCheck(id_event, id_account)
	// fmt.Println(id_account, id_event)
	if !authorizeEvent.Exception.Unauthorized {
		DetailResponse.Data = &eventDetail.Result
		if !authorizeEvent.Exception.UserNotRegisteredToEvent {
			DetailResponse.RegisterStatus = 1
		} else {
			DetailResponse.RegisterStatus = 0
		}
	} else {
		res.Result = DetailResponse
		res.Exception.Unauthorized = true
	}
	res.Error = errors.Join(eventDetail.RowsError, authorizeEvent.Error)
	return res
}
func GetEventStatus(id_event int, id_account int) (res ServiceResult[bool]) {
	eventData := EventDetailService(id_event, id_account)
	res.Error = eventData.Error
	if !eventData.Exception.Unauthorized {
		startTime := eventData.Result.Data.StartEvent
		endTime := eventData.Result.Data.EndEvent
		curTime := time.Now()
		if curTime.After(startTime) && curTime.Before(endTime) {
			res.Result = true
			res.Exception.EventOnGoing = true
		} else if curTime.Before(startTime) {
			res.Result = false
			res.Exception.EventNotStart = true
		} else if curTime.After(endTime) {
			res.Result = false
			res.Exception.EventTimeOut = true
		}
		return res
	} else {
		res.Exception.Unauthorized = true
		return res
	}

}
func AssignUserToEvent(id_event int, id_account int) (res ServiceResult[models.Events]) {
	repositories.EventRepository.Wrapper.IDEvent = uint(id_event)
	repositories.EventAssignRepository.Wrapper = models.EventAssign{
		IDEvent:   uint(id_event),
		IDAccount: uint(id_account),
	}
	eventAssigned := repositories.EventRepository.Create()
	return ServiceResult[models.Events]{
		Result: eventAssigned.Result,
		Error:  eventAssigned.RowsError,
	}
}
func EventRegisterService(id_event int, id_account int, eventCode string) (res ServiceResult[models.Events]) {
	authorizeEvent := EventRoleCheck(id_event, id_account)
	repositories.EventRepository.Wrapper.IDEvent = uint(id_event)
	eventData := repositories.EventRepository.Find()
	if authorizeEvent.Result {
		res = ServiceResult[models.Events]{
			Result: eventData.Result,
			Error:  errors.Join(authorizeEvent.Error, eventData.RowsError),
		}
	} else if authorizeEvent.Exception.UserNotRegisteredToEvent {
		res = AssignUserToEvent(id_event, id_account)
		res.Error = errors.Join(res.Error, authorizeEvent.Error, eventData.RowsError)
	} else if authorizeEvent.Exception.Unauthorized {
		repositories.EventRepository.Wrapper.SID = eventCode
		if eventData.Result.IDEvent == uint(id_event) {
			res = AssignUserToEvent(id_event, id_account)
			res.Error = errors.Join(res.Error, authorizeEvent.Error)
		} else {
			res.Exception.InvalidEventCode = true
			res.Error = authorizeEvent.Error
		}
	}
	return res
}
