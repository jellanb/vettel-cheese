package order

import (
	"context"
	"fmt"
	"vettel-backend-app/src/application/usescases/usecase_orders"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
	"vettel-backend-app/src/interfaces/geteways/order_repository_gateway/parser"
)

type ProcessNewPaymentControllerInterfaces interface {
	RegisterPayment(ctx context.Context, order request.Order) (error, *response.Order)
}

type ProcessNewPaymentController struct {
	registerNewPaymentInterfaces usecase_orders.RegisterNewPaymentUseCaseInterfaces
}

func NewProcessNewPaymentController(registerNewPaymentInterfaces usecase_orders.RegisterNewPaymentUseCaseInterfaces) ProcessNewPaymentControllerInterfaces {
	return &ProcessNewPaymentController{registerNewPaymentInterfaces}
}

func (c ProcessNewPaymentController) RegisterPayment(ctx context.Context, order request.Order) (error, *response.Order) {
	fmt.Printf("Init process to process payment for order with orderNumber: %d \n", order.OrderNumber)
	orderEntity := parser.OrderRequestToEntity(order)
	err, orderProcessed := c.registerNewPaymentInterfaces.ProcessNewPayment(ctx, orderEntity)
	if err != nil {
		fmt.Printf("Error processing payment for order: %d ago with error: %s\n", order.OrderNumber, err.Error())
		return err, nil
	}
	orderRsp := parser.OrderEntityToResponse(*orderProcessed)
	fmt.Printf("Finish process payment for order with orderNumber: %d \n", order.OrderNumber)
	return nil, &orderRsp
}
