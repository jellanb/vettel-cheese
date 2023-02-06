package parser

import (
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
)

func OrderEntityToResponse(orderEntity entity.Order) response.Order {
	var orderResponse response.Order
	orderResponse.OrderNumber = orderEntity.OrderNumber
	orderResponse.SaleAmount = orderEntity.SaleAmount
	orderResponse.Status = orderEntity.Status
	orderResponse.Date = orderEntity.Date
	return orderResponse
}
