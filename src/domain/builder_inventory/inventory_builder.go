package builder_inventory

import "vettel-backend-app/src/domain/entity"

type InventoryBuilderInterfaces interface {
}

type InventoryBuilder struct {
	inventory entity.Inventory
}

func NewInventoryBuilder() InventoryBuilderInterfaces {
	return &InventoryBuilder{}
}

func (i *InventoryBuilder) BuildNewInventory() entity.Inventory {
	return i.inventory
}
