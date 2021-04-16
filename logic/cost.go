package logic

import (
	"context"

	"cloud.google.com/go/firestore"
	cost "github.com/JBurton26/msoa-go/protos/cost"
	tools "github.com/JBurton26/msoa-go/tools"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/api/iterator"
)

// Cost Base inventory server struct, has a logger and a path
type Cost struct {
	log  hclog.Logger
	path string
}

// NewCost plpl
func NewCost(l hclog.Logger, s string) *Cost {
	return &Cost{l, s}
}

// GetUnitCost Gets the unit cost based on the id
func (c *Cost) GetUnitCost(ctx context.Context, cr *cost.CostRequest) (*cost.CostResponse, error) {
	c.log.Info("Handle GetUnitCost", "id", cr.GetID())
	app, err := tools.SetFirebase(ctx, c.log, c.path)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		c.log.Info("Error", "err2", err.Error())
		return &cost.CostResponse{}, nil
	}
	defer client.Close()

	iter := client.Collection("costs").Where("ID", "==", cr.GetID()).Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		c.log.Info("Error: Iterator", "err", err.Error())
		client.Close()
		return &cost.CostResponse{}, nil
	}
	if err != nil {
		c.log.Info("Error: Iterator", "err", err.Error())
		client.Close()
		return &cost.CostResponse{}, nil
	}
	d := doc.Data()
	c.log.Info("Thing", d)
	client.Close()
	return &cost.CostResponse{Price: d["Price"].(float64)}, nil
}

// TotalBasket Gets the total cost of a basket
func (c *Cost) TotalBasket(ctx context.Context, bask *cost.Basket) (*cost.CostResponse, error) {
	c.log.Info("Handle TotalBasket", "Basket", bask.GetItems())
	totalPrice := 0.0
	for _, item := range bask.GetItems() {
		c.log.Info("T", "ID", item.GetID())
		itemResp, _ := c.GetUnitCost(ctx, &cost.CostRequest{ID: item.GetID()})
		totalPrice += (itemResp.Price * float64(item.Count))
	}
	c.log.Info("Basket", "Items", bask.GetItems()[0])
	return &cost.CostResponse{Price: totalPrice}, nil
}

// UpdateUnitCost Updates a Unit Cost
func (c *Cost) UpdateUnitCost(ctx context.Context, ar *cost.UpdateRequest) (*cost.UpdateResponse, error) {
	c.log.Info("Handle UpdateUnitCost", "ID", ar.GetID(), "New Cost", ar.GetPrice())

	app, err := tools.SetFirebase(ctx, c.log, c.path)
	client, err := app.Firestore(context.Background())
	if err != nil {
		c.log.Info("Error: Declaring Firestore app", "err", err.Error())
	}
	defer client.Close()
	if ar.GetPrice() < 0 {
		c.log.Info("Error: UpdateUnitCost", "Error", "Item cost cannot be negative")
		client.Close()
		return &cost.UpdateResponse{Success: false}, nil
	}
	iter := client.Collection("costs").Where("ID", "==", ar.GetID()).Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		c.log.Info("Error: UpdateUnitCost", "Error", "No Item With This ID")
		client.Close()
		return &cost.UpdateResponse{Success: false}, nil
	}
	if err != nil {
		c.log.Info("Error: UpdateUnitCost", "Error", err)
		client.Close()
		return &cost.UpdateResponse{Success: false}, nil
	}
	_, err = client.Collection("inventory").Doc(doc.Ref.ID).Update(ctx, []firestore.Update{
		{
			Path:  "Price",
			Value: ar.GetPrice(),
		},
	})
	client.Close()
	c.log.Info("Success: UpdateUnitCost", "Value after operation", ar.GetPrice())
	return &cost.UpdateResponse{Success: true}, nil
}

// AddUnitCost Updates a Unit Cost
func (c *Cost) AddUnitCost(ctx context.Context, ar *cost.UpdateRequest) (*cost.UpdateResponse, error) {
	c.log.Info("Handle AddUnitCost", "ID", ar.GetID(), "New Cost", ar.GetPrice())
	app, err := tools.SetFirebase(ctx, c.log, c.path)
	client, err := app.Firestore(context.Background())
	if err != nil {
		c.log.Info("Error: Declaring Firestore app", "err", err.Error())
	}
	defer client.Close()
	if ar.GetPrice() < 0 {
		c.log.Info("Error: AddUnitCost", "Error", "Item cost cannot be negative")
		client.Close()
		return &cost.UpdateResponse{Success: false}, nil
	}
	iter := client.Collection("costs").Where("ID", "==", ar.GetID()).Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		_, err = client.Collection("costs").Doc(ar.GetID()).Set(ctx, map[string]interface{}{
			"ID":    ar.GetID(),
			"Price": ar.GetPrice(),
		})
		c.log.Info("Success: AddUnitCost", "Stock Cost Added")
		client.Close()
		return &cost.UpdateResponse{Success: true}, nil
	}
	if err != nil {
		c.log.Info("Error: AddUnitCost", "Error", err)
		client.Close()
		return &cost.UpdateResponse{Success: false}, nil
	}
	c.log.Info("Error: AddUnitCost", "Error", "Item already exists with this id", "ID", doc.Data()["ID"], "Price", doc.Data()["Price"])
	client.Close()
	return &cost.UpdateResponse{Success: false}, nil
}
