package repository

import "context"

type Repository interface {
	CreateUser(ctx context.Context, name string, age int) error
	GetUser(ctx context.Context, id int) (*User, error)
	UpdateUser(ctx context.Context, id int, attributes []interface{}) error
	DeleteUser(ctx context.Context, id int) error
}
