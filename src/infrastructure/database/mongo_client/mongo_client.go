package mongo_client

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type MongoClientInterfaces interface {
	LoadDatabase(collectionName string) (context.Context, mongo.Collection, error)
	Connect(ctx context.Context, uri string) (*mongo.Client, error)
	AttachCollection(client *mongo.Client, dbname string, collectionName string) *mongo.Collection
}

type MongoClient struct {
	MongoClientInterfaces
}

func NewMongoClientInterfaces() *MongoClient {
	return &MongoClient{}
}

func (m *MongoClient) LoadDatabase(collectionName string) (context.Context, mongo.Collection, error) {
	uri := os.Getenv("MONGO_URL")
	if uri == "" {
		log.Fatal("You must set your 'MONGO_URL' environmental variable.")
	}
	dbname := os.Getenv("DATABASE")
	if uri == "" {
		log.Fatal("You must set your 'DATABASE' environmental variable.")
	}
	ctx := context.TODO()
	client, err := m.Connect(ctx, uri)
	if err != nil {
		log.Fatalf("Cannot connecto to database %s", err)
	}
	coll := m.AttachCollection(client, dbname, collectionName)

	return ctx, *coll, nil

}

func (m *MongoClient) Connect(ctx context.Context, uri string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (m *MongoClient) AttachCollection(client *mongo.Client, dbname string, collectionName string) *mongo.Collection {
	return client.Database(dbname).Collection(collectionName)
}
