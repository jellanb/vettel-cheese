package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type Inventory struct {
	Date       primitive.DateTime `json:"date"`
	UpdateDate primitive.DateTime `json:"updateDate"`
	Product    Product            `json:"product"`
	Quantity   int                `json:"quantity"`
	SaleAmount int                `json:"saleAmount"`
}
