package parser

import (
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/database/collections"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
)

type UserEntityParserGateWayInterfaces interface {
	CreateUserRequestToEntity(userRequest request.RegisterUser) entity.User
	LoginUserRequestToEntity(userRequest request.LoginUserRequest) entity.User
	UserCollectionToEntity(coll *collections.Users) entity.User
}

type UserEntityParserGateWay struct {
	UserEntityParserGateWayInterfaces
}

func NewUserEntityParserGateWay() UserEntityParserGateWayInterfaces {
	return &UserEntityParserGateWay{}
}

func (u UserEntityParserGateWay) CreateUserRequestToEntity(userRequest request.RegisterUser) entity.User {
	var user entity.User
	user.Email = userRequest.Email
	user.Username = userRequest.Username
	user.LastName = userRequest.LastName
	user.Password = userRequest.Password
	user.Rol = userRequest.Rol
	return user
}

func (u UserEntityParserGateWay) UserCollectionToEntity(coll *collections.Users) entity.User {
	var userEntity entity.User
	userEntity.Email = coll.Email
	userEntity.Username = coll.Username
	userEntity.LastName = coll.Lastname
	userEntity.Password = coll.Password
	userEntity.Rol = coll.Rol
	return userEntity
}

func (u UserEntityParserGateWay) LoginUserRequestToEntity(userLoginRequest request.LoginUserRequest) entity.User {
	var user entity.User
	user.Email = userLoginRequest.Email
	user.Password = userLoginRequest.Password
	user.Username = ""
	user.LastName = ""
	user.Rol = 0
	return user
}
