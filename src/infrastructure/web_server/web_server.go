package web_server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"vettel-backend-app/src/infrastructure/web_server/routes"
)

func Start() {
	e := echo.New()
	e.Use(middleware.CORS())
	e = loadRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func loadRoutes(server *echo.Echo) *echo.Echo {
	server = routes.SingIn(server)
	server = routes.SingUp(server)
	server = routes.ListOrdersMonth(server)
	server = routes.LoadInventory(server)
	server = routes.AddItemInventory(server)
	server = routes.UpdateItemInventory(server)
	server = routes.DeleteInventoryItem(server)
	server = routes.LoadAllProducts(server)
	server = routes.CreateProduct(server)
	server = routes.UpdateProduct(server)
	server = routes.DeleteProduct(server)
	server = routes.AddDispatch(server)
	server = routes.EditOrder(server)
	server = routes.DeleteOrder(server)
	server = routes.ProcessPayment(server)
	return server
}
