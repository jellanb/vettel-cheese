package users

import (
	"context"
	"fmt"
	"vettel-backend-app/src/application/usescases/usecase_users"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/infrastructure/web_server/models/response"
	"vettel-backend-app/src/interfaces/geteways/user_repository_gateway/parser"
)

type SingUpControllerInterfaces interface {
	RegisterUser(ctx context.Context, userRequest request.RegisterUser) (error, *response.RegisterUserResponse)
}

type SingUpController struct {
	createNewUserUseCaseInterfaces usecase_users.CreateNewUserUseCaseInterfaces
	userEntityParserGateWay        parser.UserEntityParserGateWayInterfaces
}

type UserDetails struct {
	Username string
	Password string
}

func NewSingUpController(
	createNewUserUseCase usecase_users.CreateNewUserUseCaseInterfaces,
	userEntityParserGateWay parser.UserEntityParserGateWayInterfaces,
) SingUpControllerInterfaces {
	return &SingUpController{
		createNewUserUseCase,
		userEntityParserGateWay,
	}
}

func (c *SingUpController) RegisterUser(ctx context.Context, userRequest request.RegisterUser) (error, *response.RegisterUserResponse) {

	var userResponse response.RegisterUserResponse

	userEntity := c.userEntityParserGateWay.CreateUserRequestToEntity(userRequest)
	err, _ := c.createNewUserUseCaseInterfaces.CreateNewUSerUseCase(ctx, userEntity)
	if err != nil {
		fmt.Printf("Error creating new user with email: %s, password: %s and error: %s", userRequest.Email, userRequest.Password, err.Error())
		userResponse.Code = 500
		userResponse.Body = err.Error()
		return err, nil
	}

	userResponse.Code = 200
	userResponse.Body = userEntity

	return nil, &userResponse
}
