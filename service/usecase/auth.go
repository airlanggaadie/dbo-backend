package usecase

import (
	"context"
	"dbo/assignment-test/model"
	"dbo/assignment-test/service"
)

func NewAuthUsecase(authRepository service.AuthRepository, userRepository service.UserRepository) service.AuthUsecase {
	return auth{
		authRepository: authRepository,
		userRepository: userRepository,
	}
}

type auth struct {
	authRepository service.AuthRepository
	userRepository service.UserRepository
}

// Login implements service.AuthUsecase.
func (a auth) Login(ctx context.Context, username string, password string) (model.LoginResponse, error) {
	panic("unimplemented")
}

// Report implements service.AuthUsecase.
func (a auth) Report(ctx context.Context, page int, limit int) (model.LoginReport, error) {
	panic("unimplemented")
}
