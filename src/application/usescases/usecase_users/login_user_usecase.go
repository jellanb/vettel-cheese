package usecase_users

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"vettel-backend-app/src/domain/entity"
	"vettel-backend-app/src/interfaces/geteways/user_repository_gateway"
)

type LoginUserUseCaseInterface interface {
	LoginUserUseCase(ctx context.Context, user entity.User) (error, *entity.User)
}

type LoginUserUseCase struct {
	userRepositoryGateWayInterfaces user_repository_gateway.UserRepositoryGateWayInterfaces
}

func NewLoginUserUseCase(userRepositoryGateWayInterfaces user_repository_gateway.UserRepositoryGateWayInterfaces) LoginUserUseCaseInterface {
	return &LoginUserUseCase{userRepositoryGateWayInterfaces}
}

func (l LoginUserUseCase) LoginUserUseCase(ctx context.Context, user entity.User) (error, *entity.User) {
	fmt.Printf("Find user to email: %s and password: %s \n", user.Email, user.Password)
	err, userResult := l.userRepositoryGateWayInterfaces.FindByEmailAndPassword(ctx, user)
	if err != nil {
		return err, nil
	}
	if userResult == nil {
		return errors.New("User not found, invalid email and password"), nil
	}
	fmt.Printf("User found to email: %s and password: %s \n", user.Email, user.Password)
	return nil, userResult
}
