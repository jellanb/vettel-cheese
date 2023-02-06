package parser

import (
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/database/collections"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway/parser"
)

func DispatchEntityToCollection(entity entity.Dispatch) collections.Dispatch {
	var collection collections.Dispatch
	collection.CreditTime = entity.CreditTime
	collection.Address = entity.Address
	collection.AddressNumber = entity.AddressNumber
	collection.Commune = entity.Commune
	collection.Region = entity.Region
	collection.DeliveryDate = entity.DeliveryDate
	collection.Phone = entity.Phone
	collection.DeliveryName = entity.DeliveryName
	collection.Products = parser.ProductsEntitiesToCollection(entity.Products)
	collection.Status = entity.Status
	collection.AdministratorName = entity.AdministratorName
	collection.ReasonSocial = entity.ReasonSocial
	collection.CreatedDate = entity.CreatedDate
	return collection
}
