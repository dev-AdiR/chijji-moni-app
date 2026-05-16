package domain

type BaseDomain[T any] interface {
	FetchAll() ([]T, error)
	GetById() (T, error)
}
