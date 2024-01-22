package interfaces

// Interfaces
type Repository[R any, C Client[C]] interface {
	Clone(client C) R
	Transaction(fn func(client C) error) error
}

type Client[C any] interface {
	Transaction(fn func(client C) error) error
}
