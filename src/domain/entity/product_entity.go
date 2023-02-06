package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	CreatedDate       primitive.DateTime
	UpdateDate        primitive.DateTime
	Description       string
	Barcode           string
	Quantity          float32
	SaleAmount        int
	UnitOfMeasurement string
}
