package request

import (
	"time"
)

type Inventory struct {
	Date     time.Time `json:"date"`
	Product  Product   `json:"product"`
	Quantity int       `json:"quantity"`
}
