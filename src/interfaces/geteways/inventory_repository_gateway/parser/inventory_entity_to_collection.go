package parser

import (
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/database/collections"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway/parser"
)

type InventoryEntityToCollectionInterfaces interface {
	InventoryEntityToCollection(entityInventory entity.Inventory) *collections.Inventory
}

type InventoryEntityToCollection struct{}

func NewInventoryEntityToCollection() InventoryEntityToCollectionInterfaces {
	return &InventoryEntityToCollection{}
}

func (i InventoryEntityToCollection) InventoryEntityToCollection(entityInventory entity.Inventory) *collections.Inventory {
	var inventoryColl collections.Inventory
	inventoryColl.Quantity = entityInventory.Quantity
	inventoryColl.Product = parser.ProductEntityToCollection(entityInventory.Product)
	inventoryColl.Date = entityInventory.Date
	inventoryColl.UpdateDate = entityInventory.UpdateDate
	inventoryColl.SaleAmount = entityInventory.SaleAmount
	return &inventoryColl
}
