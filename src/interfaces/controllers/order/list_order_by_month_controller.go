package order

import (
	"fmt"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/interfaces/geteways/order_repository_gateway"
)

type ListOrderByMonthControllerInterfaces interface {
	ListOrderByMonthController() (error, []*entity.Order)
}

type ListOrderByMonthController struct {
	orderRepositoryGateWayInterfaces order_repository_gateway.OrderRepositoryGateWayInterfaces
}

func NewListOrderByMonthController(orderRepositoryGateWayInterfaces order_repository_gateway.OrderRepositoryGateWayInterfaces) ListOrderByMonthControllerInterfaces {
	return &ListOrderByMonthController{orderRepositoryGateWayInterfaces}
}

func (l ListOrderByMonthController) ListOrderByMonthController() (error, []*entity.Order) {
	fmt.Println("Init process to list order ago")
	err, ordersEntity := l.orderRepositoryGateWayInterfaces.FindOrdersLastMonth()
	if err != nil {
		fmt.Printf("Error in process to list order ago with error: %s\n", err.Error())
		return err, nil
	}
	fmt.Println("Finish process to list order ago")
	return nil, ordersEntity
}
