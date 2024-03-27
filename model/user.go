package model

import (
	"dbo/assignment-test/service/repository/utils"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(request NewUserRequest) (User, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return User{}, fmt.Errorf("[model][NewUser] uuid error: %v", err)
	}

	return User{
		Id:        id,
		Username:  request.Username,
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

type UserPassword struct {
	UserId    uuid.UUID
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserPassword(id uuid.UUID) (UserPassword, error) {
	defaultPassword := "verysecret"
	hash, err := utils.HashPassword(defaultPassword)
	if err != nil {
		return UserPassword{}, fmt.Errorf("[model][NewUserPassword] hash error: %v", err)
	}

	return UserPassword{
		UserId:    id,
		Password:  hash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func NewUserUpdate(id uuid.UUID, request NewUserRequest) User {
	return User{
		Id:       id,
		Username: request.Username,
		Name:     request.Name,
	}
}

type SimpleUser struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UserLoginHistory struct {
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLoginHistoryDetail struct {
	UserId    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

type ListUserPaginateResponse struct {
	Data  []User `json:"data"`
	Total int64  `json:"total"`
}

type UserDetailResponse struct {
	User
}

type UserListResponse[T User | SimpleUser] struct {
	Data []T `json:"data"`
}

type NewUserRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
