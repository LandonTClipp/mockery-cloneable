package main

import (
	"clone/interfaces"
	"clone/mocks"
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

type testCleanup struct{}

func (*testCleanup) Logf(format string, args ...interface{}) {}

func (*testCleanup) Errorf(format string, args ...interface{}) {}

func (*testCleanup) FailNow() {}

func (t *testCleanup) Cleanup(func()) {}

func main() {
	repository := &repository{}
	Foo(repository)

	mockRepository := mocks.NewMockRepository(&testCleanup{})
	Foo(mockRepository)
}
