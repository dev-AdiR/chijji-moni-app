package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePasswordHash(hashedPassword []byte, plainTextPassword []byte) (err error) {
	return bcrypt.CompareHashAndPassword(hashedPassword, plainTextPassword)
}

func HashPassowrd(plainTextPassword string) (hashedPassword []byte, err error) {
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(plainTextPassword), 10)

	if err != nil {
		return nil, err
	}

	return hashedPassword, nil
}
