package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"vettel-backend-app/src/application/usescases/usecase_orders"
	"vettel-backend-app/src/infrastructure/database/inventory_repository"
	"vettel-backend-app/src/infrastructure/database/mongo_client"
	"vettel-backend-app/src/infrastructure/database/order_repository"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/interfaces/controllers/dispatch"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway/parser"
	"vettel-backend-app/src/interfaces/geteways/order_repository_gateway"
)

func LoadAllDispatch(router *echo.Echo) *echo.Echo {

	return router
}

func AddDispatch(router *echo.Echo) *echo.Echo {
	router.POST("register-new-dispatch", func(c echo.Context) error {
		var dispatchRequest request.Dispatch
		mongoClient := mongo_client.NewMongoClientInterfaces()
		orderRepositoryInterfaces := order_repository.NewOrderRepository(mongoClient)
		orderRepositoryGateWayInterfaces := order_repository_gateway.NewOrderRepositoryGateWay(orderRepositoryInterfaces)
		inventoryRepository := inventory_repository.NewInventoryRepositoryGateWay(mongoClient)
		inventoryCollectionToEntityInterfaces := parser.NewInventoryCollectionEntity()
		inventoryEntityToCollectionInterfaces := parser.NewInventoryEntityToCollection()
		inventoryRepositoryGateWayInterfaces := inventory_repository_gateway.NewInventoryRepositoryGateway(inventoryRepository, inventoryCollectionToEntityInterfaces, inventoryEntityToCollectionInterfaces)
		registerNewOrderUseCaseInterfaces := usecase_orders.NewRegisterNewOrderUseCase(orderRepositoryGateWayInterfaces, inventoryRepositoryGateWayInterfaces)
		addDispatchControllerInterfaces := dispatch.NewAddDispatchController(registerNewOrderUseCaseInterfaces)

		err := c.Bind(&dispatchRequest)
		if err != nil {
			c.JSON(400, "INVALID_REQUEST")
			fmt.Printf("error : %s \n", err.Error())
			return err
		}

		err, orderCreated := addDispatchControllerInterfaces.AddDispatchController(c.Request().Context(), dispatchRequest)
		if err != nil {
			c.JSON(500, "INTERNAL_ERROR_SERVER")
			return err
		}
		c.JSON(200, orderCreated)
		return nil
	})
	return router
}
