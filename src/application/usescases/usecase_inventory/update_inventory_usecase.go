package usecase_inventory

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway"
)

type UpdateInventoryUseCaseInterfaces interface {
	UpdateInventoryUseCase(ctx context.Context, entityInventory entity.Inventory) (error, *entity.Inventory)
}

type UpdateInventoryUseCase struct {
	inventoryRepositoryGatewayInterfaces inventory_repository_gateway.InventoryRepositoryGatewayInterfaces
}

func NewUpdateInventoryUseCase(inventoryRepositoryGatewayInterfaces inventory_repository_gateway.InventoryRepositoryGatewayInterfaces) UpdateInventoryUseCaseInterfaces {
	return &UpdateInventoryUseCase{inventoryRepositoryGatewayInterfaces}
}

func (u UpdateInventoryUseCase) UpdateInventoryUseCase(ctx context.Context, entityInventory entity.Inventory) (error, *entity.Inventory) {
	fmt.Printf("Validating if iventory item exist for item barcode: %s \n", entityInventory.Product.Barcode)
	err, inventoryResult := u.inventoryRepositoryGatewayInterfaces.FindInventoryItemByBarcode(ctx, entityInventory.Product.Barcode)
	if err != nil {
		return err, nil
	}
	if inventoryResult == nil {
		fmt.Printf("Error inventory item not exist with barcode: %s \n", entityInventory.Product.Barcode)
		return errors.New("Error updating Item already exist."), nil
	}
	fmt.Printf("Validation ok item exist to update with barcode: %s \n", entityInventory.Product.Barcode)
	fmt.Printf("Init update inventory with product %v and quantity: %d \n", entityInventory.Product, entityInventory.Quantity)
	err, inventoryResult = u.inventoryRepositoryGatewayInterfaces.UpdateInventoryItem(ctx, entityInventory)
	if err != nil {
		fmt.Printf("Error updating inventory with product: %v and quantity: %d: with error: %s \n", entityInventory.Product, entityInventory.Quantity, err.Error())
		return err, nil
	}
	fmt.Printf("Finish update inventory with product %v and quantity: %d \n", entityInventory.Product, entityInventory.Quantity)
	return nil, inventoryResult
}
