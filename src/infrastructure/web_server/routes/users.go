package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
	"vettel-backend-app/src/application/usescases/usecase_users"
	"vettel-backend-app/src/infrastructure/database/mongo_client"
	"vettel-backend-app/src/infrastructure/database/user_repository"
	"vettel-backend-app/src/infrastructure/web_server/models/request"
	"vettel-backend-app/src/interfaces/controllers/users"
	"vettel-backend-app/src/interfaces/geteways/user_repository_gateway"
	"vettel-backend-app/src/interfaces/geteways/user_repository_gateway/parser"
)

type ResponseMsg struct {
	Msg string
}

func SingIn(route *echo.Echo) *echo.Echo {
	route.GET("/login", func(c echo.Context) error {
		email := c.QueryParam("email")
		password := c.QueryParam("password")

		var loginRequest request.LoginUserRequest
		loginRequest.Email = email
		loginRequest.Password = password

		validate := validator.New()
		err := validate.Struct(loginRequest)
		if err != nil {
			msg := ResponseMsg{
				Msg: "INVALID_USER_AND_PASSWORD_REQUEST",
			}
			c.JSON(400, msg)
		}

		mongoClientInterfaces := mongo_client.NewMongoClientInterfaces()
		userRepositoryInterfaces := user_repository.NewUsersRepository(mongoClientInterfaces)
		userCollectionParserGetWayInterfaces := parser.NewUserCollectionParserGateWay()
		userEntityParserGateWay := parser.NewUserEntityParserGateWay()
		userRepositoryGateWayInterfaces := user_repository_gateway.NewUserRepositoryGateWay(userRepositoryInterfaces, userCollectionParserGetWayInterfaces, userEntityParserGateWay)
		loginUserUseCase := usecase_users.NewLoginUserUseCase(userRepositoryGateWayInterfaces)
		singInController := users.NewSingInController(loginUserUseCase, userEntityParserGateWay)

		err, loginResult := singInController.SingIn(c.Request().Context(), loginRequest)
		if err != nil {
			msg := ResponseMsg{
				Msg: "UNAUTHORIZED",
			}
			c.JSON(401, msg)
			return nil
		}

		c.JSON(loginResult.Code, loginResult.Body)

		return nil
	})
	return route
}

func SingUp(route *echo.Echo) *echo.Echo {
	route.POST("register-user", func(c echo.Context) error {
		mongoClientInterfaces := mongo_client.NewMongoClientInterfaces()
		userRepositoryInterfaces := user_repository.NewUsersRepository(mongoClientInterfaces)
		userCollectionParserGetWayInterfaces := parser.NewUserCollectionParserGateWay()
		userEntityParserGateWay := parser.NewUserEntityParserGateWay()
		userRepository := user_repository_gateway.NewUserRepositoryGateWay(userRepositoryInterfaces, userCollectionParserGetWayInterfaces, userEntityParserGateWay)
		createNewUserUseCase := usecase_users.NewCreateNewUserUseCase(userRepository)
		singUpController := users.NewSingUpController(createNewUserUseCase, userEntityParserGateWay)

		var userRequest request.RegisterUser
		userRequest.Email = c.QueryParam("email")
		userRequest.Username = c.QueryParam("username")
		userRequest.LastName = c.QueryParam("lastname")
		userRequest.Password = c.QueryParam("password")
		rolParam, err := strconv.Atoi(c.QueryParam("rol"))
		if err != nil {
			c.JSON(400, err)
		}
		userRequest.Rol = rolParam

		validate := validator.New()
		err = validate.Struct(userRequest)
		if err != nil {
			msg := ResponseMsg{
				Msg: "BAD_REQUEST",
			}
			c.JSON(400, msg)
		}

		err, userCreated := singUpController.RegisterUser(c.Request().Context(), userRequest)
		if err != nil {
			msg := ResponseMsg{
				Msg: err.Error(),
			}
			log.Fatal(err)
			c.JSON(403, msg)
			return err
		}

		c.JSON(userCreated.Code, userCreated.Body)
		return nil
	})

	return route
}
