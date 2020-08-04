package main

import (
	"github.com/consumer-order-prediction/pkg/dynamodb"
	restaurantpb "github.com/consumer-order-prediction/pkg/proto/restaurant"
	"github.com/consumer-order-prediction/service"
	"google.golang.org/grpc"
	"log"
	"net"
)


func main() {

	lis, err := net.Listen("tcp", ":50052")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dynamodb.NewClient()

	s := grpc.NewServer()


	restaurantpb.RegisterRestaurantServiceServer(s, &service.Service{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}