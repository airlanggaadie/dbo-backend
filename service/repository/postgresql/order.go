package postgresql

import (
	"context"
	"database/sql"
	"dbo/assignment-test/model"
	"dbo/assignment-test/service"
	"fmt"
	"log"
	"strings"

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

// GetOrderByOrderNumber implements service.OrderRepository.
func (o orderRepository) GetOrderByOrderNumber(ctx context.Context, code string) ([]model.SimpleOrder, error) {
	var orders []model.SimpleOrder
	rows, err := o.DB.QueryContext(ctx, queryGetSimpleOrdersByName, "%"+code+"%")
	if err != nil {
		return nil, fmt.Errorf("[postgresql][GetOrderByOrderNumber] error query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order model.SimpleOrder
		if err := rows.Scan(
			&order.Id,
			&order.Code,
		); err != nil {
			return nil, fmt.Errorf("[postgresql][GetOrderByOrderNumber] error scan: %v", err)
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// DeleteOrder implements service.OrderRepository.
func (o orderRepository) DeleteOrder(ctx context.Context, id uuid.UUID) error {
	tx, err := o.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("[postgresql][DeleteOrder] begin transaction error: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, queryDeleteOrderItemByOrderId, id)
	if err != nil {
		return fmt.Errorf("[postgresql][DeleteOrder] execution delete order item error: %w", err)
	}

	_, err = tx.ExecContext(ctx, queryDeleteOrder, id)
	if err != nil {
		return fmt.Errorf("[postgresql][DeleteOrder] execution delete order error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgresql][DeleteOrder] commit error: %w", err)
	}

	return nil
}

// GetOrder implements service.OrderRepository.
func (o orderRepository) GetOrder(ctx context.Context, id uuid.UUID) (model.OrderDetail, error) {
	var order model.Order
	if err := o.DB.QueryRowContext(ctx, queryGetOrderById, id).Scan(
		&order.Id,
		&order.Code,
		&order.BuyerName,
		&order.ItemQuantity,
		&order.TotalPrice,
		&order.CreatedAt,
		&order.UpdatedAt,
	); err != nil {
		return model.OrderDetail{}, fmt.Errorf("[postgresql][GetOrder] error query: %w", err)
	}

	rows, err := o.DB.QueryContext(ctx, queryGetOrderItemByOrderId, order.Id)
	if err != nil {
		return model.OrderDetail{}, fmt.Errorf("[postgresql][GetOrder] order item error query: %w", err)
	}
	defer rows.Close()

	var orderItems []model.OrderItem
	for rows.Next() {
		var orderItem model.OrderItem
		if err := rows.Scan(
			&orderItem.Id,
			&orderItem.OrderId,
			&orderItem.Code,
			&orderItem.Name,
			&orderItem.Quantity,
			&orderItem.UnitPrice,
			&orderItem.TotalPrice,
			&orderItem.CreatedAt,
			&orderItem.UpdatedAt,
		); err != nil {
			return model.OrderDetail{}, fmt.Errorf("[postgresql][GetOrder] order item error scan: %w", err)
		}

		orderItems = append(orderItems, orderItem)
	}

	return model.OrderDetail{
		Order: order,
		Items: orderItems,
	}, nil
}

// GetOrdersPaginate implements service.OrderRepository.
func (o orderRepository) GetOrdersPaginate(ctx context.Context, search string, offset int, limit int) ([]model.Order, int64, error) {
	var (
		query *sql.Rows
		err   error
	)

	if search == "" {
		query, err = o.DB.QueryContext(ctx, queryGetOrdersPaginate, offset, limit)
	} else {
		query, err = o.DB.QueryContext(ctx, queryGetOrdersByNamePaginate, "%"+search+"%", offset, limit)
	}

	if err != nil {
		return nil, 0, fmt.Errorf("[postgresql][GetOrdersPaginate] error query: %v", err)
	}
	defer query.Close()

	var orders []model.Order
	for query.Next() {
		var order model.Order
		if err := query.Scan(
			&order.Id,
			&order.Code,
			&order.BuyerName,
			&order.ItemQuantity,
			&order.TotalPrice,
			&order.CreatedAt,
			&order.UpdatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("[postgresql][GetOrdersPaginate] error scan: %v", err)
		}

		orders = append(orders, order)
	}

	var total int64
	if search != "" {
		if err := o.DB.QueryRowContext(ctx, queryGetOrdersCountByCode, "%"+search+"%").Scan(&total); err != nil {
			return nil, 0, fmt.Errorf("[postgresql][GetOrdersPaginate] error count: %v", err)
		}
	} else {
		if err := o.DB.QueryRowContext(ctx, queryGetOrdersCount).Scan(&total); err != nil {
			return nil, 0, fmt.Errorf("[postgresql][GetOrdersPaginate] error count: %v", err)
		}
	}

	return orders, total, nil
}

// InsertOrder implements service.OrderRepository.
func (o orderRepository) InsertOrder(ctx context.Context, order model.OrderDetail) (model.OrderDetail, error) {
	tx, err := o.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return model.OrderDetail{}, fmt.Errorf("[postgresql][InsertOrder] begin transaction error: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, queryInsertOrder, order.Id, order.Code, order.BuyerName, order.ItemQuantity, order.TotalPrice)
	if err != nil {
		return model.OrderDetail{}, fmt.Errorf("[postgresql][InsertOrder] execution error: %w", err)
	}

	log.Println(order)
	log.Println(order.Items)

	var (
		queryOrderItemBuilder strings.Builder
		queryOrderItemArgs    []interface{}
	)
	queryOrderItemBuilder.WriteString(queryInsertOrderItem)
	paramCount := 0
	for i, item := range order.Items {
		log.Println("itemm: ", item)
		param1 := paramCount + 1
		param2 := param1 + 1
		param3 := param2 + 1
		param4 := param3 + 1
		param5 := param4 + 1
		param6 := param5 + 1
		param7 := param6 + 1
		paramCount = param7
		queryOrderItemBuilder.WriteString(fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d)", param1, param2, param3, param4, param5, param6, param7))
		if i != len(order.Items)-1 {
			queryOrderItemBuilder.WriteString(", ")
		}

		queryOrderItemArgs = append(queryOrderItemArgs, item.Id, item.OrderId, item.Code, item.Name, item.Quantity, item.UnitPrice, item.TotalPrice)
	}

	log.Println(queryOrderItemBuilder.String())
	log.Println(queryOrderItemArgs)
	_, err = tx.ExecContext(ctx, queryOrderItemBuilder.String(), queryOrderItemArgs...)
	if err != nil {
		return model.OrderDetail{}, fmt.Errorf("[postgresql][InsertOrder] items execution error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return model.OrderDetail{}, fmt.Errorf("[postgresql][InsertOrder] commit error: %w", err)
	}

	return order, nil
}

// UpdateOrder implements service.OrderRepository.
func (o orderRepository) UpdateOrder(ctx context.Context, id uuid.UUID, buyerName string) error {
	tx, err := o.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("[postgresql][UpdateOrder] begin transaction error: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, queryUpdateOrder, id, buyerName)
	if err != nil {
		return fmt.Errorf("[postgresql][UpdateOrder] execution error: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("[postgresql][UpdateOrder] commit error: %w", err)
	}

	return nil
}
