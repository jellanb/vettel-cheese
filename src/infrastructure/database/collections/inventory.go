package collections

import "go.mongodb.org/mongo-driver/bson/primitive"

type Inventory struct {
	Date       primitive.DateTime `bson:"date"`
	UpdateDate primitive.DateTime `bson:"updateDate"`
	Product    Product            `bson:"product"`
	Quantity   int                `bson:"quantity"`
}
