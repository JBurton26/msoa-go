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
	return &cost.CostResponse{StockCost: 5.99}, nil
}
