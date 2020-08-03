package service

import (
	"context"
	"github.com/consumer-order-prediction/pkg/dynamodb"
	restaurantpb "github.com/consumer-order-prediction/pkg/proto/restaurant"
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

	err := dynamodb.PutItem("restaurantdemo",restaurantStruct)

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

	res, err := dynamodb.GetAllItem("restaurantdemo")
	if err != nil {
		return & restaurantpb.GetAllRestaurantResponse{
		},err
	}

	return & restaurantpb.GetAllRestaurantResponse{
		Count: res,
	},nil
}

func (s *Service) GetSpecificRestaurant(ctx context.Context, req *restaurantpb.GetSpecificRestaurantRequest) (*restaurantpb.GetSpecificRestaurantResponse, error) {


	return nil,nil
}


func (s *Service) AddRestaurantItem(ctx context.Context, req *restaurantpb.AddRestaurantItemRequest) (*restaurantpb.AddRestaurantItemResponse, error) {


	return nil,nil
}

func (s *Service) UpdateRestaurantItem(ctx context.Context, req *restaurantpb.UpdateRestaurantItemRequest) (*restaurantpb.UpdateRestaurantItemResponse, error) {


	return nil,nil
}

func (s *Service) DeleteItem(ctx context.Context, req *restaurantpb.DeleteItemRequest) (*restaurantpb.DeleteItemResponse, error) {


	return nil,nil
}

func (s *Service) GetAllItems(ctx context.Context, req *restaurantpb.GetAllItemsRequest) (*restaurantpb.GetAllItemsResponse, error) {


	return nil,nil
}


func (s *Service) GetSpecificItem(ctx context.Context, req *restaurantpb.GetSpecificItemRequest) (*restaurantpb.GetSpecificItemResponse, error) {


	return nil,nil
}