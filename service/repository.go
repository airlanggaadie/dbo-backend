package service

import (
	"context"
	"dbo/assignment-test/model"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	// GetUsersPaginate return a paginated list of users
	GetUsersPaginate(ctx context.Context, search string, offset, limit int) ([]model.User, int64, error)

	// GetUser returns a single user
	GetUser(ctx context.Context, id uuid.UUID) (model.User, error)

	// GetUserByUsername returns a single user
	GetUserByUsername(ctx context.Context, username string) (model.User, error)

	// GetPasswordByUserId returns a hashed password
	GetPasswordByUserId(ctx context.Context, userId uuid.UUID) (string, error)

	// GetUserByName returns a list of users that match with the given searchName
	GetUserByName(ctx context.Context, searchName string) ([]model.SimpleUser, error)

	// InsertUser inserts a new user in the database
	InsertUser(ctx context.Context, user model.User, userPassword model.UserPassword) (model.User, error)

	// UpdateUser updates an existing user in the database
	UpdateUser(ctx context.Context, newUser model.User) (model.User, error)

	// DeleteUser deletes a user from the database
	DeleteUser(ctx context.Context, id uuid.UUID) error

	// InsertUserLoginHistory inserts a new user login history in the database
	InsertUserLoginHistory(ctx context.Context, userId uuid.UUID) error

	GetUserLoginHistoryPaginate(ctx context.Context, offset, limit int) ([]model.UserLoginHistoryDetail, int64, error)
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

type JwtRepository interface {
	// Generate a JWT token
	Generate(userId uuid.UUID, additionalClaims map[string]string, expiry time.Duration) (string, error)

	// Verify a JWT token
	Verify(token string) (uuid.UUID, error)
}
