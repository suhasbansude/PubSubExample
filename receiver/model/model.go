package model

type DataHolder struct {
	Offers []Offer `json:"offers"`
}

type Offer struct {
	CmOfferID    string       `json:"cm_offer_id"`
	Hotel        Hotel        `json:"hotel"`
	Room         Room         `json:"room"`
	RatePlan     RatePlan     `json:"rate_plan"`
	OriginalData OriginalData `json:"original_data"`
	Capacity     Capacity     `json:"capacity"`
	Number       int          `json:"number"`
	Price        int          `json:"price"`
	Currency     string       `json:"currency"`
	CheckIn      string       `json:"check_in"`
	CheckOut     string       `json:"check_out"`
	Fees         []Fee        `json:"fees"`
}

type Fee struct {
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Included    bool    `json:"included"`
	Percent     float64 `json:"percent"`
}

type OriginalData struct {
	GuaranteePolicy GuaranteePolicy `json:"GuaranteePolicy"`
}

type GuaranteePolicy struct {
	Required bool
}

type Hotel struct {
	HotelID       string     `json:"hotel_id" gorm:"primary_key"`
	Name          string     `json:"name"`
	Country       string     `json:"country"`
	Address       string     `json:"address"`
	Latitude      float64    `json:"latitude"`
	Longitude     float64    `json:"longitude"`
	Telephone     string     `json:"telephone"`
	AmenitiesList []string   `json:"amenities" gorm:"-"`
	Amenities     []Amenity  `gorm:"foreignKey:HotelID"`
	Description   string     `json:"description"`
	RoomCount     int        `json:"room_count"`
	Currency      string     `json:"currency"`
	Rooms         []Room     `json:"rooms" gorm:"foreignKey:HotelID"`
	RatePlans     []RatePlan `json:"ratePlan" gorm:"foreignKey:HotelID"`
}

type Amenity struct {
	AmenitiyID  int64  `json:"amenitiyId" gorm:"primary_key;auto_increment;not_null"`
	HotelID     string `json:"hotel_id"`
	AmenityName string `json:"amenitiy"`
}

type Capacity struct {
	HotelID       string `json:"hotel_id" gorm:"primaryKey"`
	RoomID        string `json:"room_id" gorm:"primaryKey"`
	MaxAdults     int    `json:"max_adults"`
	ExtraChildren int    `json:"extra_children"`
}

type Room struct {
	HotelID     string   `json:"hotel_id" gorm:"primaryKey"`
	RoomID      string   `json:"room_id" gorm:"primaryKey"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Capacity    Capacity `json:"capacity" gorm:"foreignKey:HotelID,RoomID"`
}

type RatePlan struct {
	HotelID              string               `json:"hotel_id" gorm:"primary_key"`
	RatePlanID           string               `json:"rate_plan_id" gorm:"primary_key"`
	CancellationPolicies []CancellationPolicy `json:"cancellation_policy" gorm:"foreignKey:HotelID,RatePlanID"`
	name                 string               `json:"name"`
	OtherConditionsList  []string             `json:"other_conditions" gorm:"-"`
	OtherConditions      []OtherCondition     `json:"-" gorm:"foreignKey:HotelID,RatePlanID"`
	MealPlan             string               `json:"meal_plan"`
}

type CancellationPolicy struct {
	ID                int64  `json:"id" gorm:"primary_key;auto_increment;not_null"`
	HotelID           string `json:"hotel_id"`
	RatePlanID        string `json:"rate_plan_id"`
	Type              string `json:"type"`
	ExpiresDaysBefore int    `json:"expires_days_before"`
}

type OtherCondition struct {
	ID         int64  `json:"id" gorm:"primary_key;auto_increment;not_null"`
	HotelID    string `json:"hotel_id"`
	RatePlanID string `json:"rate_plan_id"`
	Condition  string `json:"condition"`
}
