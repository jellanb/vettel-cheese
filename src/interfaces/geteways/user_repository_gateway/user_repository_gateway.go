package user_repository_gateway

import (
	"context"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/infrastructure/database/user_repository"
	"vettel-backend-app/src/interfaces/geteways/user_repository_gateway/parser"
)

type UserRepositoryGateWayInterfaces interface {
	FindByEmail(ctx context.Context, userEntity entity.User) (error, *entity.User)
	CreateUser(ctx context.Context, userEntity entity.User) (error, *entity.User)
	FindByEmailAndPassword(ctx context.Context, userEntity entity.User) (error, *entity.User)
}

type UserRepositoryGateWay struct {
	userRepositoryInterfaces             user_repository.UsersRepositoryInterfaces
	userCollectionParserGetWayInterfaces parser.UserCollectionParserGetWayInterfaces
	userEntityParserGateWay              parser.UserEntityParserGateWayInterfaces
}

func NewUserRepositoryGateWay(
	userRepositoryInterfaces user_repository.UsersRepositoryInterfaces,
	userCollectionParserGetWayInterfaces parser.UserCollectionParserGetWayInterfaces,
	userEntityParserGateWay parser.UserEntityParserGateWayInterfaces,
) UserRepositoryGateWayInterfaces {
	return &UserRepositoryGateWay{userRepositoryInterfaces, userCollectionParserGetWayInterfaces, userEntityParserGateWay}
}

func (u UserRepositoryGateWay) FindByEmail(ctx context.Context, userEntity entity.User) (error, *entity.User) {
	userCollection := u.userCollectionParserGetWayInterfaces.UserEntityToCollection(userEntity)
	err, userCollResult := u.userRepositoryInterfaces.FindUserByEmail(ctx, userCollection)
	if err != nil {
		return err, nil
	}
	if userCollResult == nil {
		return nil, nil
	}
	userEntityFound := u.userEntityParserGateWay.UserCollectionToEntity(userCollResult)

	return nil, &userEntityFound
}

func (u UserRepositoryGateWay) CreateUser(ctx context.Context, userEntity entity.User) (error, *entity.User) {
	userCollection := u.userCollectionParserGetWayInterfaces.UserEntityToCollection(userEntity)
	err, userCollResult := u.userRepositoryInterfaces.Register(ctx, userCollection)
	if err != nil {
		return err, nil
	}
	if userCollResult == nil {
		return nil, nil
	}
	newUser := u.userEntityParserGateWay.UserCollectionToEntity(userCollResult)

	return nil, &newUser
}

func (u UserRepositoryGateWay) FindByEmailAndPassword(ctx context.Context, userEntity entity.User) (error, *entity.User) {
	userCollection := u.userCollectionParserGetWayInterfaces.UserEntityToCollection(userEntity)
	err, userCollResult := u.userRepositoryInterfaces.FindUserByEmailAndPassword(ctx, userCollection)
	if err != nil {
		return err, nil
	}
	if userCollResult == nil {
		return nil, nil
	}
	userLogin := u.userEntityParserGateWay.UserCollectionToEntity(userCollResult)
	return nil, &userLogin
}
