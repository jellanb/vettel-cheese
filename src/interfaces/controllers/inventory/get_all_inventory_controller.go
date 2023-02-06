package inventory

import (
	"context"
	"fmt"
	"vettel-backend-app/src/application/usescases/usecase_inventory"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway/parser"
)

type GetAllInventoryControllerInterfaces interface {
	GetAllInventoryController(ctx context.Context) (error, []response.Inventory)
}

type GetAllInventoryController struct {
	inventoryRepositoryGateWayInterfaces inventory_repository_gateway.InventoryRepositoryGatewayInterfaces
	loadInventoryUseCaseInterfaces       usecase_inventory.LoadInventoryUseCaseInterfaces
	inventoryEntityToResponseInterfaces  parser.InventoryEntityToResponseInterfaces
}

func NewGetAllInventoryController(
	inventoryRepositoryGateWayInterfaces inventory_repository_gateway.InventoryRepositoryGatewayInterfaces,
	loadInventoryUseCaseInterfaces usecase_inventory.LoadInventoryUseCaseInterfaces,
	inventoryEntityToResponseInterfaces parser.InventoryEntityToResponseInterfaces,
) GetAllInventoryControllerInterfaces {
	return &GetAllInventoryController{
		inventoryRepositoryGateWayInterfaces,
		loadInventoryUseCaseInterfaces,
		inventoryEntityToResponseInterfaces,
	}
}

func (g GetAllInventoryController) GetAllInventoryController(ctx context.Context) (error, []response.Inventory) {
	fmt.Printf("Init process to load all inventory \n")
	err, allInventory := g.loadInventoryUseCaseInterfaces.LoadInventory(ctx)
	if err != nil {
		fmt.Printf("Error try to load all inventory with error: %s \n", err.Error())
		return err, nil
	}
	inventoryResponse := g.inventoryEntityToResponseInterfaces.SliceInventoryEntityToResponse(allInventory)
	fmt.Printf("Finish process to load all inventory\n")
	return nil, inventoryResponse
}
