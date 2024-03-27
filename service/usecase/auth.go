package usecase

import (
	"context"
	"database/sql"
	"dbo/assignment-test/model"
	"dbo/assignment-test/service"
	"dbo/assignment-test/service/repository/utils"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func NewAuthUsecase(userRepository service.UserRepository, jwtRepository service.JwtRepository) service.AuthUsecase {
	return auth{
		userRepository: userRepository,
		jwtRepository:  jwtRepository,
	}
}

type auth struct {
	userRepository service.UserRepository
	jwtRepository  service.JwtRepository
}

// Login implements service.AuthUsecase.
func (a auth) Login(ctx context.Context, username string, password string) (model.LoginResponse, error) {
	user, err := a.userRepository.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.LoginResponse{}, model.ErrNotFound
		}
	}

	hashedPassword, err := a.userRepository.GetPasswordByUserId(ctx, user.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.LoginResponse{}, model.ErrNotFound
		}
	}

	if err := utils.VerifyPassword(password, hashedPassword); err != nil {
		return model.LoginResponse{}, fmt.Errorf("[usecase][auth][Login] verify password error: %w", err)
	}

	// generate token
	token, err := a.jwtRepository.Generate(user.Id, map[string]string{}, 10*time.Minute)
	if err != nil {
		return model.LoginResponse{}, fmt.Errorf("[usecase][auth][Login] generate token error: %w", err)
	}

	// insert login history
	go a.logHistory(user.Id)

	return model.LoginResponse{
		Token: token,
	}, nil
}

func (a auth) logHistory(id uuid.UUID) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := a.userRepository.InsertUserLoginHistory(ctx, id)
	if err != nil {
		log.Println("[usecase][auth][logHistory] insert login history error: %w", err)
	}
}

// Report implements service.AuthUsecase.
func (a auth) Report(ctx context.Context, page int, limit int) (model.LoginReport, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	report, total, err := a.userRepository.GetUserLoginHistoryPaginate(ctx, offset, limit)
	if err != nil {
		return model.LoginReport{}, fmt.Errorf("[usecase][GetUsersPaginate] error: %v", err)
	}

	return model.LoginReport{
		Data:  report,
		Total: total,
	}, nil
}
