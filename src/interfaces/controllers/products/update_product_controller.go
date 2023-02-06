package products

import (
	"context"
	"fmt"
	"vettel-backend-app/src/application/usescases/usecase_products"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway/parser"
)

type UpdateProductControllerInterfaces interface {
	UpdateProductController(ctx context.Context, product request.Product) (error, *response.Product)
}

type UpdateProductController struct {
	updateProductUseCaseInterfaces usecase_products.UpdateProductUseCaseInterfaces
}

func NewUpdateProductController(updateProductUseCaseInterfaces usecase_products.UpdateProductUseCaseInterfaces) UpdateProductControllerInterfaces {
	return &UpdateProductController{updateProductUseCaseInterfaces}
}

func (c UpdateProductController) UpdateProductController(ctx context.Context, product request.Product) (error, *response.Product) {
	fmt.Printf("Init process to update product \n")
	entityProduct := parser.ProductRequestToEntity(product)
	err, productUpdated := c.updateProductUseCaseInterfaces.UpdateProductUseCase(ctx, entityProduct)
	if err != nil {
		fmt.Printf("Error processing product update with error: %s  \n", err.Error())
		return err, nil
	}
	productResp := parser.ProductEntityToResponse(productUpdated)
	fmt.Printf("Finish process to update product \n")
	return nil, productResp
}
