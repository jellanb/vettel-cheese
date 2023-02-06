package usecase_products

import (
	"context"
	"fmt"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway"
)

type DeleteProductUseCaseInterfaces interface {
	DeleteProductUseCase(ctx context.Context, barcode string) (error, *string)
}

type DeleteProductUseCase struct {
	productRepositoryGateWayInterfaces products_repository_gateway.ProductsRepositoryGateWayInterfaces
}

func NewDeleteProductUseCase(productRepositoryGateWayInterfaces products_repository_gateway.ProductsRepositoryGateWayInterfaces) DeleteProductUseCaseInterfaces {
	return &DeleteProductUseCase{productRepositoryGateWayInterfaces}
}

func (u DeleteProductUseCase) DeleteProductUseCase(ctx context.Context, barcode string) (error, *string) {
	fmt.Printf("Deleting product by barcode: %s \n", barcode)
	err, IdDeleted := u.productRepositoryGateWayInterfaces.DeleteProductById(ctx, barcode)
	if err != nil {
		fmt.Printf("Error deleting product with barcode: %d", barcode)
		return err, nil
	}
	fmt.Printf("Product deleted with barcode: %s \n", barcode)
	return nil, IdDeleted
}
