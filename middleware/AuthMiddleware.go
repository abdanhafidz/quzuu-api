// auth/auth.go

package middleware

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/quzuu-be/config"
	"github.com/quzuu-be/models"
	"golang.org/x/crypto/bcrypt"
)

// Define a secret key for signing the JWT token

// GenerateToken generates a JWT token for the given user
func GenerateToken(user *models.Account) (string, error) {
	salt := config.Salt
	var secretKey = []byte(salt)
	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.IDAccount
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyPassword verifies if the provided password matches the hashed password
func VerifyPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("invalid password")
	}
	return nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
