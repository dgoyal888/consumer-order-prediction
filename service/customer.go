package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgoyal888/consumer-order-prediction/pkg/dynamodb"
	customerpb "github.com/dgoyal888/consumer-order-prediction/pkg/proto/customer"
	"github.com/dgoyal888/consumer-order-prediction/util"
)

type Customer struct {
	CustomerID string
	FirstName  string
	LastName   string
	Address    string
}

func (s *Service) AddCustomer(ctx context.Context, req *customerpb.AddCustomerRequest) (*customerpb.AddCustomerResponse, error) {

	cust := req.GetCustomer()
	customerID, err := util.GenerateUUID()
	fmt.Println(customerID)
	if err != nil {
		return & customerpb.AddCustomerResponse{
			Response:"Error occurred while Adding customer",
		},err
	}
	var customerStruct = &Customer{
		CustomerID: customerID,
		FirstName: cust.GetFirstName(),
		LastName: cust.GetSecondName(),
		Address: cust.GetAddress(),
	}

	err = dynamodb.PutItem("customers",customerStruct)
	if err != nil {
		return &customerpb.AddCustomerResponse{
			Response:"Error occurred while adding customer",
		},err
	}
	res := &customerpb.AddCustomerResponse{
		Response:"Customer Added successfully",
	}
	return res,nil
}
func (s *Service) GetSpecificCustomer(ctx context.Context, req *customerpb.GetSpecificCustomerRequest) (*customerpb.GetSpecificCustomerResponse, error) {

	customerID := req.CustomerId

	res, err := dynamodb.GetItem("customers","CustomerID",customerID,"CustomerID",customerID,&Customer{})

	if err != nil {
		return nil,err
	}

	resJSON,err := json.Marshal(res)

	if err != nil {
		return & customerpb.GetSpecificCustomerResponse{
		},err
	}

	var customer *Customer
	err = json.Unmarshal(resJSON, &customer)

	if err != nil {
		return & customerpb.GetSpecificCustomerResponse{
		},err
	}


	resp := &customerpb.GetSpecificCustomerResponse{
		Customer:&customerpb.Customer{
			CustomerId:customer.CustomerID,
			FirstName: customer.FirstName,
			SecondName: customer.LastName,
			Address: customer.Address,
		},
	}
	return resp,nil
}

func (s *Service) GetAllCustomer(ctx context.Context, req *customerpb.GetAllCustomerRequest) (*customerpb.GetAllCustomerResponse, error) {

	res, err := dynamodb.GetAllItem("customers")
	if err != nil {
		return &customerpb.GetAllCustomerResponse{
		},err
	}

	return & customerpb.GetAllCustomerResponse{
		Count:res,
	},nil
}

