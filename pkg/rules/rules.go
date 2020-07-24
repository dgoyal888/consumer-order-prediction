package rules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"consumer-order-prediction/pkg/csv"
)

func PopularRestaurant(filePath string) error {

	jsonFile, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return err
	}

	var orderData []csv.Order

	err = json.Unmarshal(byteValue, &orderData)

	popularRestaurants := make(map[string]int)

	var popularRestaurantFreq int
	var popularRestaurantName string
	for i := 0; i < len(orderData); i++ {
		popularRestaurants[orderData[i].RestsurantName]++
		if popularRestaurants[orderData[i].RestsurantName] > popularRestaurantFreq {
			popularRestaurantFreq =  popularRestaurants[orderData[i].RestsurantName]
			popularRestaurantName = orderData[i].RestsurantName
		}
	}

	fmt.Println("Most Popular Restaurant is %s",popularRestaurantName)
	return nil
}

func PopularVegCuisine(filePath string) error {

	jsonFile, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return err
	}

	var orderData []csv.Order
	err = json.Unmarshal(byteValue, &orderData)

	popularCuisines := make(map[string]int)

	var popularCuisineFreq int
	var popularCuisineName string

	for i := 0; i < len(orderData); i++ {
		popularCuisines[orderData[i].VegCuisine]++
		if popularCuisines[orderData[i].VegCuisine] > popularCuisineFreq {
			popularCuisineFreq =  popularCuisines[orderData[i].VegCuisine]
			popularCuisineName = orderData[i].VegCuisine
		}
	}

	fmt.Println("Most Popular Veg cuisine is %s",popularCuisineName)

	return nil
}

func PopularNonVegCuisine(filePath string) error {

	jsonFile, err := os.Open(filePath)

	if err != nil {
		return err
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return err
	}

	var orderData []csv.Order
	err = json.Unmarshal(byteValue, &orderData)

	popularCuisines := make(map[string]int)

	var popularCuisineFreq int
	var popularCuisineName string

	for i := 0; i < len(orderData); i++ {
		if len(orderData[i].NonVegCuisine) != 0 {
			popularCuisines[orderData[i].NonVegCuisine]++
			if popularCuisines[orderData[i].NonVegCuisine] > popularCuisineFreq {
			popularCuisineFreq = popularCuisines[orderData[i].NonVegCuisine]
			popularCuisineName = orderData[i].NonVegCuisine
			}
		}
	}

	fmt.Println("Most Popular Non-Veg Cuisine is %s",popularCuisineName)

	return nil
}
