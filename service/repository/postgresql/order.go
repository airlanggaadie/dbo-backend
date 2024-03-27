package postgresql

import (
	"context"
	"database/sql"
	"dbo/assignment-test/model"
	"dbo/assignment-test/service"

	"github.com/google/uuid"
)

func NewOrderRepository(db *sql.DB) service.OrderRepository {
	return orderRepository{
		DB: db,
	}
}

type orderRepository struct {
	DB *sql.DB
}

// DeleteOrder implements service.OrderRepository.
func (o orderRepository) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetOrder implements service.OrderRepository.
func (o orderRepository) GetOrder(ctx context.Context, id uuid.UUID) (model.Order, error) {
	panic("unimplemented")
}

// GetOrdersPaginate implements service.OrderRepository.
func (o orderRepository) GetOrdersPaginate(ctx context.Context, search string, offset int, limit int) ([]model.Order, int64, error) {
	panic("unimplemented")
}

// InsertOrder implements service.OrderRepository.
func (o orderRepository) InsertOrder(ctx context.Context, Order model.Order) (model.Order, error) {
	panic("unimplemented")
}

// UpdateOrder implements service.OrderRepository.
func (o orderRepository) UpdateOrder(ctx context.Context, newOrder model.Order) (model.Order, error) {
	panic("unimplemented")
}
