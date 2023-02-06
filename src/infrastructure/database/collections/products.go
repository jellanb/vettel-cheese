package collections

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	CreatedDate       primitive.DateTime `bson:"createdDate"`
	UpdateDate        primitive.DateTime `bson:"updateDate"`
	Description       string             `bson:"description"`
	Barcode           string             `bson:"barcode"`
	Quantity          float32            `bson:"quantity"`
	SaleAmount        int                `bson:"saleAmount"`
	UnitOfMeasurement string             `bson:"unitOfMeasurement"`
}
