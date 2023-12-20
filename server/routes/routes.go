package routes

import (
	"backend-ecom/controllers/product"

	"github.com/gofiber/fiber/v2"
)

func ProductsRoute(app *fiber.App){
	app.Get("/product/:id", product.GetProduct)
	app.Post("/createproduct", product.CreateProduct)
	app.Put("/updateproduct/:id", product.EditProduct)
	app.Delete("/deleteproduct/:id", product.DeleteProduct)
	// app.Post("/upload", product.Uploadfile)
}