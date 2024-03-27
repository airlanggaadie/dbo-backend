package model

import (
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

type UserPassword struct {
	UserId    uuid.UUID
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserLoginHistory struct {
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ListUserPaginateResponse struct {
	// TODO: fill this field
}

type UserDetailResponse struct {
	// TODO: fill this field
}

type UserListResponse struct {
	// TODO: fill this field
}

type NewUserRequest struct {
	// TODO: fill this field
}
