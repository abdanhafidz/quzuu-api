package services

import (
	"errors"

	"github.com/quzuu-be/middleware"
	"github.com/quzuu-be/models"
	"github.com/quzuu-be/repositories"
)

// LoginHandler handles user login
func LoginService(loginReq *models.LoginRequest) (token string, authStatus string, err error) {
	var errtok error = nil
	username := loginReq.Username
	password := loginReq.Password
	data, count, err := repositories.GetAccount(&username)
	verifyPass := middleware.VerifyPassword(data.Password, password)
	if verifyPass != nil || count == 0 {
		authStatus = "Akun dengan kredensial (username / password) tersebut tidak ditemukan!"
	} else {
		authStatus = "ok"
		token, errtok = middleware.GenerateToken(data)
	}

	err = errors.Join(err, errtok)
	return token, authStatus, err
}
