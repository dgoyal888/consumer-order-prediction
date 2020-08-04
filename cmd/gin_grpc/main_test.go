package main

import (
	"bytes"
	"encoding/json"
	orderpb "github.com/consumer-order-prediction/pkg/proto/orders"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

var Token = "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTk2NTQ4MzQyLCJ1c2VybmFtZSI6ImFkbWluIn0.LSwKgvGEEWSfpuxVZ1rw8SrRC0lwEprDRA0axGtKXFcuSXnVpIjsqrz8YKj6lsCDzDS5LnhdPpXmKpfDeQQPmQ"


func TestPlaceOrder(t *testing.T) {
	router := gin.Default()
	api := router.Group("/api")
	api.POST("/order",PlaceOrder)

	w := httptest.NewRecorder()

	jsonOrder, _ := json.Marshal(testOrder)

	req, err := http.NewRequest("POST","/api/order",bytes.NewBuffer(jsonOrder))
	req.Header.Add("Authorization",Token)

	response := httptest.NewRecorder()
	if err != nil{
		t.Fatalf("Error While calling Create Order : %v ", err)
	}
	assert.Equal(t, 200, response.Code, "Order placed successful")

	router.ServeHTTP(w, req)
}

func TestUpdateOrder(t *testing.T)  {
	router := gin.Default()
	api := router.Group("/api")
	api.PUT("/order",UpdateOrder)

	w := httptest.NewRecorder()

	updateOrder := &orderpb.UpdateOrderRequest{
		CustomerId:"bc011d3b-7337-4abe-9e56-8005e64403ee",
		OrderId:"4541a5f8-d640-11ea-a41b-c4b301c9617d",
		ItemId:"28200a58-363e-45c7-b17b-4a9de2b6abae",
		Quantity: 25,
	}

	jsonOrder, _ := json.Marshal(updateOrder)

	req, err := http.NewRequest("PUT","/api/order",bytes.NewBuffer(jsonOrder))
	req.Header.Add("Authorization",Token)


	response := httptest.NewRecorder()
	if err != nil{
		t.Fatalf("Error While calling updating Order : %v ", err)
	}
	assert.Equal(t, 200, response.Code, "Order updated successful")

	router.ServeHTTP(w, req)
}

func TestGetSpecificOrder(t *testing.T)  {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/customer/:customerid/order/:orderid",GetSpecificOrder)

	w := httptest.NewRecorder()

	customerId := "bc011d3b-7337-4abe-9e56-8005e64403ee"
	orderId := "a45347c9-a965-4bc3-8131-bf7705885031"
	url := "/api/customer/" + customerId + "/order/" + orderId

	req, err := http.NewRequest("GET",url,nil)
	req.Header.Add("Authorization",Token)

	response := httptest.NewRecorder()
	if err != nil{
		t.Fatalf("Error While getting specific Order : %v ", err)
	}
	assert.Equal(t, 200, response.Code, "Order fetched successful")

	router.ServeHTTP(w, req)
}

//func TestDeleteOrder(t *testing.T)  {
//	router := gin.Default()
//	api := router.Group("/api")
//	api.DELETE("/deleteOrder/:id", DeleteOrder)
//
//	w := httptest.NewRecorder()
//
//	customerId := "bc011d3b-7337-4abe-9e56-8005e64403ee"
//	orderId := "a45347c9-a965-4bc3-8131-bf7705885031"
//	url := "/api/deleteOrder/" + orderId
//
//	req, err := http.NewRequest("DELETE",url,nil)
//	req.Header.Add("Authorization",Token)
//
//	response := httptest.NewRecorder()
//	if err != nil{
//		t.Fatalf("Error While deleting specific Order : %v ", err)
//	}
//	assert.Equal(t, 200, response.Code, "Order deleted successful")
//
//	router.ServeHTTP(w, req)
//}

