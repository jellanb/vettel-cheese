package parser

import (
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/database/collections"
)

func ProductEntityToCollection(productEntity entity.Product) collections.Product {
	var collectionProduct collections.Product
	collectionProduct.UpdateDate = productEntity.UpdateDate
	collectionProduct.CreatedDate = productEntity.CreatedDate
	collectionProduct.Quantity = productEntity.Quantity
	collectionProduct.Barcode = productEntity.Barcode
	collectionProduct.Description = productEntity.Description
	collectionProduct.SaleAmount = productEntity.SaleAmount
	collectionProduct.UnitOfMeasurement = productEntity.UnitOfMeasurement
	return collectionProduct
}

func ProductsEntitiesToCollection(productsEntities []entity.Product) []collections.Product {
	var productsCollections []collections.Product
	for _, productEntity := range productsEntities {
		var productCollection collections.Product
		productCollection.Barcode = productEntity.Barcode
		productCollection.SaleAmount = productEntity.SaleAmount
		productCollection.UpdateDate = productEntity.UpdateDate
		productCollection.Quantity = productEntity.Quantity
		productCollection.Description = productEntity.Description
		productCollection.CreatedDate = productEntity.CreatedDate
		productCollection.UnitOfMeasurement = productEntity.UnitOfMeasurement
		productsCollections = append(productsCollections, productCollection)
	}
	return productsCollections
}
