package main

import (
	"clone/interfaces"
	"clone/mocks"
	"testing"
)

// Interface implementation
type client struct{}

func (c *client) Transaction(cb func(client *client) error) error {
	return cb(c)
}

type repository struct {
	client *client
}

func (r *repository) Clone(client *client) *repository {
	return &repository{client: client}
}

func (r *repository) Transaction(cb func(client *client) error) error {
	return r.client.Transaction(cb)
}

// Usage
func Foo[C interfaces.Client[C], R interfaces.Repository[R, C]](repo R) {}

func main() {
	repository := &repository{}
	Foo(repository)

	var t *testing.T
	mockRepository := mocks.NewMockRepository(t)
	Foo(mockRepository)
}
