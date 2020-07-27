package main

import (
	orderspb "consumer-order-prediction/pkg/proto/orders"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World from GIN to Swiggy",
	})
}

func  GetPoplarRestaurant(c *gin.Context) {

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


func main(){
	router := gin.Default()

	api:= router.Group("/api",gin.BasicAuth(gin.Accounts{
		"team1": "team1",
	}))
	// http://localhost:5656/api/
	api.GET("/",  HomePage)
	api.GET("/getpopularrestaurant", GetPoplarRestaurant)

	router.Run("localhost:5656")
}