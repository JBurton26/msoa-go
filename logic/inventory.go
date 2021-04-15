package logic

import (
	"context"

	inv "github.com/JBurton26/msoa-go/protos/inventory"
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
func (i *Inventory) GetStock(ctx context.Context, lr *inv.LevelRequest) (*inv.StockItem, error) {
	i.log.Info("Handle GetStock", "id", lr.GetID(), "location", lr.GetLocation())
	stock := inv.StockItem{}
	return &stock, nil
}

// ChangeStock Blah Blah Blah
// TODO Access and amend from db?
func (i *Inventory) ChangeStock(ctx context.Context, cr *inv.AmendRequest) (*inv.AmendResponse, error) {
	i.log.Info("Handle ChangeStock", "id", cr.GetID(), "location", cr.GetLocation(), "change amount", cr.GetAmount())
	return &inv.AmendResponse{Response: "Placeholder Response: Hello World!"}, nil
}

// CheckShort Blah Blah Blah
// TODO Access and amend from db?
func (i *Inventory) CheckShort(ctx context.Context, cr *inv.ShortRequest) (*inv.ShortList, error) {
	i.log.Info("Handle ChangeStock", "location", cr.GetLocation())
	short := inv.ShortList{}
	return &short, nil
}
