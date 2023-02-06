package usecase_users

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/interfaces/geteways/user_repository_gateway"
)

type CreateNewUserUseCaseInterfaces interface {
	CreateNewUSerUseCase(ctx context.Context, userEntity entity.User) (error, *entity.User)
}

type CreateNewUserUseCase struct {
	userRepositoryGateWayInterfaces user_repository_gateway.UserRepositoryGateWayInterfaces
}

func NewCreateNewUserUseCase(userRepositoryGateWayInterfaces user_repository_gateway.UserRepositoryGateWayInterfaces) CreateNewUserUseCaseInterfaces {
	return &CreateNewUserUseCase{userRepositoryGateWayInterfaces}
}

func (c CreateNewUserUseCase) CreateNewUSerUseCase(ctx context.Context, userEntity entity.User) (error, *entity.User) {
	fmt.Printf("Validating if email already exist to username %s\n", userEntity.Username)
	err, user := c.userRepositoryGateWayInterfaces.FindByEmail(ctx, userEntity)
	if err != nil {
		return err, nil
	}
	if user != nil {
		return errors.New("USER_ALREADY_EXIST"), nil
	}
	err, userEntityCreated := c.userRepositoryGateWayInterfaces.CreateUser(ctx, userEntity)

	if err != nil {
		return err, nil
	}

	fmt.Printf("User created successfully with email: %s\n", userEntity.Email)

	return nil, userEntityCreated
}
