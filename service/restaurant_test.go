package service

import (
	"github.com/consumer-order-prediction/pkg/dynamodb"
	restaurantpb "github.com/consumer-order-prediction/pkg/proto/restaurant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
	"context"
)

const bufSize = 1024 * 1024
var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func init() {
	dynamodb.NewClient()
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	restaurantpb.RegisterRestaurantServiceServer(s,&Service{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

var testItems = []*restaurantpb.Item{
	{
		ItemId: "a283fa27-82fa-480a-9e61-206e04d64766",
		Name:   "panipuri",
		Price:  100,
	},
}
var testRestaurantItem = &restaurantpb.AddRestaurantItemRequest{
	RestaurantId: "7243e142-871c-4bde-82c0-35cef99b2d55",
	ItemName: "bhel",
	Price: 60,

}
var testRestaurant = &restaurantpb.AddRestaurantRequest{
	Restaurant: &restaurantpb.Restaurant{
	RestaurantId: "c6ebb3a1-4105-4317-9cf5-d93b319d5e3a",
	Name: "Mayur Sweets",
	Address: "Gandhi Chowk",
	Items: testItems,
	},
}
var TestUpdateResItem = &restaurantpb.UpdateRestaurantItemRequest{
	RestaurantId: "7243e142-871c-4bde-82c0-35cef99b2d55",
	ItemId: "6ec600c3-bb4e-49b0-95aa-eb84578873ef",
	ItemName: "c",
	Price: 100,
}
var TestDeleteIt = &restaurantpb.DeleteItemRequest{
	RestaurantId: "7243e142-871c-4bde-82c0-35cef99b2d55",
	ItemId: "6ec600c3-bb4e-49b0-95aa-eb84578873ef",
}
var TestSpecificIt = &restaurantpb.GetSpecificItemRequest{
	RestaurantId: "7243e142-871c-4bde-82c0-35cef99b2d55",
	ItemId: "6ec600c3-bb4e-49b0-95aa-eb84578873ef",
}

var TestGetAllItemsRequest = &restaurantpb.GetAllItemsRequest{
	RestaurantId: "7243e142-871c-4bde-82c0-35cef99b2d55",
}
func TestAddRestaurant(t *testing.T) {
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer),grpc.WithInsecure())
		if err != nil {
			t.Fatalf("failed to dial: %v", err)
		}
		defer conn.Close()
		cc := restaurantpb.NewRestaurantServiceClient(conn)

		req := testRestaurant

		_, err = cc.AddRestaurant(context.Background(), req)

		if err != nil {
			t.Fatalf("Error While calling Adding Restaurant : %v ", err)
		}
}

func TestGetAllRestaurant(t *testing.T) {
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
		if err != nil {
			t.Fatalf("failed to dial: %v", err)
		}
		defer conn.Close()

		if err != nil {
			t.Fatalf("Error While calling GetOrderDetail : %v ", err)
		}
		rc := restaurantpb.NewRestaurantServiceClient(conn)
		req := &restaurantpb.GetAllRestaurantRequest{}
		_, err = rc.GetAllRestaurant(context.Background(),req)
		if err != nil {
			t.Fatalf("Error While calling GetAllRestaurants : %v ", err)
		}
}

func TestGetSpecificRestaurant(t *testing.T) {
		ctx := context.Background()
		conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
		if err != nil {
			t.Fatalf("failed to dial: %v", err)
		}
		defer conn.Close()
		restaurantID := "7243e142-871c-4bde-82c0-35cef99b2d55"
	rc := restaurantpb.NewRestaurantServiceClient(conn)
	req := &restaurantpb.GetSpecificRestaurantRequest{
		RestaurantId: restaurantID,
	}
		_, err = rc.GetSpecificRestaurant(context.Background(), req)
		if err != nil {
			t.Fatalf("Error While calling GetOrderDetail : %v ", err)
		}
}

func TestAddRestaurantItem(t *testing.T){
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer),grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cc := restaurantpb.NewRestaurantServiceClient(conn)

	req := testRestaurantItem

	_, err = cc.AddRestaurantItem(context.Background(), req)

	if err != nil {
		t.Fatalf("Error While calling Adding Restaurant : %v ", err)
	}
}

func TestUpdateRestaurantItem(t *testing.T){
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer),grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cc := restaurantpb.NewRestaurantServiceClient(conn)
	req := TestUpdateResItem
	_, err = cc.UpdateRestaurantItem(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling Create Order : %v ", err)
	}
}

func TestDeleteItem(t *testing.T){
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer),grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cc := restaurantpb.NewRestaurantServiceClient(conn)
	req := TestDeleteIt
	_, err = cc.DeleteItem(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling Create Order : %v ", err)
	}
}

func TestGetAllItems(t *testing.T){
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer),grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cc := restaurantpb.NewRestaurantServiceClient(conn)
	req := TestGetAllItemsRequest
	_, err = cc.GetAllItems(context.Background(), req)

		if err != nil {
			t.Fatalf("Error While calling  GetAllItems : %v ", err)
		}
}

func TestGetSpecificItem(t *testing.T){

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer),grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cc := restaurantpb.NewRestaurantServiceClient(conn)
	req := &restaurantpb.GetSpecificItemRequest{
		RestaurantId: "9814df3f-b6b2-49c2-aa0f-7dc25bbde4f5",
		ItemId: "1460ec5c-1a1a-43ec-b232-8994ad898303",
	}
	_, err = cc.GetSpecificItem(context.Background(), req)
	if err!=nil{
		t.Fatalf("Error while calling GetSpecificItems: %v ",err)
	}
}
