package repository

import "context"

// IRepository defines the methods handle a user.
type IRepository interface {
	CreateUser(ctx context.Context, name string, age int) error
	DeleteUserByName(ctx context.Context, name string) error
	GetUserByName(ctx context.Context, name string) (*User, error)
	UpdateUser(ctx context.Context, name string, attributes map[string]interface{}) error
}
