package rules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"consumer-order-prediction/pkg/csv"
)

func PopularRestaurant(filePath string) (csv.Order,error) {

	jsonFile, err := os.Open(filePath)
	var popularrestaurantobject csv.Order
	if err != nil {
		return popularrestaurantobject,err
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return popularrestaurantobject,err
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
			popularrestaurantobject=orderData[i]
		}
	}

	fmt.Printf("Most Popular Restaurant is %s",popularRestaurantName)
	return popularrestaurantobject,nil
}

func PopularVegCuisine(filePath string) (csv.Order,error) {

	jsonFile, err := os.Open(filePath)
	var popularvegcuisineobject csv.Order
	if err != nil {
		return popularvegcuisineobject,err
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return popularvegcuisineobject,err
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
			popularvegcuisineobject = orderData[i]
		}
	}

	fmt.Printf("Most Popular Veg cuisine is %s",popularCuisineName)

	return popularvegcuisineobject,nil
}

func PopularNonVegCuisine(filePath string) (csv.Order,error) {

	jsonFile, err := os.Open(filePath)
	var ppopularVegCuisine csv.Order
	if err != nil {
		return ppopularVegCuisine,nil
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return ppopularVegCuisine,err
	}

	var orderData []csv.Order
	err = json.Unmarshal(byteValue, &orderData)

	popularCuisines := make(map[string]int)

	var popularCuisineFreq int
	var popularCuisineName string
	//var ppopularVegCuisine csv.Order

	for i := 0; i < len(orderData); i++ {
		if len(orderData[i].NonVegCuisine) != 0 {
			popularCuisines[orderData[i].NonVegCuisine]++
			if popularCuisines[orderData[i].NonVegCuisine] > popularCuisineFreq {
			popularCuisineFreq = popularCuisines[orderData[i].NonVegCuisine]
			popularCuisineName = orderData[i].NonVegCuisine
			ppopularVegCuisine = orderData[i]
			}
		}
	}

	fmt.Printf("Most Popular Non-Veg Cuisine is %s",popularCuisineName)

	return ppopularVegCuisine,nil
}
func ReturnJsonBasedOnCUSTID(custid string) (csv.Order,error){
	jsonFile, err := os.Open("../../data/orderdata.json")
	if err!=nil{
		return csv.Order{},fmt.Errorf("Error finding file")
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err!=nil{
		return csv.Order{},fmt.Errorf("Error reading file")
	}
	var orderData []csv.Order
	err = json.Unmarshal(byteValue, &orderData)
	for i := 0; i < len(orderData); i++ {
		cust,_:=strconv.ParseInt(custid,10,64)
		if orderData[i].CustomerID==cust{
			return orderData[i],nil
		}
	}
	return csv.Order{},fmt.Errorf("Error finding id")
}

func Appendtofile(order *csv.Order) error{
	jsonFile, err := os.Open("data/orderdataapi.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return err
	}

	var orderData []csv.Order
	json.Unmarshal(byteValue,&orderData)
	orderData=append(orderData,*order)
	jsonData, err := json.MarshalIndent(orderData,"","    ")
	if err != nil {
		return err
	}
	jsonFile, err= os.Create("data/orderdataapi.json")

	if err != nil {
		return err
	}

	defer jsonFile.Close()

	jsonFile.Write(jsonData)
	return nil
}