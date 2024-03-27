package model

import (
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

type ListOrderPaginateResponse struct {
	// TODO: fill this field
}

type OrderDetailResponse struct {
	// TODO: fill this field
}

type ListOrderResponse struct {
	// TODO: fill this field
}

type NewOrderRequest struct {
	// TODO: fill this field
}
