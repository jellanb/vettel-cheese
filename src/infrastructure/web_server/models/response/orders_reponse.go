package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	OrderNumber int
	Date        primitive.DateTime
	SellerName  string
	Status      string
	SaleAmount  int
}
