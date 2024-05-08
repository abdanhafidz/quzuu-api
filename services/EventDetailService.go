package services

import (
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/repositories"
)

func EventDetailService(req *models.EventDetailRequest, id_account int) (data interface{}, status string, err error) {
	id_event := req.IDEvent
	if id_account < 1 {
		id_account = 0
	}
	data, status, err = repositories.GetEventDetail(id_event, id_account)
	return data, status, err
}
