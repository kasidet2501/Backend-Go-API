package routes

import (
	"backend-ecom/controllers/authlogin"
	"backend-ecom/controllers/logout"
	"backend-ecom/controllers/order"
	"backend-ecom/controllers/product"
	"backend-ecom/controllers/register"
	"backend-ecom/controllers/user"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)


func ProductsRoute(app *fiber.App){

	// -------------------> Routes for Admin role <------------------------
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


	// -------------------> Routes for User role <------------------------
	// Group routes under /user
	// UserProductGroup := app.Group("/user")
	// // Apply the isUser middleware only to the /product routes
	// // UserProductGroup.Use(isUser)

	// UserProductGroup.Get("/product", product.GetProduct)
	// UserProductGroup.Get("/product/:id", product.GetProductById)
}
func UsersRoute(app *fiber.App){

	// Group routes under /user
	userGroup := app.Group("/user")

	userGroup.Get("/:user", user.GetUserByUsername)

	// Apply the isAdmin middleware only to the /user routes
	userGroup.Use(isAdmin)

	userGroup.Get("/", user.GetUser)
	userGroup.Get("/:id", user.GetUserByID)
	userGroup.Post("/", user.CreateUser) //for Admin if admin want to define role
	userGroup.Put("/:id", user.EditUser)
	userGroup.Delete("/:id", user.DeleteUser)
	// app.Post("/upload", product.Uploadfile)
}

func CartRoute(app *fiber.App){
	// // Group routes under /cart
	// CartGroup := app.Group("/cart")

	// // Apply the isUser middleware only to the /cart routes
	// CartGroup.Use(isUser)

	// CartGroup.Get("/", cart.GetCarts)
	// CartGroup.Post("/", cart.AddToCart)
	// CartGroup.Post("/reduce", cart.ReduceQuantity)
	// CartGroup.Delete("/:id", cart.DeleteItem)

}
func OrderRoute(app *fiber.App){
	// Group routes under /order
	OrderGroup := app.Group("/order")

	// Apply the isUser middleware only to the /order routes
	OrderGroup.Use(isUser)

	OrderGroup.Post("/", order.ConfirmOrder)
	// OrderGroup.Get("/pd", order.GetSumPrice)

}


func RegisterRoute(app *fiber.App){
	// Group routes under /register
	RegisterGroup := app.Group("/register")
	RegisterGroup.Post("/", register.Register)
}

func LoginRoute(app *fiber.App){

	app.Get("/user/product", product.GetProduct)
	app.Get("/user/product/:id", product.GetProductById)

	app.Post("/login", authlogin.Login(os.Getenv("JWT_SECRET")))

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
    }))

	app.Use(ExtractUserFromJWT)

	app.Get("/role", ExtractRoleFromJWT)
	app.Get("/username", ExtractUsernameFromJWT)
}

func LogoutRoute(app *fiber.App){
	// Group routes under /register
	RegisterGroup := app.Group("/logout")
	RegisterGroup.Get("/", logout.Logout)
}

