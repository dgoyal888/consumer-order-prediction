package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dgoyal888/consumer-order-prediction/pkg/csv"
	"github.com/dgoyal888/consumer-order-prediction/pkg/rules"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World from GIN to Swiggy",
	})
}

//Return most popular restaurant.
func GetPoplarRestaurant(c *gin.Context) {
	restaurant, _ := rules.PopularRestaurant("../../data/orderdata.json")
	c.JSON(200, gin.H{
		"Most Popular Restaurant": restaurant.RestsurantName,
	})
}

func GetPopularCuisine(c *gin.Context) {
	restaurant, _ := rules.PopularVegCuisine("../../data/orderdata.json")
	c.JSON(200, gin.H{
		"Most Popular Cuisine is": restaurant.VegCuisine,
	})
}

//Return a specific order on the basis of customer id.
func GetSpecificOrdersByQuery(c *gin.Context) {

	customerid := c.Query("CustomerID")
	order, err := rules.ReturnJsonBasedOnCUSTID(customerid)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "customer not found",
		})
	} else {
		c.JSON(200, &order)
	}
}

//Add an order in orderdataapi.json file
func PostOrder(c *gin.Context) {
	body := c.Request.Body

	content, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("Sorry No Content :", err.Error())
	}
	var order csv.Order
	json.Unmarshal(content, &order)

	rules.Appendtofile(&order)
	c.JSON(http.StatusCreated, gin.H{
		"message": string(content),
	})
}

func main() {
	router := gin.Default()

	api := router.Group("/api", gin.BasicAuth(gin.Accounts{
		"team1": "team1",
	}))
	// http://localhost:5656/api/
	api.GET("/", HomePage)
	api.GET("/popularrestaurant", GetPoplarRestaurant)
	api.GET("/popularcuisine", GetPopularCuisine)
	api.GET("/orders", GetSpecificOrdersByQuery)
	api.POST("/order", PostOrder)
	router.Run("0.0.0.0:5656")
}
func writetojson() {

}
