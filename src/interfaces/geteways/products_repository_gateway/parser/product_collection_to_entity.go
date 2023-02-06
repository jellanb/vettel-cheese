package parser

import (
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/database/collections"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
)

func ProductsCollectionToEntity(productCollection []collections.Product) []entity.Product {
	var productsEntity []entity.Product
	for _, product := range productCollection {
		var productEntity entity.Product
		productEntity.UpdateDate = product.UpdateDate
		productEntity.Quantity = product.Quantity
		productEntity.CreatedDate = product.CreatedDate
		productEntity.Barcode = product.Barcode
		productEntity.SaleAmount = product.SaleAmount
		productEntity.Description = product.Description
		productEntity.UnitOfMeasurement = product.UnitOfMeasurement
		productsEntity = append(productsEntity, productEntity)
	}
	return productsEntity
}

func ProductCollectionToEntity(productCollection collections.Product) entity.Product {
	var productEntity entity.Product
	productEntity.UpdateDate = productCollection.UpdateDate
	productEntity.Quantity = productCollection.Quantity
	productEntity.CreatedDate = productCollection.CreatedDate
	productEntity.Barcode = productCollection.Barcode
	productEntity.SaleAmount = productCollection.SaleAmount
	productEntity.Description = productCollection.Description
	productEntity.UnitOfMeasurement = productCollection.UnitOfMeasurement
	return productEntity
}

func ProductsEntitiesToResponse(entityProducts []entity.Product) []response.Product {
	var responseList []response.Product
	for _, entityProd := range entityProducts {
		var product response.Product
		product.Quantity = entityProd.Quantity
		product.Barcode = entityProd.Barcode
		product.Description = entityProd.Description
		product.SaleAmount = entityProd.SaleAmount
		product.CreatedDate = entityProd.CreatedDate
		product.UpdateDate = entityProd.UpdateDate
		product.UnitOfMeasurement = entityProd.UnitOfMeasurement
		responseList = append(responseList, product)
	}
	return responseList
}

func ProductsCollectionsToEntities(productsCollection []collections.Product) []entity.Product {
	var productsEntities []entity.Product
	for _, productCollection := range productsCollection {
		var productEntity entity.Product
		productEntity.Barcode = productCollection.Barcode
		productEntity.SaleAmount = productCollection.SaleAmount
		productEntity.UpdateDate = productCollection.UpdateDate
		productEntity.Quantity = productCollection.Quantity
		productEntity.Description = productCollection.Description
		productEntity.CreatedDate = productCollection.CreatedDate
		productsEntities = append(productsEntities, productEntity)
	}
	return productsEntities
}
