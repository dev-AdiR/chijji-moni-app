package domain

import (
	"time"
)

type User struct {
	Id        int8      `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type UserRepo interface {
	BaseDomain[User]
	GetByEmail() (User, error)
	Fetch(username string) (*User, error)
}
