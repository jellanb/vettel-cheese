package products

import (
	"context"
	"fmt"
	"vettel-backend-app/src/application/usescases/usecase_products"
)

type DeleteProductControllerInterfaces interface {
	DeleteProductController(ctx context.Context, barcode string) (error, *string)
}

type DeleteProductController struct {
	deleteProductUseCaseInterfaces usecase_products.DeleteProductUseCaseInterfaces
}

func NewDeleteProductController(deleteProductUseCaseInterfaces usecase_products.DeleteProductUseCaseInterfaces) DeleteProductControllerInterfaces {
	return &DeleteProductController{deleteProductUseCaseInterfaces}
}

func (c DeleteProductController) DeleteProductController(ctx context.Context, barcode string) (error, *string) {
	fmt.Printf("Init process to delete product with barcode: %s  \n", barcode)
	err, idDeleted := c.deleteProductUseCaseInterfaces.DeleteProductUseCase(ctx, barcode)
	if err != nil {
		fmt.Printf("Error in process to delete product with barcode: %s  \n", barcode)
		return err, nil
	}
	return nil, idDeleted
}
