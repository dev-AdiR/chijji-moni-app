package utils

import (
	"chijji-moni-backend-go/domain"
	"errors"
	"fmt"
	"log"
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

type MyClaims struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func VerifyToken(env *domain.Env, tokenString string) (*MyClaims, error) {
	secret := []byte(env.JwtSecret) // usually from env

	fmt.Printf("\n\n\n %s", secret)
	claims := &MyClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		log.Println("Error parsing jwt token", err)
		return nil, err
	}

	if !token.Valid {
		log.Println("Invalid token", tokenString)
		return nil, errors.New("Invalid token")
	}

	return claims, nil
}
