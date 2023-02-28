package usecase_inventory

import (
	"context"
	"fmt"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway"

	"github.com/pkg/errors"
)

type RegisterNewProductToInventoryUseCaseInterfaces interface {
	RegisterNewProductToInventory(ctx context.Context, newItem entity.Inventory) (error, *entity.Inventory)
}

type RegisterNewProductToInventoryUseCase struct {
	inventoryRepositoryGatewayInterfaces inventory_repository_gateway.InventoryRepositoryGatewayInterfaces
}

func NewRegisterNewProductToInventoryUseCase(inventoryRepositoryGatewayInterfaces inventory_repository_gateway.InventoryRepositoryGatewayInterfaces) RegisterNewProductToInventoryUseCaseInterfaces {
	return &RegisterNewProductToInventoryUseCase{inventoryRepositoryGatewayInterfaces}
}

func (r RegisterNewProductToInventoryUseCase) RegisterNewProductToInventory(ctx context.Context, newItem entity.Inventory) (error, *entity.Inventory) {
	fmt.Printf("Validating if iventory item exist for item barcode: %s \n", newItem.Product.Barcode)
	err, inventoryResult := r.inventoryRepositoryGatewayInterfaces.FindInventoryItemByBarcode(ctx, newItem.Product.Barcode)
	if err != nil {
		return err, nil
	}
	if inventoryResult != nil {
		fmt.Printf("Error inventory item already exist with barcode: %s \n", newItem.Product.Barcode)
		return errors.New("INVENTORY_ALREADY_EXIST."), nil
	}
	fmt.Printf("Validation ok item exist to update with barcode: %s \n", newItem.Product.Barcode)

	fmt.Printf("Adding item to inventory with product: %v and quantity: %d: \n", newItem.Product, newItem.Quantity)
	err, newInventory := r.inventoryRepositoryGatewayInterfaces.AddInventoryItem(ctx, newItem)
	if err != nil {
		fmt.Printf("Error try add item to inventory with product: %v and quantity: %d: with error: %s \n", newItem.Product, newItem.Quantity, err.Error())
		return err, nil
	}
	fmt.Printf("Item added to inventory with product: %v and quantity: %d: \n", newItem.Product, newItem.Quantity)

	return nil, newInventory
}
