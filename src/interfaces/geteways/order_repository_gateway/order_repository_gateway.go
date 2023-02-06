package order_repository_gateway

import (
	"context"
	"encoding/json"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/database/order_repository"
	"vettel-backend-app/src/interfaces/geteways/order_repository_gateway/parser"
)

type OrderRepositoryGateWayInterfaces interface {
	FindOrderByOrderId(ctx context.Context, order int) (error, *entity.Order)
	FindOrdersLastMonth() (error, []*entity.Order)
	InsertOneOrder(ctx context.Context, order entity.Order) (error, *entity.Order)
	FindLastOrderId(ctx context.Context) (error, *entity.Order)
	UpdateOrder(ctx context.Context, order entity.Order) (error, *entity.Order)
	DeleteOrderByOrderId(ctx context.Context, order entity.Order) (error, *entity.Order)
}

type OrderRepositoryGateWay struct {
	orderRepositoryInterfaces order_repository.OrderRepositoryInterfaces
}

func NewOrderRepositoryGateWay(orderRepositoryInterfaces order_repository.OrderRepositoryInterfaces) OrderRepositoryGateWayInterfaces {
	return &OrderRepositoryGateWay{orderRepositoryInterfaces}
}

func (o OrderRepositoryGateWay) FindOrdersLastMonth() (error, []*entity.Order) {
	err, ordersCollection := o.orderRepositoryInterfaces.FindOrdersLastMonth()
	if err != nil {
		return err, nil
	}
	jsonOrders, err := json.Marshal(ordersCollection)
	if err != nil {
		return err, nil
	}
	var ordersEntity []*entity.Order
	err = json.Unmarshal(jsonOrders, &ordersEntity)
	if err != nil {
		return err, nil
	}
	return nil, ordersEntity
}

func (o OrderRepositoryGateWay) InsertOneOrder(ctx context.Context, order entity.Order) (error, *entity.Order) {
	orderCollection := parser.OrderEntityToCollection(order)
	err, orderCollResult := o.orderRepositoryInterfaces.InsertOneOrder(ctx, orderCollection)
	if err != nil {
		return err, nil
	}
	orderEntityResult := parser.OrderCollectionToEntity(*orderCollResult)
	return nil, &orderEntityResult
}

func (o OrderRepositoryGateWay) FindLastOrderId(ctx context.Context) (error, *entity.Order) {
	err, lastOrderId := o.orderRepositoryInterfaces.FindLastOrderId(ctx)
	if err != nil {
		return err, nil
	}
	if lastOrderId == nil {
		var order entity.Order
		return nil, &order
	}
	orderEntity := parser.OrderCollectionToEntity(*lastOrderId)
	return nil, &orderEntity
}

func (o OrderRepositoryGateWay) FindOrderByOrderId(ctx context.Context, orderNumber int) (error, *entity.Order) {
	err, orderCollResult := o.orderRepositoryInterfaces.FindOrderByOrderId(ctx, orderNumber)
	if err != nil {
		return err, nil
	}
	orderEntityResul := parser.OrderCollectionToEntity(*orderCollResult)
	return nil, &orderEntityResul
}

func (o OrderRepositoryGateWay) UpdateOrder(ctx context.Context, order entity.Order) (error, *entity.Order) {
	orderCollection := parser.OrderEntityToCollection(order)
	err, updatedResult := o.orderRepositoryInterfaces.UpdateOrder(ctx, orderCollection)
	if err != nil {
		return err, nil
	}
	orderEntityResult := parser.OrderCollectionToEntity(*updatedResult)
	return nil, &orderEntityResult
}

func (o OrderRepositoryGateWay) DeleteOrderByOrderId(ctx context.Context, order entity.Order) (error, *entity.Order) {
	orderCollection := parser.OrderEntityToCollection(order)
	err, deletedResult := o.orderRepositoryInterfaces.DeleteOrderByOrderNumber(ctx, orderCollection)
	if err != nil {
		return err, nil
	}
	orderEntityResult := parser.OrderCollectionToEntity(*deletedResult)
	return nil, &orderEntityResult
}
