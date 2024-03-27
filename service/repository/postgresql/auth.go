package postgresql

import (
	"context"
	"database/sql"
	"dbo/assignment-test/model"
	"dbo/assignment-test/service"

	"github.com/google/uuid"
)

func NewAuthRepository(db *sql.DB) service.AuthRepository {
	return authRepository{
		DB: db,
	}
}

type authRepository struct {
	DB *sql.DB
}

// GetLoginData implements service.AuthRepository.
func (a authRepository) GetLoginData(ctx context.Context, offset int, limit int) ([]model.Auth, error) {
	panic("unimplemented")
}

// InsertLoginData implements service.AuthRepository.
func (a authRepository) InsertLoginData(ctx context.Context, userId uuid.UUID) error {
	panic("unimplemented")
}
