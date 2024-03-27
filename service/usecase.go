package service

import (
	"context"
	"dbo/assignment-test/model"

	"github.com/google/uuid"
)

type AuthUsecase interface {
	Login(ctx context.Context, username string, password string) (model.LoginResponse, error)
	Report(ctx context.Context, page, limit int) (model.LoginReport, error)
}

type UserUsecase interface {
	// GetUsersPaginate returns a list of users with total data
	GetUsersPaginate(ctx context.Context, search string, page, limit int) (model.ListUserPaginateResponse, error)

	// GetUser returns a single user. it will return [ErrNotFound] when the user is not exist in the database
	GetUser(ctx context.Context, id uuid.UUID) (model.UserDetailResponse, error)

	// GetUser returns a list of user.
	SearchUser(ctx context.Context, keyword string) (model.UserListResponse[model.SimpleUser], error)

	// AddNewUser adds a new user to the list of users
	AddNewUser(ctx context.Context, request model.NewUserRequest) (model.UserDetailResponse, error)

	// UpdateUser updates an existing user in the database. it will return [ErrNotFound] when the user is not exist in the database
	UpdateUser(ctx context.Context, id uuid.UUID, request model.NewUserRequest) (model.UserDetailResponse, error)

	// DeleteUser deletes a user from the database. it will return [ErrNotFound] when the user is not exist in the database
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type OrderUsecase interface {
	// GetOrdersPaginate returns a list of orders with total data
	GetOrdersPaginate(ctx context.Context, search string, page, limit int) (model.ListOrderPaginateResponse, error)

	// GetOrder returns a single order. it will return [ErrNotFound] when the order is not exist in the database
	GetOrder(ctx context.Context, id uuid.UUID) (model.OrderDetailResponse, error)

	// GetOrder returns a list of orders.
	SearchOrder(ctx context.Context, keyword string) (model.ListOrderResponse, error)

	// AddNewOrder adds a new order to the list of orders
	AddNewOrder(ctx context.Context, request model.NewOrderRequest) (model.OrderDetailResponse, error)

	// UpdateOrder updates an existing order in the database. it will return [ErrNotFound] when the order is not exist in the database
	UpdateOrder(ctx context.Context, id uuid.UUID, request model.NewOrderRequest) (model.OrderDetailResponse, error)

	// DeleteOrder deletes a order from the database. it will return [ErrNotFound] when the order is not exist in the database
	DeleteOrder(ctx context.Context, id uuid.UUID) error
}
