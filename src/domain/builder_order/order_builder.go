package builder_order

import (
	"vettel-backend-app/src/domain/entity"
)

type OrderBuilderInterfaces interface {
	BuildNewOrder() entity.Order
}

type OrderBuilder struct {
	order entity.Order
}

func NewOrderBuilder() OrderBuilderInterfaces {
	return &OrderBuilder{}
}

func (o *OrderBuilder) BuildNewOrder() entity.Order {
	return o.order
}
