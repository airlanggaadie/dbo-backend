package usecase

import (
	"context"
	"database/sql"
	"dbo/assignment-test/model"
	"dbo/assignment-test/service"
	"errors"
	"fmt"

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
	// TODO: check username exists

	newUser, err := model.NewUser(request)
	if err != nil {
		return model.UserDetailResponse{}, fmt.Errorf("[usecase][AddNewUser] new user error: %v", err)
	}

	newUserPassword, err := model.NewUserPassword(newUser.Id)
	if err != nil {
		return model.UserDetailResponse{}, fmt.Errorf("[usecase][AddNewUser] new user password error: %v", err)
	}

	user, err := u.userRepository.InsertUser(ctx, newUser, newUserPassword)
	if err != nil {
		return model.UserDetailResponse{}, fmt.Errorf("[usecase][AddNewUser] insert error: %v", err)
	}

	return model.UserDetailResponse{
		User: user,
	}, nil
}

// DeleteUser implements service.UserUsecase.
func (u user) DeleteUser(ctx context.Context, id uuid.UUID) error {
	user, err := u.userRepository.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrNotFound
		}
		return fmt.Errorf("[usecase][DeleteUser] error: %v", err)
	}

	err = u.userRepository.DeleteUser(ctx, user.Id)
	if err != nil {
		return fmt.Errorf("[usecase][DeleteUser] error delete: %v", err)
	}

	return nil
}

// GetUser implements service.UserUsecase.
func (u user) GetUser(ctx context.Context, id uuid.UUID) (model.UserDetailResponse, error) {
	user, err := u.userRepository.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserDetailResponse{}, model.ErrNotFound
		}
		return model.UserDetailResponse{}, fmt.Errorf("[usecase][GetUser] error: %v", err)
	}

	return model.UserDetailResponse{
		User: user,
	}, nil
}

// GetUsersPaginate implements service.UserUsecase.
func (u user) GetUsersPaginate(ctx context.Context, search string, page int, limit int) (model.ListUserPaginateResponse, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	users, total, err := u.userRepository.GetUsersPaginate(ctx, search, offset, limit)
	if err != nil {
		return model.ListUserPaginateResponse{}, fmt.Errorf("[usecase][GetUsersPaginate] error: %v", err)
	}

	return model.ListUserPaginateResponse{
		Data:  users,
		Total: total,
	}, nil
}

// SearchUser implements service.UserUsecase.
func (u user) SearchUser(ctx context.Context, keyword string) (model.UserListResponse[model.SimpleUser], error) {
	users, err := u.userRepository.GetUserByName(ctx, keyword)
	if err != nil {
		return model.UserListResponse[model.SimpleUser]{}, fmt.Errorf("[usecase][SearchUser] error: %v", err)
	}

	return model.UserListResponse[model.SimpleUser]{
		Data: users,
	}, nil
}

// UpdateUser implements service.UserUsecase.
func (u user) UpdateUser(ctx context.Context, id uuid.UUID, request model.NewUserRequest) (model.UserDetailResponse, error) {
	updatedMovie, err := u.userRepository.UpdateUser(ctx, model.NewUserUpdate(id, request))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.UserDetailResponse{}, model.ErrNotFound
		}

		return model.UserDetailResponse{}, fmt.Errorf("[usecase][UpdateMovie] error update movie: %v", err)
	}

	return model.UserDetailResponse{
		User: updatedMovie,
	}, nil
}
