package logic

import (
	"context"

	"google.golang.org/grpc"

	"cloud.google.com/go/firestore"
	cost "github.com/JBurton26/msoa-go/protos/cost"
	inv "github.com/JBurton26/msoa-go/protos/inventory"
	tools "github.com/JBurton26/msoa-go/tools"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
)

// Inventory Base inventory server struct, has a logger and a path
type Inventory struct {
	log  hclog.Logger
	path string
}

// NewInventory Initialiser, takes the arguments and returns a new inventory
// server with the arguments passed in
func NewInventory(l hclog.Logger, s string) *Inventory {
	return &Inventory{l, s}
}

// GetStock Gets the stock from the database associated with the ID
func (i *Inventory) GetStock(ctx context.Context, lr *inv.LevelRequest) (*inv.StockItem, error) {
	i.log.Info("Handle GetStock", "id", lr.GetID(), "location", lr.GetLocation())
	var stock inv.StockItem
	app, err := tools.SetFirebase(ctx, i.log, i.path)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		i.log.Info("Error", "err2", err.Error())
	}
	defer client.Close()

	iter := client.Collection("inventory").Where("ID", "==", lr.GetID()).Where("Location", "==", lr.GetLocation()).Documents(ctx)
	doc, err := iter.Next()
	if err == iterator.Done {
		i.log.Info("Error: Iterator", "err", err.Error())
		client.Close()
		return &inv.StockItem{}, nil
	}
	if err != nil {
		i.log.Info("Error: Iterator", "err", err.Error())
		client.Close()
		return &inv.StockItem{}, nil
	}

	i.log.Info("Info", "iter", doc.Data())
	mapstructure.Decode(doc.Data(), &stock)
	i.log.Info("Info", "stock", stock.String())
	client.Close()
	return &stock, nil
}

// ChangeStock Changes the amount of stock and adds stock if not present
func (i *Inventory) ChangeStock(ctx context.Context, cr *inv.AmendRequest) (*inv.AmendResponse, error) {
	i.log.Info("Handle ChangeStock", "id", cr.GetID(), "location", cr.GetLocation(), "change amount", cr.GetAmount())

	app, err := tools.SetFirebase(ctx, i.log, i.path)
	client, err := app.Firestore(context.Background())
	if err != nil {
		i.log.Info("Error: Declaring Firestore app", "err", err.Error())
	}
	defer client.Close()

	iter := client.Collection("inventory").Where("ID", "==", cr.GetID()).Where("Location", "==", cr.GetLocation()).Documents(ctx)

	doc, err := iter.Next()
	if err == iterator.Done {
		if cr.GetAmount() < 0 {
			i.log.Info("Error: AmendResponse", "Error", "Stock Does Not Exist")
			client.Close()
			return &inv.AmendResponse{Response: "Cannot Remove from stock that does not exist."}, nil
		}
		if cr.GetLocation() == "" {
			i.log.Info("Error: AmendResponse", "Error", "Stock Has No Location")
			client.Close()
			return &inv.AmendResponse{Response: "Location Empty."}, nil
		}
		if cr.GetPrice() == 0 {
			i.log.Info("Error: AmendResponse", "Error", "Stock Has No Price")
			client.Close()
			return &inv.AmendResponse{Response: "Price cannot be 0."}, nil
		}
		value := cr.GetAmount()
		newID := tools.RandSeq()
		_, err = client.Collection("inventory").Doc(newID).Set(ctx, map[string]interface{}{
			"ID":         newID,
			"Location":   cr.GetLocation(),
			"Name":       cr.GetName(),
			"StockCount": value,
		})
		i.log.Info("Success: AmendResponse", "Stock Added")
		client.Close()

		var conn *grpc.ClientConn
		conn, err := grpc.Dial(":9092", grpc.WithInsecure())
		if err != nil {
			i.log.Error("Unable to listen", err)
		}
		defer conn.Close()
		costCli := cost.NewCostClient(conn)
		mess := cost.UpdateRequest{
			ID:    newID,
			Price: cr.GetPrice(),
		}
		res, err := costCli.AddUnitCost(ctx, &mess)
		if err != nil {
			i.log.Error("Error when calling AddUnitCost", err)
			conn.Close()
		}
		i.log.Info("Cost MicroService AddUnitCost", "Response", res.GetSuccess())

		return &inv.AmendResponse{Response: "Successfully Added."}, nil
	}
	if err != nil {
		i.log.Info("Error: AmendResponse", "Error", err)
		client.Close()
		return &inv.AmendResponse{Response: "Error Occurred During Operation."}, nil
	}
	curVal := doc.Data()["StockCount"].(int64)
	postVal := curVal + cr.GetAmount()
	if postVal >= 0 {
		_, err = client.Collection("inventory").Doc(doc.Ref.ID).Update(ctx, []firestore.Update{
			{
				Path:  "StockCount",
				Value: postVal,
			},
		})

	} else {
		client.Close()
		i.log.Info("Error: AmendResponse", "Value after operation", postVal)
		return &inv.AmendResponse{Response: "Not Enough Stock."}, nil
	}

	client.Close()
	i.log.Info("Success: AmendResponse", "Value after operation", postVal)
	return &inv.AmendResponse{Response: "Successfully Changed."}, nil
}

// CheckShort Gets a list of all stocks that are short of items
//TODO Add a field to stock item saying when to replace items ie replace toilet rolls if stock < 100 but replace toothpaste when stock < 20
func (i *Inventory) CheckShort(ctx context.Context, sr *inv.ShortRequest) (*inv.ShortList, error) {
	i.log.Info("Handle CheckShort", "location", sr.GetLocation())
	var short inv.ShortList
	app, err := tools.SetFirebase(ctx, i.log, i.path)
	client, err := app.Firestore(context.Background())
	if err != nil {
		i.log.Info("Error: Declaring Firestore app", "err", err.Error())
	}
	defer client.Close()

	iter := client.Collection("inventory").Where("StockCount", "<=", 100).Where("Location", "==", sr.GetLocation()).Documents(ctx)

	for {
		doc, err := iter.Next()
		var stock inv.StockItem
		if err == iterator.Done {
			i.log.Info("Error: Iterator Finished", "err", err.Error())
			break
		}
		if err != nil {
			client.Close()
			i.log.Info("Error: Iterator Error", "err", err.Error())
			return &inv.ShortList{}, nil
		}
		i.log.Info("Info", "stock", stock.String())
		mapstructure.Decode(doc.Data(), &stock)
		short.SList = append(short.SList, &stock)
	}
	i.log.Info("Info", "short", short.String())

	client.Close()

	return &short, nil
}

// GetStore Gets a list of all stocks for a location
func (i *Inventory) GetStore(ctx context.Context, sr *inv.ShortRequest) (*inv.ShortList, error) {
	i.log.Info("Handle GetStore", "location", sr.GetLocation())
	var store inv.ShortList
	app, err := tools.SetFirebase(ctx, i.log, i.path)
	client, err := app.Firestore(context.Background())
	if err != nil {
		i.log.Info("Error: Declaring Firestore app", "err", err.Error())
	}
	defer client.Close()

	iter := client.Collection("inventory").Where("Location", "==", sr.GetLocation()).Documents(ctx)
	for {
		doc, err := iter.Next()
		var stock inv.StockItem
		if err == iterator.Done {
			i.log.Info("Error: Iterator Finished", "err", err.Error())
			break
		}
		if err != nil {
			client.Close()
			i.log.Info("Error: Iterator Error", "err", err.Error())
			return &inv.ShortList{}, nil
		}
		i.log.Info("Info", "stock", doc.Data())
		mapstructure.Decode(doc.Data(), &stock)
		store.SList = append(store.SList, &stock)
	}

	client.Close()

	return &store, nil
}
