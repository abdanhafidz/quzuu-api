package services

import (
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/repositories"
)

func EventListService(req *models.EventListRequest) (data interface{}, status string, err error) {
	offset := req.PerPage * (req.PageNumber - 1)
	limit := req.PerPage
	filter := "%" + req.Filter + "%"
	data, status, err = repositories.GetEventList(offset, limit, filter)
	return data, status, err
}
