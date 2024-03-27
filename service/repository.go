package service

import (
	"context"
	"dbo/assignment-test/model"

	"github.com/google/uuid"
)

type AuthRepository interface {
	InsertLoginData(ctx context.Context, userId uuid.UUID) error
	GetLoginData(ctx context.Context, offset, limit int) ([]model.Auth, error)
}

type UserRepository interface {
	// GetUsersPaginate return a paginated list of users
	GetUsersPaginate(ctx context.Context, search string, offset, limit int) ([]model.User, int64, error)

	// GetUser returns a single user
	GetUser(ctx context.Context, id uuid.UUID) (model.User, error)

	// GetUserByName returns a list of users that match with the given searchName
	GetUserByName(ctx context.Context, searchName string) ([]model.SimpleUser, error)

	// InsertUser inserts a new user in the database
	InsertUser(ctx context.Context, user model.User, userPassword model.UserPassword) (model.User, error)

	// UpdateUser updates an existing user in the database
	UpdateUser(ctx context.Context, newUser model.User) (model.User, error)

	// DeleteUser deletes a user from the database
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type OrderRepository interface {
	// GetOrdersPaginate return a paginated list of orders
	GetOrdersPaginate(ctx context.Context, search string, offset, limit int) ([]model.Order, int64, error)

	// GetOrder returns a single order
	GetOrder(ctx context.Context, id uuid.UUID) (model.Order, error)

	// InsertOrder inserts a new order in the database
	InsertOrder(ctx context.Context, Order model.Order) (model.Order, error)

	// UpdateOrder updates an existing order in the database
	UpdateOrder(ctx context.Context, newOrder model.Order) (model.Order, error)

	// DeleteOrder deletes a order from the database
	DeleteOrder(ctx context.Context, id uuid.UUID) error
}
