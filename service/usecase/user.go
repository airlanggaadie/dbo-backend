package usecase

import (
	"context"
	"dbo/assignment-test/model"
	"dbo/assignment-test/service"

	"github.com/google/uuid"
)

func NewUserUsecase(userRepository service.UserRepository) service.UserUsecase {
	return user{
		userRepository: userRepository,
	}
}

type user struct {
	userRepository service.UserRepository
}

// AddNewUser implements service.UserUsecase.
func (u user) AddNewUser(ctx context.Context, request model.NewUserRequest) (model.UserDetailResponse, error) {
	panic("unimplemented")
}

// DeleteUser implements service.UserUsecase.
func (u user) DeleteUser(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetUser implements service.UserUsecase.
func (u user) GetUser(ctx context.Context, id uuid.UUID) (model.UserDetailResponse, error) {
	panic("unimplemented")
}

// GetUsersPaginate implements service.UserUsecase.
func (u user) GetUsersPaginate(ctx context.Context, search string, page int, limit int) (model.ListUserPaginateResponse, error) {
	panic("unimplemented")
}

// SearchUser implements service.UserUsecase.
func (u user) SearchUser(ctx context.Context, keyword string) (model.UserListResponse, error) {
	panic("unimplemented")
}

// UpdateUser implements service.UserUsecase.
func (u user) UpdateUser(ctx context.Context, id uuid.UUID, request model.NewUserRequest) (model.UserDetailResponse, error) {
	panic("unimplemented")
}
