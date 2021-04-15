package main

import (
	"context"
	"os"

	"log"

	cost "github.com/JBurton26/msoa-go/protos/cost"
	inv "github.com/JBurton26/msoa-go/protos/inventory"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9092", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Unable to listen", err)
		os.Exit(1)
	}
	defer conn.Close()

	invCli := inv.NewInventoryClient(conn)
	mess := inv.LevelRequest{
		ID:       "hello",
		Location: "world",
	}
	response, err := invCli.GetStock(context.Background(), &mess)
	if err != nil {
		log.Fatal("Error when calling GetStock", err)
		os.Exit(1)
	}
	log.Print("Response from Server: ", response.StockCount)

	costCli := cost.NewCostClient(conn)
	mess2 := cost.CostRequest{
		ID: "hello",
	}
	response2, err := costCli.GetUnitCost(context.Background(), &mess2)
	if err != nil {
		log.Fatal("Error when calling GetUnitCost", err)
		os.Exit(1)
	}
	log.Print("Response from Server: ", response2.Cost)

	conn.Close()
}
