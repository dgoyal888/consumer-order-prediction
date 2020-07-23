package csv

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

type Order struct {
	CustomerID int64 `json:",omitempty"`
	CustomerName string `json:",omitempty"`
	RestsurantName string `json:",omitempty"`
	VegCuisine string `json:",omitempty"`
	NonVegCuisine string `json:",omitempty"`
	State string `json:",omitempty"`
}

func ReadCSVData(filePath string) ([]Order,error){

	csvfile, err := os.Open(filePath)

	if err != nil {
		return nil,err
	}

	defer csvfile.Close()

	csvData,err := csv.NewReader(csvfile).ReadAll()

	if err != nil {
		return nil,err
	}

	csvDataSlice := []Order{}

	for i, csv := range csvData {
		if i == 0 {
			continue
		}
		if i == 5 {
			break
		}
		customerID,_ :=  strconv.ParseInt(csv[1],10,32)
		order := Order{customerID,csv[2],csv[3],csv[4],csv[5],csv[12]}
		csvDataSlice = append(csvDataSlice,order)
	}

	return csvDataSlice,nil
}

func CSVToJSON(csvData []Order) error {

	jsonData, err := json.MarshalIndent(csvData,"","    ")

	if err != nil {
		return err
	}

	jsonFile, err := os.Create("data/orderdata.json")

	if err != nil {
		return err
	}

	defer jsonFile.Close()

	jsonFile.Write(jsonData)

	return nil
}
