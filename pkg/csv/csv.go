package csv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type CSV struct {
	CustomerID string
	CustomerName string
	Cuisine string
}

func ReadCSVData(filePath string) ([]CSV,error){

	csvfile, err := os.Open(filePath)

	if err != nil {
		return nil,err
	}

	defer csvfile.Close()

	csvData,err := csv.NewReader(csvfile).ReadAll()

	if err != nil {
		return nil,err
	}

	csvDataSlice := []CSV{}

	for i, csv := range csvData {
		if i == 0 {
			continue
		}
		if i == 5 {
			break
		}
		csv := CSV{csv[0],csv[1],csv[2]}
		csvDataSlice = append(csvDataSlice,csv)
	}

	return csvDataSlice,nil
}

func CSVToJSON(csvData []CSV) error {

	jsonData, err := json.MarshalIndent(csvData,"","    ")

	if err != nil {
		return err
	}

	fmt.Printf("%s\n",jsonData)

	jsonFile, err := os.Create("data/orderdata.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	jsonFile.Write(jsonData)

	return nil
}
