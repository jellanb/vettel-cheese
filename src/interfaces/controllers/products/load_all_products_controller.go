package products

import (
	"fmt"
	"vettel-backend-app/src/application/usescases/usecase_products"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway/parser"
)

type LoadAllProductsControllerInterfaces interface {
	LoadAllProductsController() (error, *[]response.Product)
}

type LoadAllProductsController struct {
	loadAllProductsUseCaseInterfaces usecase_products.LoadAllProductsUseCaseInterfaces
}

func NewLoadAllProductsController(loadAllProductsUseCase usecase_products.LoadAllProductsUseCaseInterfaces) LoadAllProductsControllerInterfaces {
	return &LoadAllProductsController{loadAllProductsUseCase}
}

func (c LoadAllProductsController) LoadAllProductsController() (error, *[]response.Product) {
	fmt.Printf("Init process to load all products \n")
	err, allProducts := c.loadAllProductsUseCaseInterfaces.LoadAllProducts()
	if err != nil {
		fmt.Printf("Error in process to load al productss with error: %s \n", err.Error())
		return err, nil
	}
	productsList := parser.ProductsEntitiesToResponse(*allProducts)
	fmt.Printf("Finish process to load all products \n")
	return nil, &productsList
}
