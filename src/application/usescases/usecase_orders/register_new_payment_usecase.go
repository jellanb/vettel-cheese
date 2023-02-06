package usecase_orders

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/domain/status_domain/payment_status"
	"vettel-backend-app/src/interfaces/geteways/order_repository_gateway"
)

type RegisterNewPaymentUseCaseInterfaces interface {
	ProcessNewPayment(ctx context.Context, order entity.Order) (error, *entity.Order)
}

type RegisterNewPaymentUseCase struct {
	orderRepositoryGateway order_repository_gateway.OrderRepositoryGateWayInterfaces
}

func NewRegisterNewPaymentUseCase(orderRepositoryGateway order_repository_gateway.OrderRepositoryGateWayInterfaces) RegisterNewPaymentUseCaseInterfaces {
	return &RegisterNewPaymentUseCase{orderRepositoryGateway}
}

func (p RegisterNewPaymentUseCase) ProcessNewPayment(ctx context.Context, order entity.Order) (error, *entity.Order) {
	fmt.Printf("Validating if order exist for orderNumber: %d \n", order.OrderNumber)
	err, orderFound := p.orderRepositoryGateway.FindOrderByOrderId(ctx, order.OrderNumber)
	if err != nil {
		return err, nil
	}
	if orderFound.OrderNumber == 0 {
		fmt.Printf("Error trying update order with orderNumber: %d,order not exist \n", order.OrderNumber)
		return errors.New("ORDER_NOT_FOUND"), nil
	}
	fmt.Printf("Validation ok order exist to update with orderNumber: %d \n", order.OrderNumber)

	orderEntity := orderFound
	orderEntity.Payment.CreditPaid = order.Payment.CreditPaid
	orderEntity.Payment.SummaryPaidAmount += totalCredit(*orderEntity)
	if orderEntity.SaleAmount <= orderEntity.Payment.SummaryPaidAmount {
		orderEntity.Payment.Status = payment_status.Paid
	}

	fmt.Printf("Updating order wih orderNumber: %d \n", order.OrderNumber)
	err, orderUpdated := p.orderRepositoryGateway.UpdateOrder(ctx, *orderEntity)
	if err != nil {
		return err, nil
	}
	fmt.Printf("Order updated with orderNumber: %d \n", order.OrderNumber)

	return nil, orderUpdated
}

func totalCredit(order entity.Order) int {
	var totalPaid int
	for _, quota := range order.Payment.CreditPaid {
		totalPaid = quota.Amount
	}
	return totalPaid
}
