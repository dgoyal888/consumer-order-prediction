package main

import (
	"consumer-order-prediction/pkg/rules"
	"context"
	"fmt"
	"log"
	"net"

	orderspb "consumer-order-prediction/pkg/proto/orders"
	"google.golang.org/grpc"
)


type server struct{}


func (s *server) GetPopularRestaurant(ctx context.Context, req *orderspb.GetPopularRestaurantRequest) (*orderspb.GetPopularRestaurantResponse, error) {

	/*csvData, err := csv.ReadCSVData("data/orderdata.csv")

	if err != nil {
		fmt.Println("Error while reading CSV data: %s",err.Error())
	}

	err = csv.CSVToJSON(csvData)

	if err != nil {
		fmt.Println("Error while converting  CSV to Json data: %s",err.Error())
	}*/

	restaurant,err := rules.PopularRestaurant("data/orderdata.json")

	if err != nil {
		fmt.Println("error from rules %v",err)
	}

	fmt.Println("hey",restaurant.RestsurantName)

	res := &orderspb.GetPopularRestaurantResponse{
		Name: restaurant.RestsurantName,
	}


	return res,nil
}

func (s *server) GetSpecificOrder(ctx context.Context, req *orderspb.GetSpecificOrderRequest) (*orderspb.GetSpecificOrderResponse, error) {

	order, err := rules.ReturnJsonBasedOnCUSTID(req.OrderId)

	if err != nil {
		fmt.Println("error from rules %v",err)
	}

	res := &orderspb.GetSpecificOrderResponse{
		Order: &orderspb.Order{
			CustomerId: order.CustomerID,
			CustomerName: order.CustomerName,
			RestsurantName:order.RestsurantName,
			VegCuisine:order.VegCuisine,
			NonvegCuisine:order.NonVegCuisine,
			State:order.State,
		},
	}

	return res,nil
}


func main() {

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	orderspb.RegisterOrderServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}