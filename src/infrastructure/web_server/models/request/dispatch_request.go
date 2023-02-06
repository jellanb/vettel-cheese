package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dispatch struct {
	ReasonSocial      string             `json:"reasonSocial"`
	DeliveryDate      primitive.DateTime `json:"deliveryDate"`
	CreditTime        primitive.DateTime `json:"creditTime"`
	Products          [][]Product        `json:"products"`
	Address           string             `json:"address"`
	Commune           string             `json:"commune"`
	AddressNumber     int                `json:"addressNumber"`
	Region            string             `json:"region"`
	Phone             string             `json:"phone"`
	DeliveryName      string             `json:"deliveryName"`
	AdministratorName string             `json:"administratorName"`
	CreatedDate       primitive.DateTime
}
