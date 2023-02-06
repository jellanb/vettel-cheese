package routes

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"vettel-backend-app/src/application/usescases/usecase_orders"
	"vettel-backend-app/src/infrastructure/database/mongo_client"
	"vettel-backend-app/src/infrastructure/database/order_repository"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/interfaces/controllers/order"
	"vettel-backend-app/src/interfaces/geteways/order_repository_gateway"
)

func ProcessPayment(router *echo.Echo) *echo.Echo {
	router.PUT("/register-payment", func(c echo.Context) error {
		var reqOrder request.Order
		mongoClientInterfaces := mongo_client.NewMongoClientInterfaces()
		orderRepositoryInterfaces := order_repository.NewOrderRepository(mongoClientInterfaces)
		orderRepositoryGateWayInterfaces := order_repository_gateway.NewOrderRepositoryGateWay(orderRepositoryInterfaces)
		useCaseRegisterPayment := usecase_orders.NewRegisterNewPaymentUseCase(orderRepositoryGateWayInterfaces)
		processPaymentController := order.NewProcessNewPaymentController(useCaseRegisterPayment)

		err := c.Bind(&reqOrder)
		if err != nil {
			c.JSON(400, "INVALID_REQUEST")
			fmt.Printf("error : %s \n", err.Error())
			return err
		}

		err, orderProcessed := processPaymentController.RegisterPayment(c.Request().Context(), reqOrder)
		if err != nil {
			fmt.Printf("Error updating order with orderNumber: %d with error: %s \n", reqOrder.OrderNumber, err.Error())
			c.JSON(500, "INTERNAL_ERROR_SERVER")
		}
		c.JSON(200, orderProcessed)
		return nil
	})

	return router
}
