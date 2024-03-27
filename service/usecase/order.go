package usecase

import (
	"context"
	"database/sql"
	"dbo/assignment-test/model"
	"dbo/assignment-test/service"
	"errors"
	"fmt"
	"log"

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
	log.Println(request)
	newOrderDetail, err := model.NewOrder(request)
	if err != nil {
		return model.OrderDetailResponse{}, fmt.Errorf("[usecase][AddNewOrder] new order error: %v", err)
	}

	order, err := o.orderRepository.InsertOrder(ctx, newOrderDetail)
	if err != nil {
		return model.OrderDetailResponse{}, fmt.Errorf("[usecase][AddNewOrder] insert error: %v", err)
	}

	return model.OrderDetailResponse{
		OrderDetail: order,
	}, nil
}

// DeleteOrder implements service.OrderUsecase.
func (o order) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	order, err := o.orderRepository.GetOrder(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.ErrNotFound
		}
		return fmt.Errorf("[usecase][DeleteOrder] error: %v", err)
	}

	err = o.orderRepository.DeleteOrder(ctx, order.Id)
	if err != nil {
		return fmt.Errorf("[usecase][DeleteOrder] error delete: %v", err)
	}

	return nil
}

// GetOrder implements service.OrderUsecase.
func (o order) GetOrder(ctx context.Context, id uuid.UUID) (model.OrderDetailResponse, error) {
	orderDetail, err := o.orderRepository.GetOrder(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.OrderDetailResponse{}, model.ErrNotFound
		}
		return model.OrderDetailResponse{}, fmt.Errorf("[usecase][GetOrder] error: %v", err)
	}

	return model.OrderDetailResponse{
		OrderDetail: orderDetail,
	}, nil
}

// GetOrdersPaginate implements service.OrderUsecase.
func (o order) GetOrdersPaginate(ctx context.Context, search string, page int, limit int) (model.ListOrderPaginateResponse, error) {
	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	orders, total, err := o.orderRepository.GetOrdersPaginate(ctx, search, offset, limit)
	if err != nil {
		return model.ListOrderPaginateResponse{}, fmt.Errorf("[usecase][GetOrdersPaginate] error: %v", err)
	}

	return model.ListOrderPaginateResponse{
		Data:  orders,
		Total: total,
	}, nil
}

// SearchOrder implements service.OrderUsecase.
func (o order) SearchOrder(ctx context.Context, keyword string) (model.ListOrderResponse, error) {
	orders, err := o.orderRepository.GetOrderByOrderNumber(ctx, keyword)
	if err != nil {
		return model.ListOrderResponse{}, fmt.Errorf("[usecase][SearchOrder] error: %v", err)
	}

	return model.ListOrderResponse{
		Data: orders,
	}, nil
}

// UpdateOrder implements service.OrderUsecase.
func (o order) UpdateOrder(ctx context.Context, id uuid.UUID, buyerName string) (model.OrderDetailResponse, error) {
	orderDetail, err := o.orderRepository.GetOrder(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.OrderDetailResponse{}, model.ErrNotFound
		}
		return model.OrderDetailResponse{}, fmt.Errorf("[usecase][UpdateOrder] error: %v", err)
	}

	orderDetail.BuyerName = buyerName
	err = o.orderRepository.UpdateOrder(ctx, orderDetail.Id, orderDetail.BuyerName)
	if err != nil {
		return model.OrderDetailResponse{}, fmt.Errorf("[usecase][UpdateOrder] error: %v", err)
	}

	return model.OrderDetailResponse{
		OrderDetail: orderDetail,
	}, nil
}
