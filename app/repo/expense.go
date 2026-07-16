package repo

import (
	"chijji-moni-backend-go/domain"
	"encoding/json"
	"errors"
)

type ExpenseRepo struct {
	Client domain.DB
}

func RegisterExpenseRepo(client domain.DB) domain.ExpenseRepo {
	return &ExpenseRepo{
		Client: client,
	}
}

func (er *ExpenseRepo) Fetch(userId int) (*[]domain.Expense, error) {
	data, err := er.Client.SelectMany("expenses", []string{"id, user_id"}, map[string]int{"user_id": 3})

	if err != nil {
		return nil, err
	}

	var result []domain.Expense

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, errors.New("Failed to parse response")
	}
	return &result, nil
}
