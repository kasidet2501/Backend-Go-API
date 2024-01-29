package routes

import (
	"backend-ecom/controllers/product"
	"backend-ecom/controllers/user"

	"github.com/gofiber/fiber/v2"
)

func ProductsRoute(app *fiber.App){
	app.Get("/product/:id", product.GetProduct)
	app.Post("/createproduct", product.CreateProduct)
	app.Put("/updateproduct/:id", product.EditProduct)
	app.Delete("/deleteproduct/:id", product.DeleteProduct)
	// app.Post("/upload", product.Uploadfile)
}
func UsersRoute(app *fiber.App){
	app.Get("/user/:id", user.GetUser)
	app.Post("/createuser", user.CreateUser)
	app.Put("/updateuser/:id", user.EditUser)
	app.Delete("/deleteuser/:id", user.DeleteUser)
	// app.Post("/upload", product.Uploadfile)
}