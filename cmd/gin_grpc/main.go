package main

import (
	"context"
	orderspb "github.com/consumer-order-prediction/pkg/proto/orders"
	customerpb "github.com/consumer-order-prediction/pkg/proto/customer"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World from GIN to Swiggy",
	})
}

//Return most popular restaurant with the help of grpc
/*func  GetPoplarRestaurant(c *gin.Context) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	defer conn.Close()

	client := orderspb.NewOrderServiceClient(conn)

	req := &orderspb.GetPopularRestaurantRequest{
	}

	res, err := client.GetPopularRestaurant(context.Background(), req)

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	c.JSON(200, gin.H{
		"Most Popular Restaurant": res.Name,
	})
}

func  GetPopularVegCuisine(c *gin.Context) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	defer conn.Close()

	client := orderspb.NewOrderServiceClient(conn)

	req := &orderspb.GetPopularVegCuisineRequest{
	}

	res, err := client.GetPopularVegCuisine(context.Background(), req)

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	c.JSON(200, gin.H{
		"Most Popular Veg Cuisine is": res.Name,
	})
}


//Return a specific order on the basis of customer id with the help of gRPC
func GetSpecificOrdersByQuery(c *gin.Context) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	defer conn.Close()

	client := orderspb.NewOrderServiceClient(conn)

	customerid := c.Query("CustomerID")

	req := &orderspb.GetSpecificOrderRequest{
		OrderId:customerid,
	}

	res, err := client.GetSpecificOrder(context.Background(), req)

	if err != nil {
		log.Fatalf("Error While calling GreetFullName : %v", err)
	}

	if err!=nil{
		c.JSON(200, gin.H{
			"message":"customer not found",
		})
	}else {
		ginRes := &csv.Order{
			CustomerID:res.GetOrder().CustomerId,
			CustomerName:res.GetOrder().CustomerName,
			RestsurantName:res.GetOrder().RestsurantName,
			VegCuisine:res.GetOrder().VegCuisine,
			NonVegCuisine:res.GetOrder().NonvegCuisine,
			State:res.GetOrder().State,
		}
		c.JSON(200,ginRes)
	}
}*/

func PlaceOrder (c *gin.Context) {
	var req orderspb.PlaceOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(req)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close();

	oc := orderspb.NewOrderServiceClient(conn)

	res, err := oc.PlaceOrder(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error While calling CreateOrder : %v ", err)
		c.JSON(500, gin.H{
		})
		return
	}

	c.JSON(200, gin.H{
		"Message":res.Response,
	})
}

func UpdateOrder (c *gin.Context) {
	var req orderspb.UpdateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(req)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := orderspb.NewOrderServiceClient(conn)

	res, err := oc.UpdateOrder(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error While calling UpdateDish : %v ", err)
		c.JSON(500, gin.H{
			"Message": res.Response,
		})
		return
	}

	c.JSON(200, gin.H{
		"Message":res.Response,
	})
}

func DeleteOrder (c *gin.Context) {
	id := c.Param("id")

	req := &orderspb.DeleteOrderRequest{
		OrderId:id,
	}


	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := orderspb.NewOrderServiceClient(conn)

	res, err := oc.DeleteOrder(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling UpdateDish : %v ", err)
		c.JSON(500, gin.H{
			"Message": res.Response,
		})
		return
	}

	c.JSON(200, gin.H{
		"Message":res.Response,
	})
}

func GetSpecificOrder (c *gin.Context) {
	customerId := c.Param("customerid")
	orderId := c.Param("orderid")

	req := &orderspb.GetSpecificOrderRequest{
		CustomerId:customerId,
		OrderId:orderId,
	}


	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := orderspb.NewOrderServiceClient(conn)

	res, err := oc.GetSpecificOrder(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling UpdateDish : %v ", err)
		c.JSON(500, gin.H{
			"Message": "Somne erro ocurred",
		})
		return
	}

	c.JSON(200, gin.H{
		"Order":res.Order,
	})
}

func AddCustomer (c *gin.Context) {
	var req customerpb.AddCustomerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(req)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close();

	oc := customerpb.NewCustomerServiceClient(conn)

	res, err := oc.AddCustomer(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error While calling CreateOrder : %v ", err)
		c.JSON(500, gin.H{
		})
		return
	}

	c.JSON(200, gin.H{
		"Message":res.Response,
	})
}

func GetSpecificCustomer (c *gin.Context) {
	customerId := c.Param("customerid")

	req := &customerpb.GetSpecificCustomerRequest{
		CustomerId:customerId,
	}


	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := customerpb.NewCustomerServiceClient(conn)

	res, err := oc.GetSpecificCustomer(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling UpdateDish : %v ", err)
		c.JSON(500, gin.H{
			"Message": "Somne erro ocurred",
		})
		return
	}

	c.JSON(200, gin.H{
		"Customer":res.Customer,
	})
}

func GetCustomerCount (c *gin.Context) {

	req := &customerpb.GetAllCustomerRequest{
	}


	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := customerpb.NewCustomerServiceClient(conn)

	res, err := oc.GetAllCustomer(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling UpdateDish : %v ", err)
		c.JSON(500, gin.H{
			"Message": "Somne erro ocurred",
		})
		return
	}

	c.JSON(200, gin.H{
		"Total Customers":res.Count,
	})
}

func main(){
	router := gin.Default()

	api:= router.Group("/api",gin.BasicAuth(gin.Accounts{
		"team1": "team1",
	}))


	// http://localhost:5656/api/
	api.GET("/",  HomePage)

	//Order API's
	api.POST("/order", PlaceOrder)
	api.PUT("/order", UpdateOrder)
	api.DELETE("/deleteOrder/:id", DeleteOrder)
	api.GET("/customer/:customerid/order/:orderid",GetSpecificOrder)


	//Customer API's
	api.GET("/customer/:customerid",GetSpecificCustomer)
	api.GET("/customers",GetCustomerCount)
	api.POST("/customer", AddCustomer)


	//api.GET("/popularrestaurant", GetPoplarRestaurant)
	//api.GET("/popularcuisine", GetPopularVegCuisine)
	//api.GET("/orders", GetSpecificOrdersByQuery)

	router.Run("localhost:5656")
}