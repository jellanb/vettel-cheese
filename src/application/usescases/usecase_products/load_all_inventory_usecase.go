package usecase_products

import (
	"fmt"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway"
)

type LoadAllProductsUseCaseInterfaces interface {
	LoadAllProducts() (error, *[]entity.Product)
}

type LoadAllProductsUseCase struct {
	productsRepositoryGateWayInterfaces products_repository_gateway.ProductsRepositoryGateWayInterfaces
}

func NewLoadAllProductsUseCase(productsRepositoryGateWayInterfaces products_repository_gateway.ProductsRepositoryGateWayInterfaces) LoadAllProductsUseCaseInterfaces {
	return &LoadAllProductsUseCase{productsRepositoryGateWayInterfaces}
}

func (u LoadAllProductsUseCase) LoadAllProducts() (error, *[]entity.Product) {
	fmt.Printf("Loading all products \n")
	err, allProductsEntities := u.productsRepositoryGateWayInterfaces.LoadAllProducts()
	if err != nil {
		fmt.Printf("Error loading all products with error: %s \n", err.Error())
		return err, nil
	}
	fmt.Printf("All products loaded \n")
	return nil, allProductsEntities
}
