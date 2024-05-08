// auth/auth.go

package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/quzuu-be/config"
	"github.com/quzuu-be/models"
	"golang.org/x/crypto/bcrypt"
)

// Define a secret key for signing the JWT token
var salt = config.Salt
var secretKey = []byte(salt)

// GenerateToken generates a JWT token for the given user
func GenerateToken(user *models.Account) (string, error) {

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

type CustomClaims struct {
	IDUser int `json:"id"`
	jwt.StandardClaims
}

func VerifyToken() (int, string, error) {
	var r *http.Request
	bearer_token := r.Header.Get("Auth-Bearer-Token")
	token, err := jwt.Parse(bearer_token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	claims, ok := token.Claims.(*CustomClaims)

	if !ok || !token.Valid {
		return 0, "invalid-token", err
	} else if claims.StandardClaims.ExpiresAt != 0 && claims.ExpiresAt < time.Now().Unix() {
		return 0, "expired", err
	}
	return claims.IDUser, "valid", err
}
