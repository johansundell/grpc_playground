package main

import (
	"log"
	"net"

	pb "github.com/johansundell/grpc_playground/geodata"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct {
}

func (s *server) GetPostalRegion(ctx context.Context, in *pb.PostalRegionRequest) (*pb.PostalRegionResponse, error) {
	return &pb.PostalRegionResponse{PostalRegion: "Not implemented yet ;)"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	pb.RegisterGeoDataServer(s, &server{})
	s.Serve(lis)
}
