package repositories

import (
	"errors"

	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"gorm.io/gorm"
)

func GetAccount(username *string) (res *models.Account, count int, err error) {
	var Account models.Account
	find := db.Find(&Account, "username = ?", username)
	res = &Account
	count = int(find.RowsAffected)
	err = find.Error
	return res, count, err
}
func CreateAccount(data *models.RegisterRequest) (res *models.Account, status string, err error) {
	hash_password, errHash := middleware.HashPassword(data.Password)
	var Account = models.Account{
		Name:        data.Name,
		Username:    data.Username,
		Email:       data.Email,
		Password:    hash_password,
		PhoneNumber: data.Phone,
	}
	createUser := db.Create(&Account)
	Account.Password = "SECRET"
	res = &Account
	errDB := createUser.Error
	if errors.Is(errDB, gorm.ErrDuplicatedKey) {
		status = "duplicate"
	}
	err = errors.Join(errHash, errDB)
	if err == nil {
		status = "ok"
	}
	return res, status, err
}
