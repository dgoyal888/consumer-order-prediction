package service

import (
	"context"
	//"fmt"
	"github.com/consumer-order-prediction/pkg/dynamodb"
	//"github.com/consumer-order-prediction/pkg/grpc"
	customerpb "github.com/consumer-order-prediction/pkg/proto/customer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	//"net"
	"testing"
)

//const bufSize = 1024 * 1024
//var lis *bufconn.Listener
//
//func bufDialer(context.Context, string) (net.Conn, error) {
//	return lis.Dial()
//}

var testCustomer = &customerpb.AddCustomerRequest{
	Customer: &customerpb.Customer{
		CustomerId: "bbf5f42d-ddb0-4e16-91ae-fe87e348da88",
		FirstName: "kishore",
		SecondName: "Kumar",
		Address: "Mumbai",
	},
}

func init() {

	dynamodb.NewClient()
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	customerpb.RegisterCustomerServiceServer(s, &Service{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func TestGetCustomerPass(t *testing.T){
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	customerID := "9261f620-d489-11ea-8a68-c4b301d68639"
	cc := customerpb.NewCustomerServiceClient(conn)
	req := &customerpb.GetSpecificCustomerRequest{
		CustomerId: customerID,
	}
	_, err = cc.GetSpecificCustomer(context.Background(), req)
	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
}
//func TestGetCustomerFail(t *testing.T){
//	ctx := context.Background()
//	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
//	if err != nil {
//		t.Fatalf("failed to dial: %v", err)
//	}
//	defer conn.Close()
//
//	customerID := "123"
//	cc := customerpb.NewCustomerServiceClient(conn)
//	req := &customerpb.GetSpecificCustomerRequest{
//		CustomerId: customerID,
//	}
//	_, err = cc.GetSpecificCustomer(context.Background(), req)
//	if err != nil {
//		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
//	}
//}
func TestAddCustomer(t *testing.T){
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer),grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	cc := customerpb.NewCustomerServiceClient(conn)

	req := testCustomer

	_, err = cc.AddCustomer(context.Background(), req)

	if err != nil {
		t.Fatalf("Error While calling Adding Customer : %v ", err)
	}
}
func TestGetAllCustomer(t *testing.T){
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	if err != nil {
		t.Fatalf("Error While calling GetOrderDetail : %v ", err)
	}
	//cc := customerpb.NewCustomerServiceClient(conn)
	//_, err = cc.GetAllCustomer(context.Background(),)
}