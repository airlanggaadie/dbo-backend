package postgresql

import (
	"context"
	"database/sql"
	"dbo/assignment-test/model"
	"dbo/assignment-test/service"
	"fmt"

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

// GetUserLoginHistoryPaginate implements service.UserRepository.
func (u userRepository) GetUserLoginHistoryPaginate(ctx context.Context, offset int, limit int) ([]model.UserLoginHistoryDetail, int64, error) {
	query, err := u.DB.QueryContext(ctx, queryGetUserLoginHistoryDetail, offset, limit)
	if err != nil {
		return nil, 0, fmt.Errorf("[postgresql][GetUserLoginHistoryPaginate] error query: %v", err)
	}
	defer query.Close()

	var users []model.UserLoginHistoryDetail
	for query.Next() {
		var user model.UserLoginHistoryDetail
		if err := query.Scan(
			&user.UserId,
			&user.Username,
			&user.LoginTime,
		); err != nil {
			return nil, 0, fmt.Errorf("[postgresql][GetUserLoginHistoryPaginate] error scan: %v", err)
		}

		users = append(users, user)
	}

	var total int64
	if err := u.DB.QueryRowContext(ctx, queryGetUserLoginHistoryCount).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("[postgresql][GetUserLoginHistoryPaginate] error count: %v", err)
	}

	return users, total, nil
}

// InsertUserLoginHistory implements service.UserRepository.
func (u userRepository) InsertUserLoginHistory(ctx context.Context, userId uuid.UUID) error {
	tx, err := u.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("[postgresql][InsertUserLoginHistory] begin transaction error: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, queryInsertUserLoginHistory, userId)
	if err != nil {
		return fmt.Errorf("[postgresql][InsertUserLoginHistory] execution error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgresql][InsertUserLoginHistory] commit error: %w", err)
	}

	return nil
}

// GetPasswordByUserId implements service.UserRepository.
func (u userRepository) GetPasswordByUserId(ctx context.Context, userId uuid.UUID) (string, error) {
	var hashedPassword string
	if err := u.DB.QueryRowContext(ctx, queryGetPasswordByUserId, userId).Scan(
		&hashedPassword,
	); err != nil {
		return "", fmt.Errorf("[postgresql][GetPasswordByUserId] error query: %w", err)
	}

	return hashedPassword, nil
}

// GetUserByUsername implements service.UserRepository.
func (u userRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	var user model.User
	if err := u.DB.QueryRowContext(ctx, queryGetUserByUsername, username).Scan(
		&user.Id,
		&user.Username,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return user, fmt.Errorf("[postgresql][GetUserByUsername] error query: %w", err)
	}

	return user, nil
}

// GetUserByName implements service.UserRepository.
func (u userRepository) GetUserByName(ctx context.Context, searchName string) ([]model.SimpleUser, error) {
	var users []model.SimpleUser
	rows, err := u.DB.QueryContext(ctx, queryGetSimpleUsersByName, "%"+searchName+"%")
	if err != nil {
		return nil, fmt.Errorf("[postgresql][GetUserByName] error query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user model.SimpleUser
		if err := rows.Scan(
			&user.Id,
			&user.Name,
		); err != nil {
			return nil, fmt.Errorf("[postgresql][GetUserByName] error scan: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

// DeleteUser implements service.UserRepository.
func (u userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	tx, err := u.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("[postgresql][DeleteUser] begin transaction error: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, queryDeleteUserPassword, id)
	if err != nil {
		return fmt.Errorf("[postgresql][DeleteUser] execution delete password error: %w", err)
	}

	_, err = tx.ExecContext(ctx, queryDeleteUser, id)
	if err != nil {
		return fmt.Errorf("[postgresql][DeleteUser] execution delete user error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgresql][DeleteUser] commit error: %w", err)
	}

	return nil
}

// GetUser implements service.UserRepository.
func (u userRepository) GetUser(ctx context.Context, id uuid.UUID) (model.User, error) {
	var user model.User
	if err := u.DB.QueryRowContext(ctx, queryGetUserById, id).Scan(
		&user.Id,
		&user.Username,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return user, fmt.Errorf("[postgresql][GetUser] error query: %w", err)
	}

	return user, nil
}

// GetUsersPaginate implements service.UserRepository.
func (u userRepository) GetUsersPaginate(ctx context.Context, search string, offset int, limit int) ([]model.User, int64, error) {
	var (
		query *sql.Rows
		err   error
	)

	if search == "" {
		query, err = u.DB.QueryContext(ctx, queryGetUsersPaginate, offset, limit)
	} else {
		query, err = u.DB.QueryContext(ctx, queryGetUsersByNamePaginate, "%"+search+"%", offset, limit)
	}

	if err != nil {
		return nil, 0, fmt.Errorf("[postgresql][GetUsersPaginate] error query: %v", err)
	}
	defer query.Close()

	var users []model.User
	for query.Next() {
		var user model.User
		if err := query.Scan(
			&user.Id,
			&user.Username,
			&user.Name,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("[postgresql][GetUsersPaginate] error scan: %v", err)
		}

		users = append(users, user)
	}

	var total int64
	if search != "" {
		if err := u.DB.QueryRowContext(ctx, queryGetUsersCountByName, "%"+search+"%").Scan(&total); err != nil {
			return nil, 0, fmt.Errorf("[postgresql][GetUsersPaginate] error count: %v", err)
		}
	} else {
		if err := u.DB.QueryRowContext(ctx, queryGetUsersCount).Scan(&total); err != nil {
			return nil, 0, fmt.Errorf("[postgresql][GetUsersPaginate] error count: %v", err)
		}
	}

	return users, total, nil
}

// InsertUser implements service.UserRepository.
func (u userRepository) InsertUser(ctx context.Context, user model.User, userPassword model.UserPassword) (model.User, error) {
	tx, err := u.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return model.User{}, fmt.Errorf("[postgresql][InsertUser] begin transaction error: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, queryInsertUser, user.Id, user.Username, user.Name)
	if err != nil {
		return model.User{}, fmt.Errorf("[postgresql][InsertUser] execution error: %w", err)
	}

	_, err = tx.ExecContext(ctx, queryInsertUserPassword, user.Id, userPassword.Password)
	if err != nil {
		return model.User{}, fmt.Errorf("[postgresql][InsertUser] password execution error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return model.User{}, fmt.Errorf("[postgresql][InsertUser] commit error: %w", err)
	}

	return user, nil
}

// UpdateUser implements service.UserRepository.
func (u userRepository) UpdateUser(ctx context.Context, newUser model.User) (model.User, error) {
	tx, err := u.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return model.User{}, fmt.Errorf("[postgresql][UpdateUser] begin transaction error: %w", err)
	}
	defer tx.Rollback()

	var updatedUser model.User
	if err := tx.QueryRowContext(ctx, queryUpdateUser, newUser.Id, newUser.Username, newUser.Name).Scan(
		&updatedUser.Id,
		&updatedUser.Username,
		&updatedUser.Name,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	); err != nil {
		return model.User{}, fmt.Errorf("[postgresql][UpdateUser] execution error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return model.User{}, fmt.Errorf("[postgresql][UpdateUser] commit error: %w", err)
	}

	return updatedUser, nil
}
