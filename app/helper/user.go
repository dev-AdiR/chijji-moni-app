package helper

import (
	"chijji-moni-backend-go/domain"
)

func ValidateUserLoginRequest(username string, password string) (err error) {

	if username == "" || password == "" {
		return domain.ErrInvalidRequest
	}

	return nil
}
