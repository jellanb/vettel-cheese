package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	CreatedDate       primitive.DateTime `json:"createdDate"`
	UpdateDate        primitive.DateTime `json:"updateDate"`
	Description       string             `json:"description"`
	Barcode           string             `json:"barcode"`
	Quantity          float32            `json:"quantity"`
	SaleAmount        int                `json:"saleAmount"`
	UnitOfMeasurement string             `json:"unitOfMeasurement"`
}
