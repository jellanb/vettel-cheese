package parser

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway/parser"
)

type InventoryRequestToEntityInterfaces interface {
	InventoryRequestToEntity(reqInventory request.Inventory) entity.Inventory
}

type InventoryRequestToEntity struct{}

func NewInventoryRequestToEntity() InventoryRequestToEntityInterfaces {
	return &InventoryRequestToEntity{}
}

func (i InventoryRequestToEntity) InventoryRequestToEntity(reqInventory request.Inventory) entity.Inventory {
	var entityInventory entity.Inventory
	entityInventory.Product = parser.ProductRequestToEntity(reqInventory.Product)
	entityInventory.Date = primitive.NewDateTimeFromTime(reqInventory.Date)
	entityInventory.UpdateDate = primitive.NewDateTimeFromTime(time.Now())
	entityInventory.Quantity = reqInventory.Quantity

	return entityInventory
}
