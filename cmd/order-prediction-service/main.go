package main

import (
	"consumer-order-prediction/pkg/csv"
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
}

