package inventory_repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"vettel-backend-app/src/infrastructure/database/collections"
	"vettel-backend-app/src/infrastructure/database/mongo_client"
	"vettel-backend-app/src/shared/utils/commontypes"
)

type InventoryRepositoryInterfaces interface {
	GetAll(ctx context.Context) (error, []*collections.Inventory)
	FindInventoryItemByBarCode(ctx context.Context, barcode string) (error, *collections.Inventory)
	AddInventoryItem(ctx context.Context, inventory collections.Inventory) (error, *collections.Inventory)
	UpdateInventory(ctx context.Context, inventory collections.Inventory) (error, *collections.Inventory)
	DeleteInventory(ctx context.Context, id string) (error, *string)
}

type InventoryRepository struct {
	mongoClientInterfaces mongo_client.MongoClientInterfaces
}

func NewInventoryRepositoryGateWay(mongoClientInterfaces mongo_client.MongoClientInterfaces) InventoryRepositoryInterfaces {
	return &InventoryRepository{mongoClientInterfaces}
}

func (r InventoryRepository) FindInventoryItemByBarCode(ctx context.Context, barcode string) (error, *collections.Inventory) {
	ctx, coll, err := r.mongoClientInterfaces.LoadDatabase(commontypes.CollectionInventory)
	if err != nil {
		fmt.Printf("Error attaching collection name %s with error :%s \n", commontypes.CollectionInventory, err.Error())
		return err, nil
	}
	var inventory collections.Inventory

	err = coll.FindOne(ctx, bson.M{"product.barcode": barcode}).Decode(&inventory)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("Noting document found: %s \n", err.Error())
		return nil, nil
	}
	if err != nil {
		fmt.Printf("Error getting inventory by barcode: %s \n", barcode)
		return err, nil
	}
	return nil, &inventory
}

func (r InventoryRepository) GetAll(ctx context.Context) (error, []*collections.Inventory) {
	var inventory []*collections.Inventory
	ctx, coll, err := r.mongoClientInterfaces.LoadDatabase(commontypes.CollectionInventory)
	if err != nil {
		fmt.Printf("Error attaching collection name %s with error :%s \n", commontypes.CollectionInventory, err.Error())
		return err, nil
	}

	cursor, err := coll.Find(ctx, bson.M{})
	err = cursor.All(ctx, &inventory)
	if err != nil {
		return err, nil
	}
	return nil, inventory
}

func (r InventoryRepository) AddInventoryItem(ctx context.Context, inventory collections.Inventory) (error, *collections.Inventory) {
	ctx, coll, err := r.mongoClientInterfaces.LoadDatabase(commontypes.CollectionInventory)

	if err != nil {
		fmt.Printf("Error attaching collection name %s with error :%s \n", commontypes.CollectionInventory, err.Error())
		return err, nil
	}

	_, err = coll.InsertOne(ctx, inventory)

	if err != nil {
		return err, nil
	}
	return nil, &inventory
}

func (r InventoryRepository) UpdateInventory(ctx context.Context, inventory collections.Inventory) (error, *collections.Inventory) {
	ctx, coll, err := r.mongoClientInterfaces.LoadDatabase(commontypes.CollectionInventory)
	if err != nil {
		fmt.Printf("Error attaching collection name %s with error :%s \n", commontypes.CollectionInventory, err.Error())
		return err, nil
	}
	fmt.Printf("inventory: %v \n", inventory)
	_, err = coll.UpdateMany(ctx, bson.D{{"product.barcode", inventory.Product.Barcode}}, bson.D{{"$set", inventory}})
	if err != nil {
		return err, nil
	}
	return nil, &inventory
}

func (r InventoryRepository) DeleteInventory(ctx context.Context, bardCode string) (error, *string) {
	ctx, coll, err := r.mongoClientInterfaces.LoadDatabase(commontypes.CollectionInventory)
	if err != nil {
		fmt.Printf("Error attaching collection name %s with error :%s \n", commontypes.CollectionInventory, err.Error())
		return err, nil
	}
	_, err = coll.DeleteOne(ctx, bson.D{{"product.barcode", bardCode}})
	if err != nil {
		fmt.Printf("Error in database cannot delete inventory item with barcode: %s and error: %s \n", bardCode, err.Error())
		return err, nil
	}
	return nil, &bardCode
}
