package parser

import (
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
)

type InventoryEntityToResponseInterfaces interface {
	SliceInventoryEntityToResponse(entityInventories []*entity.Inventory) []response.Inventory
	InventoryEntityToResponse(entityInventories *entity.Inventory) response.Inventory
}

type InventoryEntityToResponse struct{}

func NewInventoryEntityToResponse() InventoryEntityToResponseInterfaces {
	return &InventoryEntityToResponse{}
}

func (i InventoryEntityToResponse) SliceInventoryEntityToResponse(entityInventories []*entity.Inventory) []response.Inventory {
	var inventoriesResponse []response.Inventory
	for _, entity := range entityInventories {
		var inventoryResponse response.Inventory
		inventoryResponse.Date = entity.Date
		inventoryResponse.UpdateDate = entity.UpdateDate
		inventoryResponse.Quantity = entity.Quantity
		inventoryResponse.Product = response.Product(entity.Product)
		inventoryResponse.SaleAmount = entity.SaleAmount
		inventoriesResponse = append(inventoriesResponse, inventoryResponse)
	}
	return inventoriesResponse
}

func (i InventoryEntityToResponse) InventoryEntityToResponse(entityInventories *entity.Inventory) response.Inventory {
	var inventoryResponse response.Inventory
	inventoryResponse.Date = entityInventories.Date
	inventoryResponse.UpdateDate = entityInventories.UpdateDate
	inventoryResponse.Quantity = entityInventories.Quantity
	inventoryResponse.Product = response.Product(entityInventories.Product)

	return inventoryResponse
}
