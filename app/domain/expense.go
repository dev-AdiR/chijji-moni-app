package domain

type Expense struct {
	Id     int8 `json:"id" db:"id"`
	UserId int  `json:"user_id" db:"user_id"`
}

type ExpenseRepo interface {
	Fetch(username int) (*[]Expense, error)
}
