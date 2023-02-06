package inventory_repository_gateway

import (
	"context"
	"fmt"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/database/inventory_repository"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway/parser"
)

type InventoryRepositoryGatewayInterfaces interface {
	GetAll(ctx context.Context) (error, []*entity.Inventory)
	FindInventoryItemByBarcode(ctx context.Context, barCode string) (error, *entity.Inventory)
	AddInventoryItem(ctx context.Context, newInventory entity.Inventory) (error, *entity.Inventory)
	UpdateInventoryItem(ctx context.Context, newInventory entity.Inventory) (error, *entity.Inventory)
	DeleteInventoryItem(ctx context.Context, id string) (error, *string)
}

type InventoryRepositoryGateway struct {
	inventoryRepositoryInterfaces         inventory_repository.InventoryRepositoryInterfaces
	inventoryCollectionToEntityInterfaces parser.InventoryCollectionToEntityInterfaces
	inventoryEntityToCollectionInterfaces parser.InventoryEntityToCollectionInterfaces
}

func NewInventoryRepositoryGateway(
	inventoryRepository inventory_repository.InventoryRepositoryInterfaces,
	inventoryCollectionToEntityInterfaces parser.InventoryCollectionToEntityInterfaces,
	inventoryEntityToCollectionInterfaces parser.InventoryEntityToCollectionInterfaces,
) InventoryRepositoryGatewayInterfaces {
	return &InventoryRepositoryGateway{
		inventoryRepository,
		inventoryCollectionToEntityInterfaces,
		inventoryEntityToCollectionInterfaces,
	}
}

func (i InventoryRepositoryGateway) FindInventoryItemByBarcode(ctx context.Context, barCode string) (error, *entity.Inventory) {
	err, inventoryItemColl := i.inventoryRepositoryInterfaces.FindInventoryItemByBarCode(ctx, barCode)
	if err != nil {
		return err, nil
	}
	if inventoryItemColl == nil {
		return nil, nil
	}
	entityInventory := i.inventoryCollectionToEntityInterfaces.InventoryCollectionToEntity(inventoryItemColl)
	return nil, entityInventory
}

func (i InventoryRepositoryGateway) GetAll(ctx context.Context) (error, []*entity.Inventory) {
	err, inventoryColl := i.inventoryRepositoryInterfaces.GetAll(ctx)
	if err != nil {
		return err, nil
	}
	entityInventory := i.inventoryCollectionToEntityInterfaces.InventoryCollectionsToEntities(inventoryColl)

	return nil, entityInventory
}

func (i InventoryRepositoryGateway) AddInventoryItem(ctx context.Context, newInventory entity.Inventory) (error, *entity.Inventory) {
	inventoryColl := i.inventoryEntityToCollectionInterfaces.InventoryEntityToCollection(newInventory)
	err, inventoryColl := i.inventoryRepositoryInterfaces.AddInventoryItem(ctx, *inventoryColl)
	if err != nil {
		return err, nil
	}
	return nil, &newInventory
}

func (i InventoryRepositoryGateway) UpdateInventoryItem(ctx context.Context, newInventory entity.Inventory) (error, *entity.Inventory) {
	inventoryColl := i.inventoryEntityToCollectionInterfaces.InventoryEntityToCollection(newInventory)
	fmt.Printf("after convert: %v  \n", inventoryColl)
	err, inventoryColl := i.inventoryRepositoryInterfaces.UpdateInventory(ctx, *inventoryColl)
	if err != nil {
		return err, nil
	}
	return nil, &newInventory
}

func (i InventoryRepositoryGateway) DeleteInventoryItem(ctx context.Context, barcode string) (error, *string) {
	err, idResult := i.inventoryRepositoryInterfaces.DeleteInventory(ctx, barcode)
	if err != nil {
		return err, nil
	}
	return nil, idResult
}
