package main

import (
	"net"
	"os"

	"google.golang.org/grpc/reflection"

	logic "github.com/JBurton26/msoa-go/logic"
	cost "github.com/JBurton26/msoa-go/protos/cost"
	inv "github.com/JBurton26/msoa-go/protos/inventory"
	order "github.com/JBurton26/msoa-go/protos/order"
	user "github.com/JBurton26/msoa-go/protos/user"

	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()

	invConf := "creds/msoa-go-inventory.json"
	i := logic.NewInventory(log, invConf)
	costConf := "creds/msoa-go-cost.json"
	c := logic.NewCost(log, costConf)
	//userConf := "creds/msoa-go-user.json"
	u := logic.NewUser(log)
	o := logic.NewOrder(log)

	inv.RegisterInventoryServer(gs, i)
	cost.RegisterCostServer(gs, c)
	user.RegisterUserServer(gs, u)
	order.RegisterOrderServer(gs, o)

	reflection.Register(gs)
	l, err := net.Listen("tcp", ":9092")

	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
