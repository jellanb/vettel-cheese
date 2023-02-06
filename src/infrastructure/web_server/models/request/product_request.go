package request

import (
	"time"
)

type Product struct {
	CreatedDate       time.Time `json:"createdDate"`
	UpdateDate        time.Time `json:"updateDate"`
	Description       string    `json:"description"`
	Barcode           string    `json:"barcode"`
	Quantity          float32   `json:"quantity"`
	SaleAmount        int       `json:"saleAmount"`
	UnitOfMeasurement string    `json:"unitOfMeasurement"`
}
