package services

import (
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/repositories"
)

func RegisterService(regData *models.RegisterRequest) (data interface{}, status string, err error) {
	data, status, err = repositories.CreateAccount(regData)
	return data, status, err
}
