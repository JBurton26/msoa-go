package logic

import (
	"context"

	cost "github.com/JBurton26/msoa-go/protos/cost"
	inv "github.com/JBurton26/msoa-go/protos/inventory"
	order "github.com/JBurton26/msoa-go/protos/order"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc"
)

// Order in the room
type Order struct {
	log  hclog.Logger
	path string
}

// NewOrder Inherits from Joy Division
func NewOrder(l hclog.Logger, s string) *Order {
	return &Order{l, s}
}

// MakeOrder makes an order
func (o *Order) MakeOrder(ctx context.Context, or *order.OrderRequest) (*order.OrderResponse, error) {
	/*
		To be called when the order is complete
		Should take in the users details and the staff members details and call the total items from the costs service
		if an email is provided this function should send an email
	*/
	return &order.OrderResponse{Body: "Hello World!"}, nil
}

// AddToBasket da
func (o *Order) AddToBasket(ctx context.Context, addreq *order.AddToCart) (*order.Cart, error) {
	var basket order.Cart
	var cart cost.Basket
	if addreq.GetToAdd().GetLocation() != addreq.GetTrolley().GetLocation() {
		o.log.Info("Non Matching Location Values", "Item Location", addreq.GetToAdd().GetLocation(), "Trolley Location", addreq.GetTrolley().GetLocation())
		return addreq.GetTrolley(), nil
	}
	basket.Items = append(addreq.GetTrolley().GetItems(), addreq.GetToAdd())

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9092", grpc.WithInsecure())
	if err != nil {
		o.log.Error("Unable to listen", err)
	}
	defer conn.Close()

	invCli := inv.NewInventoryClient(conn)
	levReq := inv.LevelRequest{ID: addreq.GetToAdd().GetID(), Location: addreq.GetToAdd().GetLocation()}
	res1, err := invCli.GetStock(ctx, &levReq)
	if err != nil {
		o.log.Error("Error when calling GetStock", err)
		conn.Close()
	}
	o.log.Info("InventoryMicroService GetStock", "Response", res1.GetStockCount())

	if res1.GetStockCount() < addreq.GetToAdd().GetCount() {
		o.log.Info("Value Larger Than Stock Count", "Stock Count", res1.GetStockCount(), "Trolley Location", addreq.GetToAdd().GetCount())
		return addreq.GetTrolley(), nil
	}

	costCli := cost.NewCostClient(conn)
	mapstructure.Decode(addreq.GetTrolley(), &cart)
	res2, err := costCli.TotalBasket(ctx, &cart)
	if err != nil {
		o.log.Error("Error when calling TotalBasket", err)
		conn.Close()
	}
	o.log.Info("Cost MicroService TotalBasket", "Response", res2.GetPrice())
	basket.Price = res2.GetPrice()

	change := inv.AmendRequest{
		Amount: -addreq.GetToAdd().GetCount(),
	}
	res3, err := invCli.ChangeStock(ctx, &change)
	if err != nil {
		o.log.Error("Error when calling ChangeStock", err)
		conn.Close()
	}
	o.log.Info("Inventory MicroService ChangeStock", "Response", res3.GetResponse())
	conn.Close()
	return &basket, nil
}
