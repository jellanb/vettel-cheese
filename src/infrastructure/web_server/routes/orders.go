package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
	"vettel-backend-app/src/application/usescases/usecase_orders"
	"vettel-backend-app/src/infrastructure/database/mongo_client"
	"vettel-backend-app/src/infrastructure/database/order_repository"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/interfaces/controllers/order"
	"vettel-backend-app/src/interfaces/geteways/order_repository_gateway"
)

func ListOrdersMonth(route *echo.Echo) *echo.Echo {
	route.GET("/orders", func(c echo.Context) error {
		mongoClientInterfaces := mongo_client.NewMongoClientInterfaces()
		orderRepositoryInterfaces := order_repository.NewOrderRepository(mongoClientInterfaces)
		orderRepositoryGateWayInterfaces := order_repository_gateway.NewOrderRepositoryGateWay(orderRepositoryInterfaces)
		listOrderByMonthController := order.NewListOrderByMonthController(orderRepositoryGateWayInterfaces)

		err, orderList := listOrderByMonthController.ListOrderByMonthController()
		if err != nil {
			c.JSON(500, "ERROR_GETTING_ORDERS")
			return err
		}

		c.JSON(200, orderList)
		return nil
	})

	return route
}

func EditOrder(router *echo.Echo) *echo.Echo {
	router.PUT("/update-order", func(c echo.Context) error {
		var reqOrder request.Order
		mongoClientInterfaces := mongo_client.NewMongoClientInterfaces()
		orderRepositoryInterfaces := order_repository.NewOrderRepository(mongoClientInterfaces)
		orderRepositoryGateWayInterfaces := order_repository_gateway.NewOrderRepositoryGateWay(orderRepositoryInterfaces)
		useCaseModifyOrder := usecase_orders.NewModifyOrderUseCase(orderRepositoryGateWayInterfaces)
		modifyOrderController := order.NewModifyOrderController(useCaseModifyOrder)

		err := c.Bind(&reqOrder)
		if err != nil {
			c.JSON(400, "INVALID_REQUEST")
			fmt.Printf("error : %s \n", err.Error())
			return err
		}

		err, orderRsp := modifyOrderController.ProcessEditRequest(c.Request().Context(), reqOrder)
		if err != nil {
			fmt.Printf("Error updating order with orderNumber: %d with error: %s \n", reqOrder.OrderNumber, err.Error())
			c.JSON(500, "INTERNAL_ERROR_SERVER")
		}
		c.JSON(200, orderRsp)
		return nil
	})
	return router
}

func DeleteOrder(router *echo.Echo) *echo.Echo {
	router.DELETE("/delete-order", func(c echo.Context) error {
		orderRequest := c.QueryParam("orderNumber")
		mongoClientInterfaces := mongo_client.NewMongoClientInterfaces()
		orderRepositoryInterfaces := order_repository.NewOrderRepository(mongoClientInterfaces)
		orderRepositoryGateWayInterfaces := order_repository_gateway.NewOrderRepositoryGateWay(orderRepositoryInterfaces)
		DeleteOrderUseCaseInterfaces := usecase_orders.NewDeleteOrderUseCase(orderRepositoryGateWayInterfaces)
		deleteOrderController := order.NewDeleteOrderController(DeleteOrderUseCaseInterfaces)
		if orderRequest == "" {
			return c.JSON(404, "INVALID_REQUEST")
		}
		orderNumber, err := strconv.Atoi(orderRequest)
		if err != nil {
			fmt.Printf("Error delting order with orderNumber: %s with error: %s \n", orderRequest, err.Error())
			return c.JSON(500, "INTERNAL_ERROR_SERVER")
		}
		err, orderDeleted := deleteOrderController.ProcessDeleteOrder(c.Request().Context(), orderNumber)
		if err != nil {
			fmt.Printf("Error delting order with orderNumber: %s with error: %s \n", orderRequest, err.Error())
			return c.JSON(500, "INTERNAL_ERROR_SERVER")
		}
		c.JSON(200, orderDeleted)
		return nil
	})
	return router
}
