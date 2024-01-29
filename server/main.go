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
	routes.ProductsRoute(app)
	routes.UsersRoute(app)

    // app.Get("/", func(c *fiber.Ctx) error {
    //     return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
    // })

    app.Listen(":6000")
}