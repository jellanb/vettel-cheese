package usecase_orders

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/interfaces/geteways/inventory_repository_gateway"
	"vettel-backend-app/src/interfaces/geteways/order_repository_gateway"
)

type RegisterNewOrderUseCaseInterfaces interface {
	RegisterNewOrderUseCase(ctx context.Context, dispatch entity.Order) (error, *entity.Order)
}

type RegisterNewOrderUseCase struct {
	orderRepositoryGateWayInterfaces     order_repository_gateway.OrderRepositoryGateWayInterfaces
	inventoryRepositoryGateWayInterfaces inventory_repository_gateway.InventoryRepositoryGatewayInterfaces
}

func NewRegisterNewOrderUseCase(
	orderRepositoryGateWayInterfaces order_repository_gateway.OrderRepositoryGateWayInterfaces,
	inventoryRepositoryGateWayInterfaces inventory_repository_gateway.InventoryRepositoryGatewayInterfaces,
) RegisterNewOrderUseCaseInterfaces {
	return &RegisterNewOrderUseCase{
		orderRepositoryGateWayInterfaces,
		inventoryRepositoryGateWayInterfaces,
	}
}

func (u RegisterNewOrderUseCase) RegisterNewOrderUseCase(ctx context.Context, order entity.Order) (error, *entity.Order) {
	err, lastOrder := u.orderRepositoryGateWayInterfaces.FindLastOrderId(ctx)
	if err != nil {
		return err, nil
	}

	order.OrderNumber = lastOrder.OrderNumber + 1
	order.Payment.Iva = calculateIVA(order.SaleAmount)
	order.Payment.NetAmount = order.Payment.Amount - order.Payment.Iva

	for _, product := range order.Products {
		err, inventoryFound := u.inventoryRepositoryGateWayInterfaces.FindInventoryItemByBarcode(ctx, product.Barcode)
		if err != nil {
			return err, nil
		}
		if inventoryFound == nil {
			message := fmt.Sprintf("Cannot decrease inventory for barco: %s", inventoryFound.Product.Barcode)
			return errors.New(message), nil
		}
		fmt.Printf("Inventory decreasing item by barcode: %s from order number: %d", product.Barcode, &order.OrderNumber)
		inventoryDecrease := inventoryFound
		inventoryDecrease.Quantity -= 1
		err, _ = u.inventoryRepositoryGateWayInterfaces.UpdateInventoryItem(ctx, *inventoryDecrease)
		if err != nil {
			return err, nil
		}
		fmt.Printf("Inventory decreased item by barcode: %s from order number: %d", product.Barcode, &order.OrderNumber)
	}
	err, orderEntitySaved := u.orderRepositoryGateWayInterfaces.InsertOneOrder(ctx, order)
	if err != nil {
		return err, nil
	}
	return nil, orderEntitySaved
}

func calculateIVA(orderAmount int) int {
	return orderAmount * 19 / 100
}
