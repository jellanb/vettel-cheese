package products_repository_gateway

import (
	"context"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/database/products_repository"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway/parser"
)

type ProductsRepositoryGateWayInterfaces interface {
	LoadAllProducts() (error, *[]entity.Product)
	FindProductByBarcode(barcode string) (error, *entity.Product)
	CreateProduct(ctx context.Context, newProduct entity.Product) (error, *entity.Product)
	UpdateProduct(ctx context.Context, entityProduct entity.Product) (error, *entity.Product)
	DeleteProductById(ctx context.Context, barcode string) (error, *string)
}

type ProductsRepositoryGateWay struct {
	productsRepositoryInterfaces products_repository.ProductsRepositoryInterfaces
}

func NewProductsRepositoryGateWay(productsRepositoryInterfaces products_repository.ProductsRepositoryInterfaces) ProductsRepositoryGateWayInterfaces {
	return &ProductsRepositoryGateWay{productsRepositoryInterfaces}
}

func (g ProductsRepositoryGateWay) LoadAllProducts() (error, *[]entity.Product) {
	err, productsList := g.productsRepositoryInterfaces.FindAllProducts()
	if err != nil {
		return err, nil
	}
	productsEntity := parser.ProductsCollectionToEntity(productsList)
	return nil, &productsEntity
}

func (g ProductsRepositoryGateWay) FindProductByBarcode(barcode string) (error, *entity.Product) {
	err, productCreatedCollection := g.productsRepositoryInterfaces.FindProductByBarcode(barcode)
	if err != nil {
		return err, nil
	}
	if productCreatedCollection == nil {
		return nil, nil
	}
	productEntityCreated := parser.ProductCollectionToEntity(*productCreatedCollection)
	return nil, &productEntityCreated
}

func (g ProductsRepositoryGateWay) CreateProduct(ctx context.Context, newProduct entity.Product) (error, *entity.Product) {
	productCollection := parser.ProductEntityToCollection(newProduct)
	err, productCreatedCollection := g.productsRepositoryInterfaces.InsertOneProduct(ctx, productCollection)
	if err != nil {
		return err, nil
	}
	productEntityCreated := parser.ProductCollectionToEntity(*productCreatedCollection)
	return nil, &productEntityCreated
}

func (g ProductsRepositoryGateWay) UpdateProduct(ctx context.Context, entityProduct entity.Product) (error, *entity.Product) {
	productCollection := parser.ProductEntityToCollection(entityProduct)
	err, product := g.productsRepositoryInterfaces.UpdateOneProductById(ctx, productCollection)
	if err != nil {
		return err, nil
	}
	productEntityCreated := parser.ProductCollectionToEntity(*product)
	return nil, &productEntityCreated
}

func (g ProductsRepositoryGateWay) DeleteProductById(ctx context.Context, barcode string) (error, *string) {
	err, IdDeleted := g.productsRepositoryInterfaces.DeleteProductById(ctx, barcode)
	if err != nil {
		return err, nil
	}
	return nil, IdDeleted
}
