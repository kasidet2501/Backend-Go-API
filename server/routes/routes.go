package routes

import (
	"backend-ecom/controllers/authlogin"
	"backend-ecom/controllers/product"
	"backend-ecom/controllers/user"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)


func ProductsRoute(app *fiber.App){
	// Group routes under /product
	ProductGroup := app.Group("/product")

	// Apply the isAdmin middleware only to the /product routes
	ProductGroup.Use(isAdmin)

	ProductGroup.Get("/", product.GetProduct)
	ProductGroup.Get("/:id", product.GetProductById)
	ProductGroup.Post("/", product.CreateProduct)
	ProductGroup.Put("/:id", product.EditProduct)
	ProductGroup.Delete("/:id", product.DeleteProduct)
	// app.Post("/upload", product.Uploadfile)
}
func UsersRoute(app *fiber.App){

	// Group routes under /user
	userGroup := app.Group("/user")

	// Apply the isAdmin middleware only to the /user routes
	userGroup.Use(isAdmin)

	userGroup.Get("/", user.GetUser)
	userGroup.Get("/:id", user.GetUserByID)
	userGroup.Post("/", user.CreateUser)
	userGroup.Put("/:id", user.EditUser)
	userGroup.Delete("/:id", user.DeleteUser)
	// app.Post("/upload", product.Uploadfile)
}

func LoginRoute(app *fiber.App){
	app.Post("/login", authlogin.Login(os.Getenv("JWT_SECRET")))

	secret := os.Getenv("JWT_SECRET");
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
	}))

	app.Use(extractUserFromJWT)

	
	
	// app.Post("/upload", product.Uploadfile)
}