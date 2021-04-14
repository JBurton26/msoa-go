package main

import (
	"context"
	"os"

	"log"

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

	c := inv.NewInventoryClient(conn)
	mess := inv.LevelRequest{
		ID:       "hello",
		Location: "world",
	}
	response, err := c.GetStock(context.Background(), &mess)
	if err != nil {
		log.Fatal("Error when calling GetStock", err)
		os.Exit(1)
	}
	log.Print("Response from Server: ", response.StockCount)
	/*
		mess2 := inv.CostRequest{
			ID: "hello",
		}
		response2, err := c.GetUnitCost(context.Background(), &mess2)
		if err != nil {
			log.Fatalf("Error when calling GetUnitCost", err)
			os.Exit(1)
		}
		log.Print("Response from Server: ", response2.StockCost)
	*/
	conn.Close()
}
