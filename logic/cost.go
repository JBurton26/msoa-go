package logic

import (
	"context"

	cost "github.com/JBurton26/msoa-go/protos/cost"
	hclog "github.com/hashicorp/go-hclog"
)

// Cost Blah Blah Blah
type Cost struct {
	log hclog.Logger
}

// NewCost plpl
func NewCost(l hclog.Logger) *Cost {
	return &Cost{l}
}

// GetUnitCost Blah Blah Blah
// TODO Currency Conversions?
func (c *Cost) GetUnitCost(ctx context.Context, cr *cost.CostRequest) (*cost.CostResponse, error) {
	c.log.Info("Handle GetUnitCost", "id", cr.GetID())
	return &cost.CostResponse{Price: 5.99}, nil
}

// TotalBasket Blah Blah Blah
// TODO Currency Conversions?
func (c *Cost) TotalBasket(ctx context.Context, bask *cost.Basket) (*cost.CostResponse, error) {
	c.log.Info("Handle TotalBasket", "Basket", bask.String())
	totalPrice := 0.0
	//req := cost.CostRequest{}

	for _, item := range bask.GetItems() {
		c.log.Info("T", "ID", item.GetID())
		itemResp, _ := c.GetUnitCost(ctx, &cost.CostRequest{ID: item.GetID()})
		totalPrice += (itemResp.Price * float64(item.Count))
	}
	c.log.Info("Basket", "Items", bask.GetItems()[0])
	return &cost.CostResponse{Price: totalPrice}, nil
}
