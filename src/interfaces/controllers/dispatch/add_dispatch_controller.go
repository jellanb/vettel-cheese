package dispatch

import (
	"context"
	"vettel-backend-app/src/application/usescases/usecase_orders"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
	"vettel-backend-app/src/interfaces/geteways/dispatch_repository_gateway/parser"
	orderResponseParser "vettel-backend-app/src/interfaces/geteways/order_repository_gateway/parser"
)

type AddDispatchControllerInterfaces interface {
	AddDispatchController(ctx context.Context, dispatch request.Dispatch) (error, *response.Order)
}

type AddDispatchController struct {
	registerNewOrderUseCaseInterfaces usecase_orders.RegisterNewOrderUseCaseInterfaces
}

func NewAddDispatchController(registerNewOrderUseCaseInterfaces usecase_orders.RegisterNewOrderUseCaseInterfaces) AddDispatchControllerInterfaces {
	return &AddDispatchController{registerNewOrderUseCaseInterfaces}
}

func (c AddDispatchController) AddDispatchController(ctx context.Context, dispatch request.Dispatch) (error, *response.Order) {
	dispatchEntity := parser.DispatchRequestToEntity(dispatch)
	orderEntity := orderResponseParser.GenerateOrderEntityFromDispatch(dispatchEntity)
	err, order := c.registerNewOrderUseCaseInterfaces.RegisterNewOrderUseCase(ctx, orderEntity)
	if err != nil {
		return err, nil
	}
	orderResponse := orderResponseParser.OrderEntityToResponse(*order)
	return nil, &orderResponse
}
