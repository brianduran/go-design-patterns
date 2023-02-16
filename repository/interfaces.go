package repository

import "context"

// IRepository defines the methods handle a user.
type IRepository interface {
	CreateUser(ctx context.Context, name string, age int) error
	DeleteUser(ctx context.Context, id int) error
	GetUser(ctx context.Context, id int) (*User, error)
	UpdateUser(ctx context.Context, id int, attributes []interface{}) error
}
