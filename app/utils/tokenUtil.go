package utils

import (
	"chijji-moni-backend-go/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(env *domain.Env, user *domain.User) (accessToken string, err error) {

	secret := []byte(env.JwtSecret) // usually from env

	claims := jwt.MapClaims{
		"userId":   user.Id,
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
