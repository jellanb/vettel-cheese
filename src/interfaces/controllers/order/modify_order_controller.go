package order

import (
	"context"
	"fmt"
	"vettel-backend-app/src/application/usescases/usecase_orders"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
	"vettel-backend-app/src/interfaces/geteways/order_repository_gateway/parser"
)

type ModifyOrderControllerInterfaces interface {
	ProcessEditRequest(ctx context.Context, orderChanges request.Order) (error, *response.Order)
}

type ModifyOrderController struct {
	modifyOrderUseCaseInterfaces usecase_orders.ModifyOrderUseCaseInterfaces
}

func NewModifyOrderController(modifyOrderUseCase usecase_orders.ModifyOrderUseCaseInterfaces) ModifyOrderControllerInterfaces {
	return &ModifyOrderController{modifyOrderUseCase}
}

func (m ModifyOrderController) ProcessEditRequest(ctx context.Context, orderChanges request.Order) (error, *response.Order) {
	fmt.Printf("Init process update order with orderNumber: %d \n", orderChanges.OrderNumber)
	orderEntity := parser.OrderRequestToEntity(orderChanges)
	err, orderUpdated := m.modifyOrderUseCaseInterfaces.ProcessEditOrder(ctx, orderEntity)
	if err != nil {
		fmt.Printf("Error in updating order: %d ago with error: %s\n", orderChanges.OrderNumber, err.Error())
		return err, nil
	}
	orderRsp := parser.OrderEntityToResponse(*orderUpdated)
	fmt.Printf("Finish process update order with orderNumber: %d \n", orderChanges.OrderNumber)
	return nil, &orderRsp
}
