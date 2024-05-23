package repositories

import (
	"github.com/quzuu-be/models"
	"gorm.io/gorm"
)

func GetEventList(offset int, limit int, filter string, id_account int) (data interface{}, eventList *gorm.DB) {
	var e []models.Events
	eventList = db.Raw("(SELECT events.id_event,events.title, events.start_event, events.end_event, events.s_id, events.public FROM events WHERE public = 'Y' ) UNION (SELECT events.id_event,events.title, events.start_event, events.end_event, events.s_id, events.public FROM events INNER JOIN event_assign ON events.id_event=event_assign.id_event WHERE event_assign.id_account = ? )", id_account).Limit(limit).Offset(offset).Where("title LIKE ? OR 1=1", filter).Find(&e)
	return &e, eventList
}

func GetEventAssign(id_account int, id_event int) (res interface{}, eventAssign *gorm.DB) {
	var e []models.EventAssign
	if id_event != 0 {
		eventAssign = db.Where("id_account = ? AND id_event = ?", id_account, id_event).Find(&e)
	} else {
		eventAssign = db.Where("id_account = ?", id_account).Find(&e)
	}
	return &e, eventAssign
}

func GetEventDetail(id_event int) (data *models.Events, eventDetail *gorm.DB) {
	var e models.Events
	eventDetail = db.Where("id_event = ?", id_event).Find(&e)
	return &e, eventDetail
}

// func AssignEvent(id_account int, id_event int) (res interface{}, status string, err error) {
// 	return 0, 0, error()
// }
