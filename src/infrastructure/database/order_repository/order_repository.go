package order_repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"vettel-backend-app/src/infrastructure/database/collections"
	"vettel-backend-app/src/infrastructure/database/mongo_client"
	"vettel-backend-app/src/shared/utils/commontypes"
)

type OrderRepositoryInterfaces interface {
	FindOrderByOrderId(ctx context.Context, order int) (error, *collections.Order)
	FindOrdersLastMonth() (error, *[]collections.Order)
	InsertOneOrder(ctx context.Context, order collections.Order) (error, *collections.Order)
	FindLastOrderId(ctx context.Context) (error, *collections.Order)
	UpdateOrder(ctx context.Context, order collections.Order) (error, *collections.Order)
	DeleteOrderByOrderNumber(ctx context.Context, order collections.Order) (error, *collections.Order)
}

type OrderRepository struct {
	mongoClientInterfaces mongo_client.MongoClientInterfaces
}

func NewOrderRepository(mongoClientInterfaces mongo_client.MongoClientInterfaces) OrderRepositoryInterfaces {
	return &OrderRepository{mongoClientInterfaces}
}

func (o OrderRepository) FindOrdersLastMonth() (error, *[]collections.Order) {
	var orders []collections.Order
	days := bson.M{
		"$gte": primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -30))}

	ctx, coll, err := o.mongoClientInterfaces.LoadDatabase(commontypes.CollectionOrder)
	if err != nil {
		fmt.Printf("Error attaching collection name %s \n", commontypes.CollectionOrder)
		return err, nil
	}

	cursor, err := coll.Find(ctx, bson.M{"date": days})
	err = cursor.All(ctx, &orders)
	if err != nil {
		return err, nil
	}
	return nil, &orders
}

func (o OrderRepository) InsertOneOrder(ctx context.Context, order collections.Order) (error, *collections.Order) {
	ctx, coll, err := o.mongoClientInterfaces.LoadDatabase(commontypes.CollectionOrder)
	if err != nil {
		fmt.Printf("Error attaching collection name %s \n", commontypes.CollectionOrder)
		return err, nil
	}

	_, err = coll.InsertOne(ctx, order)
	if err != nil {
		return err, nil
	}
	return nil, &order
}

func (o OrderRepository) FindLastOrderId(ctx context.Context) (error, *collections.Order) {
	ctx, coll, err := o.mongoClientInterfaces.LoadDatabase(commontypes.CollectionOrder)
	if err != nil {
		fmt.Printf("Error attaching collection name %s \n", commontypes.CollectionOrder)
		return err, nil
	}
	var orderCollection collections.Order
	var queryOption options.FindOneOptions
	queryOption.SetSort(bson.D{{"orderNumber", -1}})
	err = coll.FindOne(ctx, bson.M{}, &queryOption).Decode(&orderCollection)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("Noting document found: %s \n", err.Error())
		return nil, nil
	}
	return nil, &orderCollection
}

func (o OrderRepository) FindOrderByOrderId(ctx context.Context, orderNumber int) (error, *collections.Order) {
	ctx, coll, err := o.mongoClientInterfaces.LoadDatabase(commontypes.CollectionOrder)
	if err != nil {
		fmt.Printf("Error attaching collection name %s \n", commontypes.CollectionOrder)
		return err, nil
	}
	var orderCollection collections.Order
	err = coll.FindOne(ctx, bson.M{"orderNumber": orderNumber}).Decode(&orderCollection)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("Noting document found: %s \n", err.Error())
		return nil, nil
	}
	return nil, &orderCollection

}

func (o OrderRepository) UpdateOrder(ctx context.Context, order collections.Order) (error, *collections.Order) {
	ctx, coll, err := o.mongoClientInterfaces.LoadDatabase(commontypes.CollectionOrder)
	if err != nil {
		fmt.Printf("Error attaching collection name %s \n", commontypes.CollectionOrder)
		return err, nil
	}
	filter := bson.D{{"orderNumber", order.OrderNumber}}
	update := bson.D{{"$set", order}}
	_, err = coll.UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Printf("Error updating order with orderNumber: %d, %s \n", order.OrderNumber, commontypes.CollectionOrder)
		return err, nil
	}
	return nil, &order
}

func (o OrderRepository) DeleteOrderByOrderNumber(ctx context.Context, order collections.Order) (error, *collections.Order) {
	ctx, coll, err := o.mongoClientInterfaces.LoadDatabase(commontypes.CollectionOrder)
	if err != nil {
		fmt.Printf("Error attaching collection name %s \n", commontypes.CollectionOrder)
		return err, nil
	}
	filter := bson.D{{"orderNumber", order.OrderNumber}}
	_, err = coll.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Printf("Error deleting order with orderNumber: %d, %s \n", order.OrderNumber, commontypes.CollectionOrder)
		return err, nil
	}
	return nil, &order
}
