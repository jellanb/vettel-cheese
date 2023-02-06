package usecase_products

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway"
)

type CreateProductUseCaseInterfaces interface {
	CreateProductUseCase(ctx context.Context, newProduct entity.Product) (error, *entity.Product)
}

type CreateProductUseCase struct {
	productsRepositoryGateWayInterfaces products_repository_gateway.ProductsRepositoryGateWayInterfaces
}

func NewCreateProductUseCase(productsRepositoryGateWayInterfaces products_repository_gateway.ProductsRepositoryGateWayInterfaces) CreateProductUseCaseInterfaces {
	return &CreateProductUseCase{productsRepositoryGateWayInterfaces}
}

func (u CreateProductUseCase) CreateProductUseCase(ctx context.Context, newProduct entity.Product) (error, *entity.Product) {
	fmt.Printf("Init validation to check if product already exists with barcode: %s \n", newProduct.Barcode)
	err, productFound := u.productsRepositoryGateWayInterfaces.FindProductByBarcode(newProduct.Barcode)
	if err != nil {
		return err, nil
	}
	if productFound != nil {
		fmt.Printf("Error creating new product, the product with barcode: %s already axist. \n", newProduct.Barcode)
		return errors.New("PRODUCT_ALREADY_EXIST."), nil
	}
	fmt.Printf("Finish validation ok, product ready to create with barcode: %s \n", newProduct.Barcode)

	fmt.Printf("Creating product with barcode: %s, description: %s and price: %d \n", newProduct.Barcode, newProduct.Description, newProduct.SaleAmount)
	err, productEntityResult := u.productsRepositoryGateWayInterfaces.CreateProduct(ctx, newProduct)
	if err != nil {
		fmt.Printf("Error creating product with barcode: %s, description: %s and price %d \n", newProduct.Barcode, newProduct.Description, newProduct.SaleAmount)
		return err, nil
	}
	fmt.Printf("Product created with barcode: %s, description: %s and price %d \n", newProduct.Barcode, newProduct.Description, newProduct.SaleAmount)
	return nil, productEntityResult
}
