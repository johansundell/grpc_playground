package main

import (
	"context"
	"log"

	"github.com/johansundell/grpc_playground/geodata"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := geodata.NewGeoDataClient(conn)
	resp, err := client.GetPostalRegion(context.Background(), &geodata.PostalRegionRequest{ZipCode: 16834})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Responce: %v", resp)

}
