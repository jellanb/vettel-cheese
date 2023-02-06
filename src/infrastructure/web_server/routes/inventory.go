package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"vettel-backend-app/src/application/usescases/usecase_inventory"
	"vettel-backend-app/src/infrastructure/database/inventory_repository"
	"vettel-backend-app/src/infrastructure/database/mongo_client"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/interfaces/controllers/inventory"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway/parser"
)

func LoadInventory(router *echo.Echo) *echo.Echo {
	router.GET("get-all-inventory", func(c echo.Context) error {
		mongoClient := mongo_client.NewMongoClientInterfaces()
		inventoryCollectionToEntity := parser.NewInventoryCollectionEntity()
		inventoryRepository := inventory_repository.NewInventoryRepositoryGateWay(mongoClient)
		inventoryEntityToCollectionInterfaces := parser.NewInventoryEntityToCollection()
		inventoryEntityToResponseInterfaces := parser.NewInventoryEntityToResponse()
		inventoryRepositoryGateWay := inventory_repository_gateway.NewInventoryRepositoryGateway(inventoryRepository, inventoryCollectionToEntity, inventoryEntityToCollectionInterfaces)
		loadInventoryUseCase := usecase_inventory.NewLoadInventoryUseCase(inventoryRepositoryGateWay)
		getAllInventoryController := inventory.NewGetAllInventoryController(inventoryRepositoryGateWay, loadInventoryUseCase, inventoryEntityToResponseInterfaces)

		err, inventoryResponse := getAllInventoryController.GetAllInventoryController(c.Request().Context())
		if err != nil {
			c.JSON(500, "ERROR_LOAD_INVENTORIES")
			return err
		}
		c.JSON(200, inventoryResponse)
		return nil
	})

	return router
}

func AddItemInventory(router *echo.Echo) *echo.Echo {
	router.POST("register-new-item-inventory", func(c echo.Context) error {
		var reqInventory request.Inventory
		mongoClient := mongo_client.NewMongoClientInterfaces()
		inventoryCollectionToEntity := parser.NewInventoryCollectionEntity()
		inventoryRepository := inventory_repository.NewInventoryRepositoryGateWay(mongoClient)
		inventoryRequestToEntityInterfaces := parser.NewInventoryRequestToEntity()
		inventoryEntityToCollectionInterfaces := parser.NewInventoryEntityToCollection()
		inventoryEntityToResponseInterfaces := parser.NewInventoryEntityToResponse()
		inventoryRepositoryGateWay := inventory_repository_gateway.NewInventoryRepositoryGateway(inventoryRepository, inventoryCollectionToEntity, inventoryEntityToCollectionInterfaces)
		addItemToInventoryUseCase := usecase_inventory.NewRegisterNewProductToInventoryUseCase(inventoryRepositoryGateWay)
		addItemToInventoryController := inventory.NewAddItemToInventoryController(addItemToInventoryUseCase, inventoryRequestToEntityInterfaces, inventoryEntityToResponseInterfaces)

		err := c.Bind(&reqInventory)
		if err != nil {
			fmt.Printf("Error binding request inventory with descryption : %s \n", err.Error())
			c.JSON(400, "INVALID_REQUEST")
			return err
		}

		err, inventoryResponse := addItemToInventoryController.AddItemToInventoryController(c.Request().Context(), reqInventory)
		if err != nil {
			fmt.Printf("Error inserting new inventory with product: %v with error %s", reqInventory.Product, err.Error())
			c.JSON(500, err.Error())
		}

		c.JSON(200, inventoryResponse)
		return nil
	})

	return router
}

func UpdateItemInventory(router *echo.Echo) *echo.Echo {
	router.POST("update-inventory-item", func(c echo.Context) error {
		var reqInventory request.Inventory
		mongoClient := mongo_client.NewMongoClientInterfaces()
		inventoryCollectionToEntity := parser.NewInventoryCollectionEntity()
		inventoryRepository := inventory_repository.NewInventoryRepositoryGateWay(mongoClient)
		inventoryRequestToEntityInterfaces := parser.NewInventoryRequestToEntity()
		inventoryEntityToCollectionInterfaces := parser.NewInventoryEntityToCollection()
		inventoryEntityToResponseInterfaces := parser.NewInventoryEntityToResponse()
		inventoryRepositoryGateWay := inventory_repository_gateway.NewInventoryRepositoryGateway(inventoryRepository, inventoryCollectionToEntity, inventoryEntityToCollectionInterfaces)
		updateInventoryUseCaseInterfaces := usecase_inventory.NewUpdateInventoryUseCase(inventoryRepositoryGateWay)
		updateInventoryController := inventory.NewUpdateInventoryController(updateInventoryUseCaseInterfaces, inventoryRequestToEntityInterfaces, inventoryEntityToResponseInterfaces)

		err := c.Bind(&reqInventory)
		if err != nil {
			c.JSON(400, "INVALID_REQUEST")
			fmt.Printf("error : %s \n", err.Error())
			return err
		}

		err, inventoryResponse := updateInventoryController.UpdateInventoryController(c.Request().Context(), reqInventory)
		if err != nil {
			fmt.Printf("Error updating inventory for barcode: %s with error: %s \n", reqInventory.Product.Barcode, err.Error())
			c.JSON(500, "INTERNAL_ERROR_SERVER")
		}

		c.JSON(200, inventoryResponse)
		return nil
	})
	return router
}

func DeleteInventoryItem(router *echo.Echo) *echo.Echo {
	router.DELETE("delete-inventory-item", func(c echo.Context) error {
		itemIDInventory := c.QueryParam("barcode")

		mongoClient := mongo_client.NewMongoClientInterfaces()
		inventoryCollectionToEntity := parser.NewInventoryCollectionEntity()
		inventoryRepository := inventory_repository.NewInventoryRepositoryGateWay(mongoClient)
		inventoryEntityToCollectionInterfaces := parser.NewInventoryEntityToCollection()
		inventoryRepositoryGateWay := inventory_repository_gateway.NewInventoryRepositoryGateway(inventoryRepository, inventoryCollectionToEntity, inventoryEntityToCollectionInterfaces)
		deleteItemInventoryUseCase := usecase_inventory.NewDeleteItemInventoryUseCase(inventoryRepositoryGateWay)
		deleteItemInventoryController := inventory.NewDeleteInventoryController(deleteItemInventoryUseCase)

		err, idDeleted := deleteItemInventoryController.DeleteInventoryController(c.Request().Context(), itemIDInventory)
		if err != nil {
			c.JSON(500, err.Error())
			return err
		}

		c.JSON(200, idDeleted)

		return nil
	})
	return router
}
