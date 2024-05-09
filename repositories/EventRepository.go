package repositories

import (
	"github.com/quzuu-be/models"
)

func GetEventList(offset int, limit int, filter string, id_account int) (res interface{}, status string, err error) {
	var e []models.Events
	eventList := db.Rawr("(SELECT events.id_event,events.title, events.start_event, events.end_event, events.s_id, events.public FROM events WHERE public = 'Y' ) UNION (SELECT events.id_event,events.title, events.start_event, events.end_event, events.s_id, events.public FROM events INNER JOIN event_assign ON events.id_event=event_assign.id_event WHERE event_assign.id_account = ? )", id_account).Limit(limit).Offset(offset).Where("title LIKE ? OR 1=1", filter).Find(&e)
	err = eventList.Error
	if eventList.RowsAffected == 0 {
		status = "no-record"
	} else if err == nil {
		status = "ok"
	}
	return &e, status, err
}

func GetEventAssign(id_account int) (res interface{}, status string, err error) {
	var e []models.EventAssign
	eventAssign := db.Where("id_account = ?", id_account).Find(&e)
	err = eventAssign.Error
	if eventAssign.RowsAffected == 0 {
		status = "no-record"
	} else if err == nil {
		status = "ok"
	}
	return &e, status, err
}

func CheckEventAssign(id_account int, id_event int) (res interface{}, status string, err error) {
	var e models.EventAssign
	eventAssign := db.Where("id_account = ? AND id_event = ?", id_account, id_event).Find(&e)
	err = eventAssign.Error
	if eventAssign.RowsAffected == 0 {
		status = "no-record"
	} else if err == nil {
		status = "ok"
	}
	return &e, status, err
}
func GetEventDetail(id_event int, id_account int) (res interface{}, status string, err error) {
	var e models.Events
	eventDetail := db.Where("id_event = ?", id_event).Find(&e)
	err = eventDetail.Error
	if eventDetail.RowsAffected == 0 {
		status = "no-record"
	} else if err == nil {
		status = "ok"
	}
	if e.Public == "Y" {
		return &e, status, err
	} else {
		data, status, err := CheckEventAssign(id_account, id_event)
		if status != "no-record" && data != nil {
			return &e, status, err
		} else {
			return false, "unauthorized", err
		}
	}
}
