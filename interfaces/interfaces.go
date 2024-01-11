package interfaces

// Interfaces
type Repository[T any, C Client[C]] interface {
	Clone(client C) T
	Transaction(fn func(client C) error) error
}

type Client[C any] interface {
	Transaction(fn func(client C) error) error
}
