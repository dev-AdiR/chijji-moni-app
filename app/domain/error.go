package domain

import "errors"

var (
	ErrInvalidRequest = errors.New("Invalid Request")
	UserNotFound      = errors.New("User not found")
)
