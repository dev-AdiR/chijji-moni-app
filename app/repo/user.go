package repo

import (
	"chijji-moni-backend-go/domain"
	"encoding/json"
	"errors"
)

func RegisterUserRepo(client domain.DB) domain.UserRepo {
	return &UserRepo{
		Client: client,
	}
}

type UserRepo struct {
	Client domain.DB
}

func (ur *UserRepo) Create(username string, hashedPassword string) error {

	err := ur.Client.Insert("users", map[string]any{
		"username": username,
		"password": hashedPassword,
	}, false)

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepo) Fetch(username string) (*domain.User, error) {

	result := new(domain.User)

	data, err := ur.Client.Select("users", []string{"*"}, map[string]string{"username": username})

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, result); err != nil {
		return nil, errors.New("Failed to parse response")
	}

	return result, nil
}

func (ur *UserRepo) GetById() (domain.User, error) {
	return domain.User{}, nil
}

func (ur *UserRepo) GetByEmail() (domain.User, error) {
	return domain.User{}, nil
}

func (ur *UserRepo) FetchAll() ([]domain.User, error) {
	return []domain.User{}, nil
}
