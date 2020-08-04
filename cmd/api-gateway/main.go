package main

import (
	"context"
	"fmt"
	"github.com/consumer-order-prediction/pkg/auth"
	customerpb "github.com/consumer-order-prediction/pkg/proto/customer"
	orderspb "github.com/consumer-order-prediction/pkg/proto/orders"
	restaurantpb "github.com/consumer-order-prediction/pkg/proto/restaurant"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"log"
	"net/http"
)


// variables for prometheus
var (
	PlaceOrderCnt = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "place_order_cnt",
			Help: "no of times Placeorder was hit",
		})

	GetSpecificOrderCnt = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "get_specific_order_cnt",
			Help: "no of times GetSpecificOrder was hit",
		})

	UpdateOrderCnt = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "update_order_cnt",
			Help: "no of times UpdateOrder was hit",
		})

	DeleteOrderCnt = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "delete_order_cnt",
			Help: "no of times DeleteOrder was hit",
		})
)

//variables for jwt
var (
	adminUsername = "admin"
	adminPassword = "admin"
	jwtSecret     = "iu4fcn0qnua"
)


func init()  {
	prometheus.MustRegister(PlaceOrderCnt)
	prometheus.MustRegister(GetSpecificOrderCnt)
	prometheus.MustRegister(UpdateOrderCnt)
	prometheus.MustRegister(DeleteOrderCnt)
}

type LoginRequest struct {
	Username string
	Password string
}
type AuthResponse struct {
	Token string
}
type ErrorResponse struct {
	Error string
}


func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World from GIN to Swiggy",
	})
}


// for generating token
func Login(c *gin.Context) {
	var loginReq LoginRequest
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid payload"})
		return
	}

	if loginReq.Username == adminUsername && loginReq.Password == adminPassword {
		token, err := auth.CreateToken(loginReq.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, &AuthResponse{Token: token})
		return
	}

	c.JSON(http.StatusUnauthorized, ErrorResponse{Error: "Invalid credentials"})
}

func PlaceOrder (c *gin.Context) {

	PlaceOrderCnt.Inc()

	// authenticating user
	_, err := auth.AuthenticateUser(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &ErrorResponse{Error: err.Error()})
		return
	}

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

	UpdateOrderCnt.Inc()

	// authenticating user
	_, err := auth.AuthenticateUser(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &ErrorResponse{Error: err.Error()})
		return
	}

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

	DeleteOrderCnt.Inc()

	// authenticating user
	_, err := auth.AuthenticateUser(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &ErrorResponse{Error: err.Error()})
		return
	}

	id := c.Param("id")

	req := &orderspb.DeleteOrderRequest{
		CustomerId:c.Param("customerid"),
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

	GetSpecificOrderCnt.Inc()

	// authenticating user
	_, err := auth.AuthenticateUser(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, &ErrorResponse{Error: err.Error()})
		return
	}

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

func AddRestaurant (c *gin.Context) {
	fmt.Println("request recieved")
	var req restaurantpb.AddRestaurantRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(req)

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close();

	oc := restaurantpb.NewRestaurantServiceClient(conn)

	res, err := oc.AddRestaurant(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error While calling Add Restaurant : %v ", err)
		c.JSON(500, gin.H{
		})
		return
	}

	c.JSON(200, gin.H{
		"Message":res.Response,
	})
}

func GetAllRestaurants (c *gin.Context) {

	req := &restaurantpb.GetAllRestaurantRequest{}


	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := restaurantpb.NewRestaurantServiceClient(conn)

	res, err := oc.GetAllRestaurant(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling UpdateDish : %v ", err)
		c.JSON(500, gin.H{
			"Message": "Somne erro ocurred",
		})
		return
	}

	c.JSON(200, gin.H{
		"Total Restaurants on board":res.Count,
	})
}

func GetSpecificRestaurant (c *gin.Context) {
	restaurantId := c.Param("restaurantid")

	req := &restaurantpb.GetSpecificRestaurantRequest{
		RestaurantId:restaurantId,
	}

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := restaurantpb.NewRestaurantServiceClient(conn)

	res, err := oc.GetSpecificRestaurant(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling UpdateDish : %v ", err)
		c.JSON(500, gin.H{
			"Message": "Somne erro ocurred",
		})
		return
	}

	c.JSON(200, gin.H{
		"Restaurant":res.Restaurant,
	})
}

func AddSpecificItem (c *gin.Context) {
	restaurantId := c.Param("restaurantid")

	var req *restaurantpb.AddRestaurantItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.RestaurantId = restaurantId

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := restaurantpb.NewRestaurantServiceClient(conn)

	res, err := oc.AddRestaurantItem(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling UpdateDish : %v ", err)
		c.JSON(500, gin.H{
			"Message": "Somne erro ocurred",
		})
		return
	}

	c.JSON(200, gin.H{
		"Response":res.Response,
	})
}

func UpdateRestaurantItem (c *gin.Context) {
	restaurantId := c.Param("restaurantid")

	var req *restaurantpb.UpdateRestaurantItemRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.RestaurantId = restaurantId

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := restaurantpb.NewRestaurantServiceClient(conn)

	res, err := oc.UpdateRestaurantItem(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling UpdateDish : %v ", err)
		c.JSON(500, gin.H{
			"Message": "Somne erro ocurred",
		})
		return
	}

	c.JSON(200, gin.H{
		"Response":res.Response,
	})
}

func DeleteRestaurant (c *gin.Context) {
	restaurantId := c.Param("restaurantid")

	req := &restaurantpb.DeleteItemRequest{
		RestaurantId:restaurantId,
		ItemId:c.Param("itemid"),
	}

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := restaurantpb.NewRestaurantServiceClient(conn)

	res, err := oc.DeleteItem(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling UpdateDish : %v ", err)
		c.JSON(500, gin.H{
			"Message": "Somne erro ocurred",
		})
		return
	}

	c.JSON(200, gin.H{
		"Response":res.Response,
	})
}


func GetAllItems (c *gin.Context) {
	restaurantId := c.Param("restaurantid")
	fmt.Println(restaurantId)
	req := &restaurantpb.GetAllItemsRequest{
		RestaurantId:restaurantId,
	}

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Sorry client cannot talk to server: %v: ", err)
	}
	defer conn.Close()

	oc := restaurantpb.NewRestaurantServiceClient(conn)

	res, err := oc.GetAllItems(context.Background(), req)
	if err != nil {
		log.Fatalf("Error While calling UpdateDish : %v ", err)
		c.JSON(500, gin.H{
			"Message": "Somne erro ocurred",
		})
		return
	}

	c.JSON(200, gin.H{
		"Items":res.Items,
	})
}


func main(){
	router := gin.Default()

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	api:= router.Group("/api")


	// http://localhost:5656/api/
	api.GET("/",  HomePage)

	api.POST("/login", Login)

	//Order API's
	api.POST("/order", PlaceOrder)
	api.PUT("/order", UpdateOrder)
	api.DELETE("/customer/:customerid/order/:id", DeleteOrder)
	api.GET("/customer/:customerid/order/:orderid",GetSpecificOrder)


	//Customer API's
	api.GET("/customer/:customerid",GetSpecificCustomer)
	api.GET("/customers",GetCustomerCount)
	api.POST("/customer", AddCustomer)

	//Restaurant API's
	api.POST("/restaurant",AddRestaurant)
	api.GET("/restaurants",GetAllRestaurants)
	api.GET("/restaurant/:restaurantid",GetSpecificRestaurant)
	api.POST("/restaurant/:restaurantid/item",AddSpecificItem)
	api.PUT("restaurant/:restaurantid/item",UpdateRestaurantItem)
	api.GET("/restaurant/:restaurantid/items",GetAllItems)
	api.DELETE("restaurant/:restaurantid/item/:itemid",DeleteRestaurant)

	router.Run("localhost:5656")
}