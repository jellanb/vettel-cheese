package parser

import (
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/database/collections"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway/parser"
)

type InventoryCollectionToEntityInterfaces interface {
	InventoryCollectionsToEntities(inventoryCollections []*collections.Inventory) []*entity.Inventory
	InventoryCollectionToEntity(inventoryCollections *collections.Inventory) *entity.Inventory
}

type InventoryCollectionEntity struct{}

func NewInventoryCollectionEntity() InventoryCollectionToEntityInterfaces {
	return &InventoryCollectionEntity{}
}

func (i InventoryCollectionEntity) InventoryCollectionsToEntities(inventoryCollections []*collections.Inventory) []*entity.Inventory {
	var entitiesInventory []*entity.Inventory
	for _, inventoryCollection := range inventoryCollections {
		var entityInventory entity.Inventory
		entityInventory.Date = inventoryCollection.Date
		entityInventory.UpdateDate = inventoryCollection.UpdateDate
		entityInventory.Quantity = inventoryCollection.Quantity
		entityInventory.Product = parser.ProductCollectionToEntity(inventoryCollection.Product)
		entityInventory.SaleAmount = inventoryCollection.SaleAmount
		entitiesInventory = append(entitiesInventory, &entityInventory)
	}

	return entitiesInventory
}

func (i InventoryCollectionEntity) InventoryCollectionToEntity(inventoryCollections *collections.Inventory) *entity.Inventory {
	var entityInventory entity.Inventory
	entityInventory.Date = inventoryCollections.Date
	entityInventory.UpdateDate = inventoryCollections.UpdateDate
	entityInventory.Quantity = inventoryCollections.Quantity
	entityInventory.Product = parser.ProductCollectionToEntity(inventoryCollections.Product)

	return &entityInventory
}
