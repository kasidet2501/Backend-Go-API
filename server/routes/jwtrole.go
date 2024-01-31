package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// UserData represents the user data extracted from the JWT token
type UserData struct {
	Username string
	Role  string
}
  
// userContextKey is the key used to store user data in the Fiber context
const userContextKey = "user"

// isAdmin checks if the user is an admin
func isAdmin(c *fiber.Ctx) error {

	if CheckRole(c) != "admin" {
		return fiber.ErrUnauthorized
	  }
	
	  return c.Next()
}

// isAdmin checks if the user is a User
func isUser(c *fiber.Ctx) error {

	if CheckRole(c) != "user" {
	  return fiber.ErrUnauthorized
	}
  
	return c.Next()
}

func CheckRole(c *fiber.Ctx) string {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		log.Fatal("Unauthorized1")
	}

	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

	// ตรวจสอบว่าการ Parse สำเร็จหรือไม่
	if err != nil {
		log.Fatal("Unauthorized2")
	}

	var user UserData

	claims := token.Claims.(jwt.MapClaims)
    user.Role = claims["role"].(string)
	fmt.Println( claims)

	return user.Role
}


// extractUserFromJWT is a middleware that extracts user data from the JWT token
func ExtractUserFromJWT(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	secret := os.Getenv("JWT_SECRET")

	_ , err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
  

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	
	// fmt.Println(token)

	return c.Next()

}
// func extractUserFromJWT(c *fiber.Ctx) error {
// 	user := &UserData{}

// 	// Extract the token from the Fiber context (inserted by the JWT middleware)
// 	token := c.Locals("user").(*jwt.Token)
// 	claims := token.Claims.(jwt.MapClaims)

// 	fmt.Println(claims)

// 	user.Username = claims["username"].(string)
// 	user.Role = claims["role"].(string)

// 	// Store the user data in the Fiber context
// 	c.Locals(userContextKey, user)

// 	return c.Next()
// }