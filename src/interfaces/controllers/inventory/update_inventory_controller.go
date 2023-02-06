package inventory

import (
	"context"
	"fmt"
	"vettel-backend-app/src/application/usescases/usecase_inventory"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway/parser"
)

type UpdateInventoryControllerInterfaces interface {
	UpdateInventoryController(ctx context.Context, reqUpdateInv request.Inventory) (error, *response.Inventory)
}

type UpdateInventoryController struct {
	UpdateInventoryUseCaseInterfaces    usecase_inventory.UpdateInventoryUseCaseInterfaces
	inventoryRequestToEntityInterfaces  parser.InventoryRequestToEntityInterfaces
	inventoryEntityToResponseInterfaces parser.InventoryEntityToResponseInterfaces
}

func NewUpdateInventoryController(
	UpdateInventoryUseCaseInterfaces usecase_inventory.UpdateInventoryUseCaseInterfaces,
	inventoryRequestToEntityInterfaces parser.InventoryRequestToEntityInterfaces,
	inventoryEntityToResponseInterfaces parser.InventoryEntityToResponseInterfaces,
) UpdateInventoryControllerInterfaces {
	return &UpdateInventoryController{
		UpdateInventoryUseCaseInterfaces,
		inventoryRequestToEntityInterfaces,
		inventoryEntityToResponseInterfaces,
	}
}

func (c UpdateInventoryController) UpdateInventoryController(ctx context.Context, reqUpdateInv request.Inventory) (error, *response.Inventory) {
	fmt.Printf("Init process to add item to inventory with product: %v and quantity: %d: \n", reqUpdateInv.Product, reqUpdateInv.Quantity)
	inventoryEntity := c.inventoryRequestToEntityInterfaces.InventoryRequestToEntity(reqUpdateInv)
	err, inventoryResult := c.UpdateInventoryUseCaseInterfaces.UpdateInventoryUseCase(ctx, inventoryEntity)
	if err != nil {
		fmt.Printf("Error in process for update inventory %s \n", err.Error())
		return err, nil
	}
	inventoryResp := c.inventoryEntityToResponseInterfaces.InventoryEntityToResponse(inventoryResult)
	fmt.Printf("Finish process to add item to inventory with product: %v and quantity: %d: \n", reqUpdateInv.Product, reqUpdateInv.Quantity)
	return nil, &inventoryResp
}
