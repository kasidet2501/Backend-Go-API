package routes

import (
	"encoding/json"
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

	user := c.Locals(userContextKey).(*UserData)
	if user.Role != "admin"{
		return fiber.ErrUnauthorized
	}

	return c.Next()
}

// isAdmin checks if the user is a User
func isUser(c *fiber.Ctx) error {
	user := c.Locals(userContextKey).(*UserData)
	if user.Role != "user"{
		return fiber.ErrUnauthorized
	}
  
	return c.Next()
}

func ExtractRoleFromJWT(c *fiber.Ctx) error {
	user := c.Locals(userContextKey).(*UserData)
	role := user.Role

	jsonRole, err := json.Marshal(role)
	if err != nil {
		return c.Status(400).SendString(string(err.Error()))
	}

	return c.Status(200).SendString(string(jsonRole))
}

func ExtractUsernameFromJWT(c *fiber.Ctx) error {
	user := c.Locals(userContextKey).(*UserData)
	username := user.Username

	jsonRole, err := json.Marshal(username)
	if err != nil {
		return c.Status(400).SendString(string(err.Error()))
	}

	return c.Status(200).SendString(string(jsonRole))
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

	// var user UserData

	claims := token.Claims.(jwt.MapClaims)
    // user.Role = claims["role"].(string)

	return claims["role"].(string)
}




// extractUserFromJWT is a middleware that extracts user data from the JWT token
func ExtractUserFromJWT(c *fiber.Ctx) error {
	user := &UserData{}

	// Extract the token from the Fiber context (inserted by the JWT middleware)
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	fmt.Println("Cookies : " ,claims)

	user.Username = claims["username"].(string)
	user.Role = claims["role"].(string)

	// Store the user data in the Fiber context
	c.Locals(userContextKey, user)

	return c.Next()
}