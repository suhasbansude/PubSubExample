package service

import (
	"../dao"
	"../model"
)

func SaveOfferDataService(myData model.DataHolder) error {
	for i := 0; i < len(myData.Offers); i++ {
		for _, aminity := range myData.Offers[i].Hotel.AmenitiesList {
			a := model.Amenity{}
			a.AmenityName = aminity
			a.HotelID = myData.Offers[i].Hotel.HotelID
			// db.Create(&a)
			myData.Offers[i].Hotel.Amenities = append(myData.Offers[i].Hotel.Amenities, a)
		}

		for _, condition := range myData.Offers[i].RatePlan.OtherConditionsList {
			c := model.OtherCondition{}
			c.Condition = condition
			myData.Offers[i].RatePlan.OtherConditions = append(myData.Offers[i].RatePlan.OtherConditions, c)
		}
	}
	myData.Offers[0].Room.Capacity.HotelID = myData.Offers[0].Room.HotelID
	myData.Offers[0].Room.Capacity.RoomID = myData.Offers[0].Room.RoomID
	myData.Offers[0].Hotel.Rooms = append(myData.Offers[0].Hotel.Rooms, myData.Offers[0].Room)
	myData.Offers[0].Hotel.RatePlans = append(myData.Offers[0].Hotel.RatePlans, myData.Offers[0].RatePlan)
	return dao.SaveOfferDataDAO(myData)
}
