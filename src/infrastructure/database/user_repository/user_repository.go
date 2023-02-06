package user_repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"vettel-backend-app/src/infrastructure/database/collections"
	"vettel-backend-app/src/infrastructure/database/mongo_client"
	"vettel-backend-app/src/shared/utils/commontypes"
)

type UsersRepositoryInterfaces interface {
	FindUserByEmailAndPassword(ctx context.Context, user collections.Users) (error, *collections.Users)
	FindUserByEmail(ctx context.Context, user collections.Users) (error, *collections.Users)
	Register(ctx context.Context, user collections.Users) (error, *collections.Users)
}

type UsersRepository struct {
	mongoClientInterfaces mongo_client.MongoClientInterfaces
}

func NewUsersRepository(mongoClientInterfaces mongo_client.MongoClientInterfaces) UsersRepositoryInterfaces {
	return &UsersRepository{mongoClientInterfaces}
}

func (u *UsersRepository) FindUserByEmailAndPassword(ctx context.Context, user collections.Users) (error, *collections.Users) {
	ctx, coll, err := u.mongoClientInterfaces.LoadDatabase(commontypes.CollectionUsers)
	if err != nil {
		fmt.Printf("Error attaching collection name %s", commontypes.CollectionUsers)
		return err, nil
	}
	var result collections.Users
	err = coll.FindOne(ctx, bson.M{"email": user.Email, "password": user.Password}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("User not found by email %s and password %s \n", user.Email, user.Password)
		return nil, nil
	}
	if err != mongo.ErrNoDocuments && err != nil {
		return err, nil
	}
	return nil, &result
}

func (u *UsersRepository) FindUserByEmail(ctx context.Context, user collections.Users) (error, *collections.Users) {
	ctx, coll, err := u.mongoClientInterfaces.LoadDatabase(commontypes.CollectionUsers)
	if err != nil {
		fmt.Printf("Error attaching collection name %s", commontypes.CollectionOrder)
		return err, nil
	}
	var userResult collections.Users
	err = coll.FindOne(ctx, bson.M{"email": user.Email}).Decode(&userResult)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("User not found by email: %s\n", user.Email)
		return nil, nil
	}
	return nil, &userResult
}

func (u *UsersRepository) Register(ctx context.Context, user collections.Users) (error, *collections.Users) {
	ctx, coll, err := u.mongoClientInterfaces.LoadDatabase(commontypes.CollectionUsers)
	if err != nil {
		fmt.Printf("Error attaching collection name %s", commontypes.CollectionOrder)
		return err, nil
	}
	_, err = coll.InsertOne(
		ctx,
		user)

	if err != nil {
		return err, nil
	}
	return nil, &user
}
