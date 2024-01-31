package main

import (
	"backend-ecom/configs"
	"backend-ecom/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	routes.RegisterRoute(app)
	routes.LoginRoute(app)
	routes.LogoutRoute(app)
	routes.ProductsRoute(app)
	routes.UsersRoute(app)
	routes.CartRoute(app)
	routes.OrderRoute(app)
	

    app.Listen(":6000")
}