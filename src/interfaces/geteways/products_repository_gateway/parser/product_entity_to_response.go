package parser

import (
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
)

func ProductEntityToResponse(productEntity *entity.Product) *response.Product {
	var productRep response.Product
	productRep.Barcode = productEntity.Barcode
	productRep.Description = productEntity.Description
	productRep.SaleAmount = productEntity.SaleAmount
	productRep.Quantity = productEntity.Quantity
	productRep.CreatedDate = productEntity.CreatedDate
	productRep.UpdateDate = productEntity.UpdateDate
	productRep.UnitOfMeasurement = productEntity.UnitOfMeasurement
	return &productRep
}
