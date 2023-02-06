package usecase_inventory

import (
	"context"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway"
)

type LoadInventoryUseCaseInterfaces interface {
	LoadInventory(ctx context.Context) (error, []*entity.Inventory)
}

type LoadInventoryUseCase struct {
	inventoryRepositoryGateway inventory_repository_gateway.InventoryRepositoryGatewayInterfaces
}

func NewLoadInventoryUseCase(inventoryRepositoryGateway inventory_repository_gateway.InventoryRepositoryGatewayInterfaces) LoadInventoryUseCaseInterfaces {
	return &LoadInventoryUseCase{inventoryRepositoryGateway}
}

func (l LoadInventoryUseCase) LoadInventory(ctx context.Context) (error, []*entity.Inventory) {
	err, allInventory := l.inventoryRepositoryGateway.GetAll(ctx)
	if err != nil {
		return err, nil
	}
	return nil, allInventory
}
