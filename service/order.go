package service

import (
	"context"
	"encoding/json"
	"fmt"
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
	fmt.Println(orderID)
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

	var order *Order
	err = json.Unmarshal(resJSON, &order)

	if err != nil {
		return & orderpb.GetSpecificOrderResponse{
		},err
	}

	var itempbSlice []*orderpb.Item

	for _,item := range order.Items {
		itempb := &orderpb.Item{
			ItemId:item.ItemID,
			Name:item.Name,
			Cost:item.Cost,
			Quantity:item.Quantity,
		}
		itempbSlice = append(itempbSlice,itempb)
	}


	resp := &orderpb.GetSpecificOrderResponse{
		Order:&orderpb.Order{
			CustomerId:order.CustomerID,
			OrderId:order.OrderID,
			RestaurantId:order.RestaurantID,
			Discount:order.Discount,
			Items:itempbSlice,
		},
	}
	return resp,nil
}


func (s *Service) UpdateOrder(ctx context.Context, req *orderpb.UpdateOrderRequest) (*orderpb.UpdateOrderResponse, error) {
	customerId := req.GetCustomerId()
	orderId := req.GetOrderId()
	itemId := req.GetItemId()
	quantity := req.GetQuantity()

	existingOrder, err := dynamodb.GetItem("orderDemo","CustomerID",customerId,"OrderID",orderId,&Order{})

	if err != nil {
		return & orderpb.UpdateOrderResponse{
		},err
	}

	existingOrderJSON,err := json.Marshal(existingOrder)

	if err != nil {
		return & orderpb.UpdateOrderResponse{
		},err
	}

	var order *Order
	err = json.Unmarshal(existingOrderJSON, &order)

	for i,item := range order.Items {
		if item.ItemID == itemId {
			item.Quantity = quantity
			order.Items[i] = item
			break
		}
	}

	i := 0
	for _, item := range order.Items {
		if item.Quantity > 0 {
			order.Items[i] = item
			i++
		}
	}

	order.Items = order.Items[:i]

	err = dynamodb.PutItem("orderDemo",order)

	return & orderpb.UpdateOrderResponse{
		Response:"Order Updated",
	},err
}


func (s *Service) DeleteOrder(ctx context.Context, req *orderpb.DeleteOrderRequest) (*orderpb.DeleteOrderResponse, error) {
	customerId := req.GetCustomerId()
	orderId := req.GetOrderId()

	err := dynamodb.DeleteItem("orderDemo","CustomerID",customerId,"OrderID",orderId)

	if err != nil {
		return & orderpb.DeleteOrderResponse{
		},err
	}

	return & orderpb.DeleteOrderResponse{
		Response:"Order Deleted Successfully",
	},err
}

