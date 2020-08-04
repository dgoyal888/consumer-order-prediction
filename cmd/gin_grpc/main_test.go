package main

import (
	"bytes"
	"encoding/json"
	"github.com/consumer-order-prediction/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
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

func TestPlaceOrder(t *testing.T) {
	router := gin.Default()
	w := httptest.NewRecorder()

	item := []*service.Item{
		{
			ItemID: "28200a58-363e-45c7-b17b-4a9de2b6abae",
			Name: "hello",
			Cost: 2,
			Quantity: 10,
		},
	}

	order := & service.Order{
		CustomerID: "",
		OrderID: "a45347c9-a965-4bc3-8131-bf7705885031",
		RestaurantID: "0db4dde4-1c3f-45db-9159-2c02736566d9",
		Discount: 1.1,
		Items: item,
	}
	jsonOrder, _ := json.Marshal(order)

	req, err := http.NewRequest("POST","/api/order",bytes.NewBuffer(jsonOrder))

	response := httptest.NewRecorder()
	if err != nil{
		t.Fatalf("Error While calling Create Order : %v ", err)
	}
	assert.Equal(t, 200, response.Code, "Order placed successful")

	router.ServeHTTP(w, req)
}