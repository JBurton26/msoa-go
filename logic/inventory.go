package logic

import (
	"context"

	inv "github.com/JBurton26/protos/inventory"
	hclog "github.com/hashicorp/go-hclog"
)

// Inventory blahblablah
type Inventory struct {
	log hclog.Logger
}

// NewInventory new Inv
func NewInventory(l hclog.Logger) *Inventory {
	return &Inventory{l}
}

// GetStock gets the stock
func (i *Inventory) GetStock(ctx context.Context, lr *inv.LevelRequest) (*inv.LevelResponse, error) {
	i.log.Info("Handle GetStock", "id", lr.GetID(), "location", lr.GetLocation())
	return &inv.LevelResponse{StockCount: 5}, nil
}

// ChangeStock Blah Blah Blah
// TODO Access and amend from db?
func (i *Inventory) ChangeStock(ctx context.Context, cr *inv.AmendRequest) (*inv.AmendResponse, error) {
	i.log.Info("Handle ChangeStock", "id", cr.GetID(), "location", cr.GetLocation(), "change amount", cr.GetAmount())
	return &inv.AmendResponse{Response: "Placeholder Response: Hello World!"}, nil
}
