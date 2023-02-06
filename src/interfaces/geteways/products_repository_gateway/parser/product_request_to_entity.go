package parser

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
)

func ProductRequestToEntity(productReq request.Product) entity.Product {
	var productEntity entity.Product
	productEntity.UpdateDate = primitive.NewDateTimeFromTime(time.Now())
	productEntity.CreatedDate = primitive.NewDateTimeFromTime(productReq.CreatedDate)
	productEntity.SaleAmount = productReq.SaleAmount
	productEntity.Description = productReq.Description
	productEntity.Barcode = productReq.Barcode
	productEntity.Quantity = productReq.Quantity
	productEntity.UnitOfMeasurement = productReq.UnitOfMeasurement
	return productEntity
}
