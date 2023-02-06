package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Inventory struct {
	Date       primitive.DateTime
	UpdateDate primitive.DateTime
	Product    Product
	Quantity   int
	SaleAmount int
}
