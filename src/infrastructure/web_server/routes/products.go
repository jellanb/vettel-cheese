package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"vettel-backend-app/src/application/usescases/usecase_products"
	"vettel-backend-app/src/infrastructure/database/inventory_repository"
	"vettel-backend-app/src/infrastructure/database/mongo_client"
	"vettel-backend-app/src/infrastructure/database/products_repository"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/interfaces/controllers/products"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway/parser"
	"vettel-backend-app/src/interfaces/geteways/products_repository_gateway"
)

func LoadAllProducts(router *echo.Echo) *echo.Echo {
	router.GET("load-all-products", func(c echo.Context) error {
		mongoClientInterfaces := mongo_client.NewMongoClientInterfaces()
		productsRepositoryInterfaces := products_repository.NewProductsRepository(mongoClientInterfaces)
		productsRepositoryGateGayInterfaces := products_repository_gateway.NewProductsRepositoryGateWay(productsRepositoryInterfaces)
		loadAllProductsUseCaseInterfaces := usecase_products.NewLoadAllProductsUseCase(productsRepositoryGateGayInterfaces)
		loadAllProductsController := products.NewLoadAllProductsController(loadAllProductsUseCaseInterfaces)

		err, allProducts := loadAllProductsController.LoadAllProductsController()
		if err != nil {
			c.JSON(500, "ERROR_LOADING_ALL_PRODUCTS")
			return err
		}

		c.JSON(200, allProducts)
		return nil
	})
	return router
}

func CreateProduct(router *echo.Echo) *echo.Echo {
	router.POST("create-product", func(c echo.Context) error {
		var newProductRequest request.Product
		mongoClientInterfaces := mongo_client.NewMongoClientInterfaces()
		productsRepositoryInterfaces := products_repository.NewProductsRepository(mongoClientInterfaces)
		productsRepositoryGateGayInterfaces := products_repository_gateway.NewProductsRepositoryGateWay(productsRepositoryInterfaces)
		createProductUseCaseInterfaces := usecase_products.NewCreateProductUseCase(productsRepositoryGateGayInterfaces)
		createProductControllerInterfaces := products.NewCreateProductController(createProductUseCaseInterfaces)

		err := c.Bind(&newProductRequest)
		if err != nil {
			fmt.Printf("Bad request error: %s", err.Error())
			c.JSON(401, "BAD_REQUEST")
			return err
		}

		err, newProduct := createProductControllerInterfaces.CreateProductController(c.Request().Context(), newProductRequest)
		if err != nil {
			c.JSON(500, err.Error())
			return err
		}
		c.JSON(200, newProduct)
		return nil
	})
	return router
}

func UpdateProduct(router *echo.Echo) *echo.Echo {
	router.PUT("update-product", func(c echo.Context) error {
		var newProductRequest request.Product
		mongoClientInterfaces := mongo_client.NewMongoClientInterfaces()
		productsRepositoryInterfaces := products_repository.NewProductsRepository(mongoClientInterfaces)
		inventoryRepositoryGateway := inventory_repository.NewInventoryRepositoryGateWay(mongoClientInterfaces)
		inventoryEntityToCollection := parser.NewInventoryEntityToCollection()
		inventoryCollectionToEntity := parser.NewInventoryCollectionEntity()
		productsRepositoryGatewayInterfaces := products_repository_gateway.NewProductsRepositoryGateWay(productsRepositoryInterfaces)
		inventoryRepositoryGatewayInterfaces := inventory_repository_gateway.NewInventoryRepositoryGateway(inventoryRepositoryGateway, inventoryCollectionToEntity, inventoryEntityToCollection)
		updateProductUseCaseInterfaces := usecase_products.NewUpdateProductUseCase(productsRepositoryGatewayInterfaces, inventoryRepositoryGatewayInterfaces)
		updateProductControllerInterfaces := products.NewUpdateProductController(updateProductUseCaseInterfaces)

		err := c.Bind(&newProductRequest)
		if err != nil {
			fmt.Printf("Error biding product request with error: %s", err.Error())
			c.JSON(401, "BAD_REQUEST")
			return err
		}

		err, productResponse := updateProductControllerInterfaces.UpdateProductController(c.Request().Context(), newProductRequest)
		if err != nil {
			c.JSON(500, "INTERNAL_ERROR_SERVER")
			return err
		}
		c.JSON(200, productResponse)
		return nil
	})
	return router
}

func DeleteProduct(router *echo.Echo) *echo.Echo {
	router.DELETE("delete-product", func(c echo.Context) error {
		barcode := c.QueryParam("barcode")
		mongoClientInterfaces := mongo_client.NewMongoClientInterfaces()
		productsRepositoryInterfaces := products_repository.NewProductsRepository(mongoClientInterfaces)
		productsRepositoryGateGayInterfaces := products_repository_gateway.NewProductsRepositoryGateWay(productsRepositoryInterfaces)
		deleteProductUseCaseInterfaces := usecase_products.NewDeleteProductUseCase(productsRepositoryGateGayInterfaces)
		deleteProductControllerInterfaces := products.NewDeleteProductController(deleteProductUseCaseInterfaces)

		if barcode == "" {
			c.JSON(400, "BAD_REQUEST")
		}

		err, IdDeleted := deleteProductControllerInterfaces.DeleteProductController(c.Request().Context(), barcode)
		if err != nil {
			c.JSON(500, "INTERNAL_ERROR_SERVER")
			return err
		}
		c.JSON(200, IdDeleted)
		return nil
	})
	return router
}
