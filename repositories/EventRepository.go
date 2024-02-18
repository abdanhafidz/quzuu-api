package repositories

import (
	"github.com/quzuu-be/config"
	"github.com/quzuu-be/models"
)

func GetEventList(offset int, limit int, filter string) (res interface{}, status string, err error) {
	var e models.Events
	db := config.DB
	eventList := db.Limit(limit).Offset(offset).Where("title LIKE ? OR 1=1", filter).Find(&e)
	err = eventList.Error
	if eventList.RowsAffected == 0 {
		status = "no-record"
	} else if err == nil {
		status = "ok"
	}
	return eventList, status, err
}
