package logic

import (
	"context"

	order "github.com/JBurton26/msoa-go/protos/order"
	hclog "github.com/hashicorp/go-hclog"
)

// Order in the room
type Order struct {
	log hclog.Logger
}

// NewOrder Inherits from Joy Division
func NewOrder(l hclog.Logger) *Order {
	return &Order{l}
}

// MakeOrder makes an order
func (o *Order) MakeOrder(ctx context.Context, or *order.OrderRequest) (*order.OrderResponse, error) {
	return &order.OrderResponse{Body: "Hello World!"}, nil
}
