package rules

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"consumer-order-prediction/pkg/csv"
)

func PopularRestaurants(filePath string) error {

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

	for i := 0; i < len(orderData); i++ {
		popularRestaurants[orderData[i].RestsurantName]++
	}

	return nil
}

func PopularVegCuisines(filePath string) error {

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

	for i := 0; i < len(orderData); i++ {
		popularCuisines[orderData[i].VegCuisine]++
	}

	return nil
}

func PopularNonVegCuisines(filePath string) error {

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

	for i := 0; i < len(orderData); i++ {
		popularCuisines[orderData[i].NonVegCuisine]++
	}

	return nil
}
