package usecase_inventory

import (
	"context"
	"fmt"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway"
)

type DeleteItemInventoryUseCaseInterfaces interface {
	DeleteItemInventoryUseCase(ctx context.Context, id string) (error, *string)
}

type DeleteItemInventoryUseCase struct {
	inventoryRepositoryGateWayInterfaces inventory_repository_gateway.InventoryRepositoryGatewayInterfaces
}

func NewDeleteItemInventoryUseCase(inventoryRepositoryGateWayInterfaces inventory_repository_gateway.InventoryRepositoryGatewayInterfaces) DeleteItemInventoryUseCaseInterfaces {
	return &DeleteItemInventoryUseCase{inventoryRepositoryGateWayInterfaces}
}

func (u DeleteItemInventoryUseCase) DeleteItemInventoryUseCase(ctx context.Context, barcode string) (error, *string) {
	fmt.Printf("Init delete inventory with barcode %s \n", barcode)
	err, idDeleted := u.inventoryRepositoryGateWayInterfaces.DeleteInventoryItem(ctx, barcode)
	if err != nil {
		fmt.Printf("Error try delete inventory with barcode: %s \n", barcode)
	}
	fmt.Printf("Finish delete inventory with barcode %s \n", barcode)
	return nil, idDeleted
}
