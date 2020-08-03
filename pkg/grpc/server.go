package main

import (
	"github.com/consumer-order-prediction/pkg/dynamodb"
	customerpb "github.com/consumer-order-prediction/pkg/proto/customer"
	orderspb "github.com/consumer-order-prediction/pkg/proto/orders"
	restaurantpb "github.com/consumer-order-prediction/pkg/proto/restaurant"
	"github.com/consumer-order-prediction/service"
	"google.golang.org/grpc"
	"log"
	"net"
)


//type server struct{}


// Function implementation for rpc GetPopularRestaurant
/*func (s *server) GetPopularRestaurant(ctx context.Context, req *orderspb.GetPopularRestaurantRequest) (*orderspb.GetPopularRestaurantResponse, error) {

	restaurant,err := rules.PopularRestaurant("../../data/orderdata.json")

	if err != nil {
		fmt.Println("error from rules %v",err)
	}

	res := &orderspb.GetPopularRestaurantResponse{
		Name: restaurant.RestsurantName,
	}

	return res,nil
}

func (s *server) GetPopularVegCuisine(ctx context.Context, req *orderspb.GetPopularVegCuisineRequest) (*orderspb.GetPopularVegCuisineResponse, error) {

	cuisine,err := rules.PopularVegCuisine("../../data/orderdata.json")

	if err != nil {
		fmt.Println("error from rules %v",err)
	}

	res := &orderspb.GetPopularVegCuisineResponse{
		Name: cuisine.VegCuisine,
	}

	return res,nil
}


//Function implementation for rpc GetSpecificOrder
func (s *server) GetSpecificOrder(ctx context.Context, req *orderspb.GetSpecificOrderRequest) (*orderspb.GetSpecificOrderResponse, error) {

	order, err := rules.ReturnJsonBasedOnCUSTID(req.OrderId)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
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
*/

func main() {

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dynamodb.NewClient()

	s := grpc.NewServer()

	customerpb.RegisterCustomerServiceServer(s, &service.Service{})
	orderspb.RegisterOrderServiceServer(s, &service.Service{})
	restaurantpb.RegisterRestaurantServiceServer(s, &service.Service{})

	//wg:= new(sync.WaitGroup)
	//wg.Add(1)
	//
	//go func() {
	//	http.Handle("/metrics", promhttp.Handler())
	//	panic(http.ListenAndServe(":6565", nil))
	//	wg.Done()
	//}()


	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}