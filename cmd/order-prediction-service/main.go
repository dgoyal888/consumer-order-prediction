package main

import (
	"consumer-order-prediction/pkg/csv"
	"consumer-order-prediction/pkg/rules"
	"fmt"
)


func main() {

	csvData, err := csv.ReadCSVData("data/orderdata.csv")

	if err != nil {
		fmt.Println("Error while reading CSV data: %s",err.Error())
	}

	err = csv.CSVToJSON(csvData)

	if err != nil {
		fmt.Println("Error while converting  CSV to Json data: %s",err.Error())
	}

	_,err = rules.PopularRestaurant("data/orderdata.json")

	if err != nil {
		fmt.Println("Error while finding popular restaurants: %s",err.Error())
	}

	_,err = rules.PopularVegCuisine("data/orderdata.json")

	if err != nil {
		fmt.Println("Error while finding popular veg cuisine: %s",err.Error())
	}

	_,err = rules.PopularNonVegCuisine("data/orderdata.json")

	if err != nil {
		fmt.Println("Error while finding popular non-veg cuisine: %s",err.Error())
	}

}

