package main

import (
    "github.com/consumer-order-prediction/pkg/dynamodb"
    customerpb "github.com/consumer-order-prediction/pkg/proto/customer"
    orderspb "github.com/consumer-order-prediction/pkg/proto/orders"
    "github.com/consumer-order-prediction/service"
    "google.golang.org/grpc"
    "log"
    "net"
)


func main() {

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dynamodb.NewClient()

	s := grpc.NewServer()

	customerpb.RegisterCustomerServiceServer(s, &service.Service{})
	orderspb.RegisterOrderServiceServer(s, &service.Service{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}