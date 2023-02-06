package usecase_orders

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/interfaces/geteways/order_repository_gateway"
)

type DeleteOrderUseCaseInterfaces interface {
	ProcessDeleteOrder(ctx context.Context, orderNumber int) (error, *entity.Order)
}

type DeleteOrderUseCase struct {
	orderRepositoryGateway order_repository_gateway.OrderRepositoryGateWayInterfaces
}

func NewDeleteOrderUseCase(orderRepositoryGateway order_repository_gateway.OrderRepositoryGateWayInterfaces) DeleteOrderUseCaseInterfaces {
	return &DeleteOrderUseCase{orderRepositoryGateway}
}

func (d DeleteOrderUseCase) ProcessDeleteOrder(ctx context.Context, orderNumber int) (error, *entity.Order) {
	fmt.Printf("Validating if order exist for orderNumber: %d \n", orderNumber)
	err, orderFound := d.orderRepositoryGateway.FindOrderByOrderId(ctx, orderNumber)
	if err != nil {
		return err, nil
	}
	if orderFound.OrderNumber == 0 {
		fmt.Printf("Error trying update order with orderNumber: %d,order not exist \n", orderNumber)
		return errors.New("ORDER_NOT_FOUND"), nil
	}
	fmt.Printf("Validation ok order exist to update with orderNumber: %d \n", orderNumber)
	fmt.Printf("Deleting order wih orderNumber: %d \n", orderNumber)
	err, orderDeleted := d.orderRepositoryGateway.DeleteOrderByOrderId(ctx, *orderFound)
	if err != nil {
		return err, nil
	}
	fmt.Printf("Order deleted wih orderNumber: %d \n", orderNumber)
	return nil, orderDeleted
}
