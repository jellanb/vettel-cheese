package order

import (
	"context"
	"fmt"
	"vettel-backend-app/src/application/usescases/usecase_orders"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
	"vettel-backend-app/src/interfaces/geteways/order_repository_gateway/parser"
)

type DeleteOrderControllerInterfaces interface {
	ProcessDeleteOrder(ctx context.Context, order int) (error, *response.Order)
}

type DeleteOrderController struct {
	orderDeleteUseCaseInterfaces usecase_orders.DeleteOrderUseCaseInterfaces
}

func NewDeleteOrderController(orderDeleteUseCaseInterfaces usecase_orders.DeleteOrderUseCaseInterfaces) DeleteOrderControllerInterfaces {
	return &DeleteOrderController{orderDeleteUseCaseInterfaces}
}

func (d DeleteOrderController) ProcessDeleteOrder(ctx context.Context, orderNumber int) (error, *response.Order) {
	fmt.Printf("Init process to delete order with orderNumber: %d \n", orderNumber)
	err, orderDeleted := d.orderDeleteUseCaseInterfaces.ProcessDeleteOrder(ctx, orderNumber)
	if err != nil {
		return err, nil
	}
	orderDeleteResp := parser.OrderEntityToResponse(*orderDeleted)
	fmt.Printf("Finish process to delete order with orderNumber: %d \n", orderNumber)
	return nil, &orderDeleteResp
}
