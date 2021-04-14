package main

import (
	"net"
	"os"

	"google.golang.org/grpc/reflection"

	logic "github.com/JBurton26/msoa-go/logic"
	cost "github.com/JBurton26/msoa-go/protos/cost"
	inv "github.com/JBurton26/msoa-go/protos/inventory"
	user "github.com/JBurton26/msoa-go/protos/user"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	i := logic.NewInventory(log)
	c := logic.NewCost(log)
	u := logic.NewUser(log)
	inv.RegisterInventoryServer(gs, i)
	cost.RegisterCostServer(gs, c)
	user.RegisterUserServer(gs, u)
	reflection.Register(gs)
	l, err := net.Listen("tcp", ":9092")

	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
