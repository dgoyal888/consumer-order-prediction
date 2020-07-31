package service

import (
	"context"
	"encoding/json"

	"github.com/consumer-order-prediction/pkg/dynamodb"
	orderpb "github.com/consumer-order-prediction/pkg/proto/orders"
	"github.com/consumer-order-prediction/util"
)

type Item struct {
	ItemID   string
	Name     string
	Cost     float32
	Quantity int64
}

type Order struct {
		CustomerID   string
		OrderID      string
		RestaurantID string
		Items        []*Item
		Discount float32
}


func (s *Service) PlaceOrder(ctx context.Context, req *orderpb.PlaceOrderRequest) (*orderpb.PlaceOrderResponse, error) {

	order := req.GetOrder()

	var itemSlice []*Item
	var item *Item
	for _, v := range order.Items {
		item = &Item{
			ItemID:v.ItemId,
			Name:v.GetName(),
			Cost:v.GetCost(),
			Quantity:v.GetQuantity(),
		}
		itemSlice = append(itemSlice,item)
	}

	orderID, err := util.GenerateUUID()
	//fmt.Println("order id is ",orderID)
	if err != nil {
		return & orderpb.PlaceOrderResponse{
			Response:"Error occurred while placing order",
		},err
	}

	var orderStruct = &Order{
		CustomerID:order.GetCustomerId(),
		OrderID:orderID,
		RestaurantID:order.GetRestaurantId(),
		Items:itemSlice,
		Discount:order.GetDiscount(),
	}

	err = dynamodb.PutItem("orderDemo",orderStruct)

	if err != nil {
		return & orderpb.PlaceOrderResponse{
			Response:"Error occurred while placing order",
		},err
	}

	res := &orderpb.PlaceOrderResponse{
		Response:"Order placed successfully",
	}

	return res,nil
}

func (s *Service) GetSpecificOrder(ctx context.Context, req *orderpb.GetSpecificOrderRequest) (*orderpb.GetSpecificOrderResponse, error) {

	customerID := req.CustomerId
	orderID := req.OrderId

	res, err := dynamodb.GetItem("orderDemo","CustomerID",customerID,"OrderID",orderID,&Order{})

	if err != nil {
		return nil,err
	}

	resJSON,err := json.Marshal(res)

	if err != nil {
		return & orderpb.GetSpecificOrderResponse{
		},err
	}

	var order *orderpb.Order
	err = json.Unmarshal(resJSON, &order)

	if err != nil {
		return & orderpb.GetSpecificOrderResponse{
		},err
	}

	resp := &orderpb.GetSpecificOrderResponse{
		Order:order,
	}
	return resp,nil
}


func (s *Service) DeleteOrder(ctx context.Context, req *orderpb.DeleteOrderRequest) (*orderpb.DeleteOrderResponse, error) {


	return nil,nil
}

func (s *Service) UpdateOrder(ctx context.Context, req *orderpb.UpdateOrderRequest) (*orderpb.UpdateOrderResponse, error) {


	return nil,nil
}