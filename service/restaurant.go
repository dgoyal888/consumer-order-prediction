package service

import (
	"context"
	"encoding/json"
	"github.com/consumer-order-prediction/pkg/dynamodb"
	restaurantpb "github.com/consumer-order-prediction/pkg/proto/restaurant"
	"github.com/consumer-order-prediction/util"
)

type Ritem struct {
	ItemID   string
	Name     string
	Price     float32
}

type Restaurant struct {
	RestaurantID   string
	RestaurantName      string
	Address string
	Ritems        []*Ritem
}


func (s *Service) AddRestaurant(ctx context.Context, req *restaurantpb.AddRestaurantRequest) (*restaurantpb.AddRestaurantResponse, error) {

	restaurant := req.GetRestaurant()

	var ritemSlice []*Ritem
	var ritem *Ritem
	for _,z := range restaurant.Items{
		ritem = &Ritem{
			ItemID: z.ItemId,
			Name: z.Name,
			Price: z.Price,
		}
		ritemSlice = append(ritemSlice,ritem)
	}

	var restaurantStruct = &Restaurant{
		RestaurantID: restaurant.RestaurantId,
		RestaurantName: restaurant.Name,
		Address: restaurant.Address,
		Ritems: ritemSlice,
	}

	err := AddRestaurant(restaurantStruct)

	if err != nil {
		return & restaurantpb.AddRestaurantResponse{
			Response:"Error occurred while adding restaurant",
		},err
	}

	res := & restaurantpb.AddRestaurantResponse{
		Response:"Restaurant added successfully",
	}

	return res,nil
}

func (s *Service) GetAllRestaurant(ctx context.Context, req *restaurantpb.GetAllRestaurantRequest) (*restaurantpb.GetAllRestaurantResponse, error) {

	res, err := dynamodb.GetAllItem("t1-restaurant")
	if err != nil {
		return & restaurantpb.GetAllRestaurantResponse{
		},err
	}

	return & restaurantpb.GetAllRestaurantResponse{
		Count: res,
	},nil
}

func (s *Service) GetSpecificRestaurant(ctx context.Context, req *restaurantpb.GetSpecificRestaurantRequest) (*restaurantpb.GetSpecificRestaurantResponse, error) {

	restaurantId := req.RestaurantId

	restaurant, err := GetRestaurant(restaurantId)

	if err != nil {
		return nil,err
	}

	var itempbSlice []*restaurantpb.Item

	for _,item := range restaurant.Ritems {
		itempb := &restaurantpb.Item{
			ItemId:item.ItemID,
			Name:item.Name,
			Price: item.Price,
		}
		itempbSlice = append(itempbSlice,itempb)
	}


	resp := &restaurantpb.GetSpecificRestaurantResponse{
		Restaurant:&restaurantpb.Restaurant{
			RestaurantId:restaurant.RestaurantID,
			Name:restaurant.RestaurantName,
			Address:restaurant.Address,
			Items: itempbSlice,
		},
	}
	return resp,nil
}


func (s *Service) AddRestaurantItem(ctx context.Context, req *restaurantpb.AddRestaurantItemRequest) (*restaurantpb.AddRestaurantItemResponse, error) {

	restaurant,err:=GetRestaurant(req.RestaurantId)
	if err!=nil{
		return &restaurantpb.AddRestaurantItemResponse{},err
	}

	itemID,err := util.GenerateUUID()

	if err!=nil{
		return &restaurantpb.AddRestaurantItemResponse{},err
	}

	itemRequest := &Ritem{
		ItemID: itemID,
		Name:req.ItemName,
		Price:req.Price,
	}

	restaurant.Ritems=append(restaurant.Ritems,itemRequest)

	err = AddRestaurant(restaurant)

	if err!=nil{
		return &restaurantpb.AddRestaurantItemResponse{},err
	}

	return &restaurantpb.AddRestaurantItemResponse{
		Response:"Item Added Successfully",
	},nil
}

func (s *Service) UpdateRestaurantItem(ctx context.Context, req *restaurantpb.UpdateRestaurantItemRequest) (*restaurantpb.UpdateRestaurantItemResponse, error) {

	restaurant,err:=GetRestaurant(req.RestaurantId)
	if err!=nil{
		return &restaurantpb.UpdateRestaurantItemResponse{

		},err
	}
	itemPresent:=false

	itemRequest := &Ritem{
		ItemID:req.ItemId,
		Name:req.ItemName,
		Price:req.Price,
	}
	for i,item:=range restaurant.Ritems{
		if item.ItemID==req.ItemId{
			restaurant.Ritems[i]=itemRequest
			itemPresent=true
			break
		}
	}
	if !itemPresent{
		restaurant.Ritems=append(restaurant.Ritems,itemRequest)
	}

	err = AddRestaurant(restaurant)

	if err!=nil{
		return &restaurantpb.UpdateRestaurantItemResponse{},err
	}

	return &restaurantpb.UpdateRestaurantItemResponse{
		Response:"Item Updated Successfully",
	},nil
}

func (s *Service) DeleteItem(ctx context.Context, req *restaurantpb.DeleteItemRequest) (*restaurantpb.DeleteItemResponse, error) {

	restaurant,err:=GetRestaurant(req.RestaurantId)
	if err!=nil{
		return &restaurantpb.DeleteItemResponse{},err
	}

	itemPresent := false
	var itemIndex int
	for i,item:=range restaurant.Ritems{
		if item.ItemID == req.ItemId{
			itemIndex=i
			itemPresent = true
			break
		}
	}

	if !itemPresent {
		return &restaurantpb.DeleteItemResponse{
			Response:"Item is not present",
		},nil
	}
	restaurant.Ritems= append(restaurant.Ritems[:itemIndex], restaurant.Ritems[itemIndex+1:]...)

	err= AddRestaurant(restaurant)

	if err!=nil{
		return &restaurantpb.DeleteItemResponse{},err
	}

	return &restaurantpb.DeleteItemResponse{
		Response:"Item Deleted Successfully",
	},nil
}

func (s *Service) GetAllItems(ctx context.Context, req *restaurantpb.GetAllItemsRequest) (*restaurantpb.GetAllItemsResponse, error) {

	restaurant,err := GetRestaurant(req.RestaurantId)

	if err!=nil {
		return &restaurantpb.GetAllItemsResponse{

		},err
	}

	var itempbSlice []*restaurantpb.Item

	for _,item := range restaurant.Ritems {
		itempb := &restaurantpb.Item{
			ItemId:item.ItemID,
			Name:item.Name,
			Price: item.Price,
		}
		itempbSlice = append(itempbSlice,itempb)
	}

	return &restaurantpb.GetAllItemsResponse{
		Items:itempbSlice,
	},nil
}


func (s *Service) GetSpecificItem(ctx context.Context, req *restaurantpb.GetSpecificItemRequest) (*restaurantpb.GetSpecificItemResponse, error) {

	restaurant,err := GetRestaurant(req.RestaurantId)

	if err != nil {
		return &restaurantpb.GetSpecificItemResponse{},err
	}

	var presentItem *Ritem
	for _,item:=range restaurant.Ritems{
		if item.ItemID==req.ItemId{
			presentItem = item
			break
		}
	}

	item := &restaurantpb.Item{
		ItemId:presentItem.ItemID,
		Name:presentItem.Name,
		Price:presentItem.Price,
	}
	return &restaurantpb.GetSpecificItemResponse{
		Item:item,
	},nil
}

func GetRestaurant(restaurantId string) (*Restaurant,error){

	res, err := dynamodb.GetItem("t1-restaurant","RestaurantID",restaurantId,"RestaurantID",restaurantId,&Restaurant{})

	if err != nil {
		return &Restaurant{},err
	}

	resJSON,err := json.Marshal(res)

	if err != nil {
		return &Restaurant{},err
	}

	var restaurant *Restaurant
	err = json.Unmarshal(resJSON, &restaurant)

	if err != nil {
		return &Restaurant{},err
	}

	return restaurant,nil
}

func AddRestaurant(restaurant *Restaurant) (error){

	err := dynamodb.PutItem("t1-restaurant",restaurant)

	return err
}