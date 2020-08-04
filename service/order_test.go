package service

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
	"google.golang.org/grpc/test/bufconn"

	orderpb "github.com/cdgoyal888/onsumer-order-prediction/pkg/proto/orders"
	"github.com/dgoyal888/consumer-order-prediction/pkg/dynamodb"
)

const bufSize = 1024 * 1024
var lis *bufconn.Listener


var orderItem = []*orderpb.Item{
	{
		ItemId:"28200a58-363e-45c7-b17b-4a9de2b6abae",
		Name:"Hello",
		Cost:1.1,
		Quantity:20,
	},
}

var testOrder = &orderpb.PlaceOrderRequest{
	Order:&orderpb.Order{
		CustomerId:"bc011d3b-7337-4abe-9e56-8005e64403ee",
		OrderId:"a45347c9-a965-4bc3-8131-bf7705885031",
		RestaurantId:"0db4dde4-1c3f-45db-9159-2c02736566d9",
		Discount:1.1,
		Items:orderItem,
	},
}

var updateOrder = &orderpb.UpdateOrderRequest{
	CustomerId:"bc011d3b-7337-4abe-9e56-8005e64403ee",
	OrderId:"28200a58-363e-45c7-b17b-4a9de2b6abae",
	ItemId:"28200a58-363e-45c7-b17b-4a9de2b6abae",
	Quantity: 25,
}

func init() {
	lis = bufconn.Listen(bufSize)
	dynamodb.NewClient()
	s := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(s, &Service{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestPlaceOrder(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer),grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderpb.NewOrderServiceClient(conn)

	req := testOrder

	_, err = oc.PlaceOrder(context.Background(), req)

	if err != nil {
		t.Fatalf("Error While calling Create Order : %v ", err)
	}
}

/*func TestUpdateOrder(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer),grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderpb.NewOrderServiceClient(conn)

	req := updateOrder

	_, err = oc.UpdateOrder(context.Background(), req)

	if err != nil {
		t.Fatalf("Error While calling Create Order : %v ", err)
	}
}*/

func TestGetSpecificOrder(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer),grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderpb.NewOrderServiceClient(conn)

	req := &orderpb.GetSpecificOrderRequest{
		CustomerId:"bc011d3b-7337-4abe-9e56-8005e64403ee",
		OrderId:"a45347c9-a965-4bc3-8131-bf7705885031",
	}

	_, err = oc.GetSpecificOrder(context.Background(), req)

	if err != nil {
		t.Fatalf("Error While calling Create Order : %v ", err)
	}
}

func TestDeleteOrder(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer),grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	oc := orderpb.NewOrderServiceClient(conn)

	req := &orderpb.DeleteOrderRequest{
		CustomerId:"bc011d3b-7337-4abe-9e56-8005e64403ee",
		OrderId:"a45347c9-a965-4bc3-8131-bf7705885031",
	}

	_, err = oc.DeleteOrder(context.Background(), req)

	if err != nil {
		t.Fatalf("Error While calling Create Order : %v ", err)
	}
}