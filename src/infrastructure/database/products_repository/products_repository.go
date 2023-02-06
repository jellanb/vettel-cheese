package products_repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"vettel-backend-app/src/infrastructure/database/collections"
	"vettel-backend-app/src/infrastructure/database/mongo_client"
	"vettel-backend-app/src/shared/utils/commontypes"
)

type ProductsRepositoryInterfaces interface {
	FindAllProducts() (error, []collections.Product)
	FindProductByBarcode(barcode string) (error, *collections.Product)
	InsertOneProduct(ctx context.Context, newProduct collections.Product) (error, *collections.Product)
	UpdateOneProductById(ctx context.Context, product collections.Product) (error, *collections.Product)
	DeleteProductById(ctx context.Context, barcode string) (error, *string)
}

type ProductsRepository struct {
	mongoClientInterfaces mongo_client.MongoClientInterfaces
}

func NewProductsRepository(mongoClientInterfaces mongo_client.MongoClientInterfaces) ProductsRepositoryInterfaces {
	return &ProductsRepository{mongoClientInterfaces}
}

func (r ProductsRepository) FindAllProducts() (error, []collections.Product) {
	var productsCollection []collections.Product
	ctx, coll, err := r.mongoClientInterfaces.LoadDatabase(commontypes.CollectionProducts)
	if err != nil {
		return err, nil
	}
	cursor, err := coll.Find(ctx, bson.M{})
	err = cursor.All(ctx, &productsCollection)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("Noting document found: %s \n", err.Error())
		return nil, nil
	}
	if err != nil {
		return err, nil
	}
	return nil, productsCollection
}

func (r ProductsRepository) FindProductByBarcode(barcode string) (error, *collections.Product) {
	ctx, coll, err := r.mongoClientInterfaces.LoadDatabase(commontypes.CollectionProducts)
	var productsCollection collections.Product
	if err != nil {
		return err, nil
	}
	err = coll.FindOne(ctx, bson.D{{"barcode", barcode}}).Decode(&productsCollection)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("Noting document found: %s \n", err.Error())
		return nil, nil
	}
	if err != nil {
		return err, nil
	}
	return nil, &productsCollection
}

func (r ProductsRepository) InsertOneProduct(ctx context.Context, newProduct collections.Product) (error, *collections.Product) {
	ctx, coll, err := r.mongoClientInterfaces.LoadDatabase(commontypes.CollectionProducts)
	if err != nil {
		return err, nil
	}
	_, err = coll.InsertOne(ctx, newProduct)
	if err != nil {
		return err, nil
	}
	return nil, &newProduct
}

func (r ProductsRepository) UpdateOneProductById(ctx context.Context, product collections.Product) (error, *collections.Product) {
	ctx, coll, err := r.mongoClientInterfaces.LoadDatabase(commontypes.CollectionProducts)
	if err != nil {
		return err, nil
	}
	_, err = coll.UpdateOne(ctx, bson.D{{"barcode", product.Barcode}}, bson.D{{"$set", product}})
	if err != nil {
		return err, nil
	}
	return nil, &product
}

func (r ProductsRepository) DeleteProductById(ctx context.Context, barcode string) (error, *string) {
	ctx, coll, err := r.mongoClientInterfaces.LoadDatabase(commontypes.CollectionProducts)
	if err != nil {
		return err, nil
	}
	_, err = coll.DeleteOne(ctx, bson.D{{"barcode", barcode}})
	if err != nil {
		return err, nil
	}
	return nil, &barcode
}
