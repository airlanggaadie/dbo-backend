package postgresql

import (
	"context"
	"database/sql"
	"dbo/assignment-test/model"
	"dbo/assignment-test/service"

	"github.com/google/uuid"
)

func NewUserRepository(db *sql.DB) service.UserRepository {
	return userRepository{
		DB: db,
	}
}

type userRepository struct {
	DB *sql.DB
}

// DeleteUser implements service.UserRepository.
func (u userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetUser implements service.UserRepository.
func (u userRepository) GetUser(ctx context.Context, id uuid.UUID) (model.User, error) {
	panic("unimplemented")
}

// GetUsersPaginate implements service.UserRepository.
func (u userRepository) GetUsersPaginate(ctx context.Context, search string, offset int, limit int) ([]model.User, int64, error) {
	panic("unimplemented")
}

// InsertUser implements service.UserRepository.
func (u userRepository) InsertUser(ctx context.Context, User model.User) (model.User, error) {
	panic("unimplemented")
}

// UpdateUser implements service.UserRepository.
func (u userRepository) UpdateUser(ctx context.Context, newUser model.User) (model.User, error) {
	panic("unimplemented")
}
