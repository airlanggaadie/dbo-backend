package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id           uuid.UUID `json:"id"`
	Code         string    `json:"code"`
	BuyerName    string    `json:"buyer_name"`
	ItemQuantity int64     `json:"item_quantity"`
	TotalPrice   int64     `json:"total_price"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewOrder(order NewOrderRequest) (OrderDetail, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return OrderDetail{}, fmt.Errorf("[model][NewOrder] uuid order error: %v", err)
	}

	var totalQuantity int64 = 0
	var totalPrice int64 = 0
	var items []OrderItem
	for _, item := range order.Items {
		itemId, err := uuid.NewRandom()
		if err != nil {
			return OrderDetail{}, fmt.Errorf("[model][NewOrder] uuid item error: %v", err)
		}

		items = append(items, OrderItem{
			Id:         itemId,
			OrderId:    id,
			Code:       item.Code,
			Name:       item.Name,
			Quantity:   item.Quantity,
			UnitPrice:  item.UnitPrice,
			TotalPrice: item.GetTotalPrice(),
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		})

		totalQuantity += item.Quantity
		totalPrice += item.GetTotalPrice()
	}

	return OrderDetail{
		Order: Order{
			Id:           id,
			Code:         "DBO" + time.Now().Format("20060102150405"),
			BuyerName:    order.BuyerName,
			ItemQuantity: totalQuantity,
			TotalPrice:   totalPrice,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		Items: items,
	}, nil
}

type OrderItem struct {
	Id         uuid.UUID `json:"id"`
	OrderId    uuid.UUID `json:"order_id"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	Quantity   int64     `json:"quantity"`
	UnitPrice  int64     `json:"unit_price"`
	TotalPrice int64     `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderDetail struct {
	Order
	Items []OrderItem `json:"items"`
}

type SimpleOrder struct {
	Id   uuid.UUID `json:"id"`
	Code string    `json:"code"`
}

type ListOrderPaginateResponse struct {
	Data  []Order `json:"data"`
	Total int64   `json:"total"`
}

type OrderDetailResponse struct {
	OrderDetail
}

type ListOrderResponse struct {
	Data []SimpleOrder `json:"data"`
}

type NewOrderRequest struct {
	BuyerName string           `json:"buyer_name"`
	Items     []NewItemRequest `json:"items"`
}

type NewItemRequest struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Quantity  int64  `json:"quantity"`
	UnitPrice int64  `json:"unit_price"`
}

func (i *NewItemRequest) GetTotalPrice() int64 {
	return i.Quantity * i.UnitPrice
}

type UpdateOrderRequest struct {
	BuyerName string `json:"buyer_name"`
}
