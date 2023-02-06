package inventory

import (
	"context"
	"fmt"
	"vettel-backend-app/src/application/usescases/usecase_inventory"
)

type DeleteInventoryControllerInterfaces interface {
	DeleteInventoryController(ctx context.Context, id string) (error, *string)
}

type DeleteInventoryController struct {
	deleteInventoryControllerUseCaseInterfaces usecase_inventory.DeleteItemInventoryUseCaseInterfaces
}

func NewDeleteInventoryController(deleteInventoryControllerUseCaseInterfaces usecase_inventory.DeleteItemInventoryUseCaseInterfaces) DeleteInventoryControllerInterfaces {
	return &DeleteInventoryController{deleteInventoryControllerUseCaseInterfaces}
}

func (c DeleteInventoryController) DeleteInventoryController(ctx context.Context, barcode string) (error, *string) {
	fmt.Printf("Init process to delete item inventory with barcode:%s \n", barcode)
	err, idDeleted := c.deleteInventoryControllerUseCaseInterfaces.DeleteItemInventoryUseCase(ctx, barcode)
	if err != nil {
		fmt.Printf("Error in process for delete inventory item with barcode: %s \n", barcode)
		return err, nil
	}
	fmt.Printf("Finish process to delete item inventory with barcode:%s \n", barcode)
	return nil, idDeleted
}
