package usecase

import (
	"context"
	"dbo/assignment-test/model"
	"dbo/assignment-test/service"

	"github.com/google/uuid"
)

func NewOrderUsecase(orderRepository service.OrderRepository) service.OrderUsecase {
	return order{
		orderRepository: orderRepository,
	}
}

type order struct {
	orderRepository service.OrderRepository
}

// AddNewOrder implements service.OrderUsecase.
func (o order) AddNewOrder(ctx context.Context, request model.NewOrderRequest) (model.OrderDetailResponse, error) {
	panic("unimplemented")
}

// DeleteOrder implements service.OrderUsecase.
func (o order) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetOrder implements service.OrderUsecase.
func (o order) GetOrder(ctx context.Context, id uuid.UUID) (model.OrderDetailResponse, error) {
	panic("unimplemented")
}

// GetOrdersPaginate implements service.OrderUsecase.
func (o order) GetOrdersPaginate(ctx context.Context, search string, page int, limit int) (model.ListOrderPaginateResponse, error) {
	panic("unimplemented")
}

// SearchOrder implements service.OrderUsecase.
func (o order) SearchOrder(ctx context.Context, keyword string) (model.ListOrderResponse, error) {
	panic("unimplemented")
}

// UpdateOrder implements service.OrderUsecase.
func (o order) UpdateOrder(ctx context.Context, id uuid.UUID, request model.NewOrderRequest) (model.OrderDetailResponse, error) {
	panic("unimplemented")
}
