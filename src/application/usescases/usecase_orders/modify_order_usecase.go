package usecase_orders

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/interfaces/geteways/order_repository_gateway"
)

type ModifyOrderUseCaseInterfaces interface {
	ProcessEditOrder(ctx context.Context, order entity.Order) (error, *entity.Order)
}

type ModifyOrderUseCase struct {
	orderRepositoryGateway order_repository_gateway.OrderRepositoryGateWayInterfaces
}

func NewModifyOrderUseCase(orderRepositoryGateway order_repository_gateway.OrderRepositoryGateWayInterfaces) ModifyOrderUseCaseInterfaces {
	return &ModifyOrderUseCase{orderRepositoryGateway}
}

func (m ModifyOrderUseCase) ProcessEditOrder(ctx context.Context, order entity.Order) (error, *entity.Order) {
	fmt.Printf("Validating if order exist for orderNumber: %d \n", order.OrderNumber)
	err, orderFound := m.orderRepositoryGateway.FindOrderByOrderId(ctx, order.OrderNumber)
	if err != nil {
		return err, nil
	}
	if orderFound.OrderNumber == 0 {
		fmt.Printf("Error trying update order with orderNumber: %d,order not exist \n", order.OrderNumber)
		return errors.New("ORDER_NOT_FOUND"), nil
	}
	fmt.Printf("Validation ok order exist to update with orderNumber: %d \n", order.OrderNumber)
	fmt.Printf("Updating order wih orderNumber: %d \n", order.OrderNumber)
	err, orderUpdated := m.orderRepositoryGateway.UpdateOrder(ctx, order)
	if err != nil {
		return err, nil
	}
	fmt.Printf("Order updated with orderNumber: %d \n", order.OrderNumber)
	return nil, orderUpdated
}
