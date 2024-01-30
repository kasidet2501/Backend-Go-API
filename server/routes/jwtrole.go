package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// isAdmin checks if the user is an admin
func isAdmin(c *fiber.Ctx) error {
	user := c.Locals(userContextKey).(*UserData)
  
	if user.Role != "admin" {
	  return fiber.ErrUnauthorized
	}
  
	return c.Next()
}

// isAdmin checks if the user is a User
func isUser(c *fiber.Ctx) error {
	user := c.Locals(userContextKey).(*UserData)
  
	if user.Role != "user" {
	  return fiber.ErrUnauthorized
	}
  
	return c.Next()
}

// UserData represents the user data extracted from the JWT token
type UserData struct {
	Username string
	Role  string
}
  
// userContextKey is the key used to store user data in the Fiber context
const userContextKey = "user"

// extractUserFromJWT is a middleware that extracts user data from the JWT token
func extractUserFromJWT(c *fiber.Ctx) error {
	user := &UserData{}

	// Extract the token from the Fiber context (inserted by the JWT middleware)
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	fmt.Println(claims)

	user.Username = claims["username"].(string)
	user.Role = claims["role"].(string)

	// Store the user data in the Fiber context
	c.Locals(userContextKey, user)

	return c.Next()
}