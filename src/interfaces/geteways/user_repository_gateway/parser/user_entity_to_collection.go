package parser

import (
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/database/collections"
)

type UserCollectionParserGetWayInterfaces interface {
	UserEntityToCollection(userEntity entity.User) collections.Users
}

type UserCollectionParserGateWay struct{}

func NewUserCollectionParserGateWay() *UserCollectionParserGateWay {
	return &UserCollectionParserGateWay{}
}

func (u UserCollectionParserGateWay) UserEntityToCollection(userEntity entity.User) collections.Users {
	var userCollection collections.Users
	userCollection.Email = userEntity.Email
	userCollection.Username = userEntity.Username
	userCollection.Lastname = userEntity.LastName
	userCollection.Password = userEntity.Password
	userCollection.Rol = userEntity.Rol

	return userCollection
}
