package parser

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/domain/status_domain/dispatch_status"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
)

func DispatchRequestToEntity(requestDispatch request.Dispatch) entity.Dispatch {
	var dispatch entity.Dispatch
	dispatch.CreatedDate = primitive.NewDateTimeFromTime(time.Now())
	dispatch.Address = requestDispatch.Address
	dispatch.DeliveryDate = requestDispatch.DeliveryDate
	dispatch.ReasonSocial = requestDispatch.ReasonSocial
	dispatch.Status = dispatch_status.Pending
	dispatch.AddressNumber = requestDispatch.AddressNumber
	dispatch.Commune = requestDispatch.Commune
	dispatch.Region = requestDispatch.Region
	dispatch.AdministratorName = requestDispatch.AdministratorName
	dispatch.DeliveryName = requestDispatch.DeliveryName
	dispatch.Phone = requestDispatch.Phone
	dispatch.CreditTime = requestDispatch.CreditTime
	dispatch.Products = DispatchRequestProductsToEntity(requestDispatch.Products)
	dispatch.Amount = calculateAmount(requestDispatch.Products)
	return dispatch
}

func DispatchRequestProductsToEntity(requestDispatchProducts [][]request.Product) []entity.Product {
	var productsEntity []entity.Product
	for _, global := range requestDispatchProducts {
		for _, requestProduct := range global {
			var productEntity entity.Product
			productEntity.Barcode = requestProduct.Barcode
			productEntity.Quantity = requestProduct.Quantity
			productEntity.SaleAmount = requestProduct.SaleAmount
			productEntity.Description = requestProduct.Description
			productEntity.CreatedDate = primitive.NewDateTimeFromTime(requestProduct.CreatedDate)
			productsEntity = append(productsEntity, productEntity)
		}
	}
	return productsEntity
}

func calculateAmount(requestDispatchProducts [][]request.Product) int {
	amount := 0
	for _, global := range requestDispatchProducts {
		for _, requestProduct := range global {
			amount += requestProduct.SaleAmount
		}
	}
	return amount
}
