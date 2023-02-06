package users

import (
	"context"
	"fmt"
	"vettel-backend-app/src/application/usescases/usecase_users"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
	"vettel-backend-app/src/interfaces/geteways/user_repository_gateway/parser"
)

type UsersControllerInterfaces interface {
	SingIn(ctx context.Context, userLogin request.LoginUserRequest) (error, *response.RegisterUserResponse)
}

type SingInController struct {
	loginUserUseCaseInterface usecase_users.LoginUserUseCaseInterface
	userEntityParserGateWay   parser.UserEntityParserGateWayInterfaces
}

func NewSingInController(
	loginUserUseCaseInterface usecase_users.LoginUserUseCaseInterface,
	userEntityParserGateWay parser.UserEntityParserGateWayInterfaces,
) UsersControllerInterfaces {
	return &SingInController{loginUserUseCaseInterface, userEntityParserGateWay}
}

func (c *SingInController) SingIn(ctx context.Context, userLogin request.LoginUserRequest) (error, *response.RegisterUserResponse) {
	var userResponse response.RegisterUserResponse

	userEntity := c.userEntityParserGateWay.LoginUserRequestToEntity(userLogin)
	err, _ := c.loginUserUseCaseInterface.LoginUserUseCase(ctx, userEntity)
	if err != nil {
		fmt.Printf("Error try login for email: %s and password: %s \n with error: %s", userEntity.Email, userEntity.Password, err.Error())
		userResponse.Code = 500
		userResponse.Body = err.Error()
		return err, nil
	}
	userResponse.Code = 200
	userResponse.Body = userEntity

	return nil, &userResponse
}
