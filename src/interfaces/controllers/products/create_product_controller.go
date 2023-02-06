package products

import (
	"context"
	"fmt"
	"vettel-backend-app/src/application/usescases/usecase_products"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway/parser"
)

type CreateProductControllerInterfaces interface {
	CreateProductController(ctx context.Context, newProduct request.Product) (error, *response.Product)
}

type CreateProductController struct {
	useCaseCreateProductUseCaseInterfaces usecase_products.CreateProductUseCaseInterfaces
}

func NewCreateProductController(useCaseCreateProductUseCaseInterfaces usecase_products.CreateProductUseCaseInterfaces) CreateProductControllerInterfaces {
	return &CreateProductController{useCaseCreateProductUseCaseInterfaces}
}

func (c CreateProductController) CreateProductController(ctx context.Context, newProduct request.Product) (error, *response.Product) {
	fmt.Printf("Init process to create product with barcode %s, description: %s and price: %d \n", newProduct.Barcode, newProduct.Description, newProduct.SaleAmount)
	newProductEntity := parser.ProductRequestToEntity(newProduct)
	err, productCreatedEntity := c.useCaseCreateProductUseCaseInterfaces.CreateProductUseCase(ctx, newProductEntity)
	if err != nil {
		return err, nil
	}
	productCreatedResp := parser.ProductEntityToResponse(productCreatedEntity)
	fmt.Printf("Finish process to create product with barcode %s, description: %s and price: %d  \n", newProduct.Barcode, newProduct.Description, newProduct.SaleAmount)
	return nil, productCreatedResp

}
