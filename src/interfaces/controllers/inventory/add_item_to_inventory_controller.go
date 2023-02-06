package inventory

import (
	"context"
	"fmt"
	"vettel-backend-app/src/application/usescases/usecase_inventory"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway/parser"
)

type AddItemToInventoryControllerInterfaces interface {
	AddItemToInventoryController(ctx context.Context, reqInventory request.Inventory) (error, *response.Inventory)
}

type AddItemToInventoryController struct {
	registerNewProductToInventoryUseCaseInterfaces usecase_inventory.RegisterNewProductToInventoryUseCaseInterfaces
	inventoryRequestToEntityInterfaces             parser.InventoryRequestToEntityInterfaces
	inventoryEntityToResponseInterfaces            parser.InventoryEntityToResponseInterfaces
}

func NewAddItemToInventoryController(
	registerNewProductToInventoryUseCaseInterfaces usecase_inventory.RegisterNewProductToInventoryUseCaseInterfaces,
	inventoryRequestToEntityInterfaces parser.InventoryRequestToEntityInterfaces,
	inventoryEntityToResponseInterfaces parser.InventoryEntityToResponseInterfaces,
) AddItemToInventoryControllerInterfaces {
	return &AddItemToInventoryController{
		registerNewProductToInventoryUseCaseInterfaces,
		inventoryRequestToEntityInterfaces,
		inventoryEntityToResponseInterfaces,
	}
}

func (c AddItemToInventoryController) AddItemToInventoryController(ctx context.Context, reqInventory request.Inventory) (error, *response.Inventory) {
	fmt.Printf("Init process to add item to inventory with product: %v and quantity: %d: \n", reqInventory.Product, reqInventory.Quantity)
	inventory := c.inventoryRequestToEntityInterfaces.InventoryRequestToEntity(reqInventory)

	err, inventoryResult := c.registerNewProductToInventoryUseCaseInterfaces.RegisterNewProductToInventory(ctx, inventory)
	if err != nil {
		fmt.Printf("Error in process for add new item to inventory %s \n", err.Error())
		return err, nil
	}
	inventoryRes := c.inventoryEntityToResponseInterfaces.InventoryEntityToResponse(inventoryResult)
	fmt.Printf("Finish process to add item to inventory with product: %v and quantity: %d: \n", reqInventory.Product, reqInventory.Quantity)
	return nil, &inventoryRes
}
