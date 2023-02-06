package usecase_products

import (
	"context"
	"fmt"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway"
)

type UpdateProductUseCaseInterfaces interface {
	UpdateProductUseCase(ctx context.Context, product entity.Product) (error, *entity.Product)
}

type UpdateProductUseCase struct {
	productsRepositoryGateWayInterfaces  products_repository_gateway.ProductsRepositoryGateWayInterfaces
	inventoryRepositoryGatewayInterfaces inventory_repository_gateway.InventoryRepositoryGatewayInterfaces
}

func NewUpdateProductUseCase(
	productsRepositoryGateWayInterfaces products_repository_gateway.ProductsRepositoryGateWayInterfaces,
	inventoryRepositoryGatewayInterfaces inventory_repository_gateway.InventoryRepositoryGatewayInterfaces,
) UpdateProductUseCaseInterfaces {
	return &UpdateProductUseCase{
		productsRepositoryGateWayInterfaces,
		inventoryRepositoryGatewayInterfaces,
	}
}

func (u UpdateProductUseCase) UpdateProductUseCase(ctx context.Context, product entity.Product) (error, *entity.Product) {
	fmt.Printf("Find products in inventory to barcode: %s \n", product.Barcode)

	fmt.Printf("Updating product with barcode: %s, despcription: %s and price: %d  \n", product.Barcode, product.Description, product.SaleAmount)
	err, inventoryProduct := u.inventoryRepositoryGatewayInterfaces.FindInventoryItemByBarcode(ctx, product.Barcode)
	if err != nil {
		fmt.Printf("Error finding product in invetory for barcode: %s  \n", product.Barcode)
		return err, nil
	}
	if inventoryProduct == nil {
		fmt.Printf("Noting product found in invetory for barcode: %s  \n", product.Barcode)
		return err, nil
	}
	fmt.Printf("Updating product and invetory for barcode: %s  \n", product.Barcode)

	var inventoryUpdate *entity.Inventory
	inventoryUpdate = inventoryProduct
	inventoryUpdate.SaleAmount = product.SaleAmount
	inventoryUpdate.Product.Description = product.Description
	inventoryUpdate.Product.SaleAmount = product.SaleAmount

	err, _ = u.inventoryRepositoryGatewayInterfaces.UpdateInventoryItem(ctx, *inventoryUpdate)
	if err != nil {
		fmt.Printf("Error updating product and invetory for barcode: %s  \n", product.Barcode)
		return err, nil
	}
	err, productUpdated := u.productsRepositoryGateWayInterfaces.UpdateProduct(ctx, product)
	if err != nil {
		return err, nil
	}
	fmt.Printf("Product and inventory updated with barcode: %s, despcription: %s and price: %d  \n", product.Barcode, product.Description, product.SaleAmount)
	return nil, productUpdated
}
