package authlogin

import (
	"backend-ecom/configs"
	"backend-ecom/models"
	"backend-ecom/responses"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = configs.GetCollection(configs.DB,"users")


func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func Login(secretKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		var request models.LoginRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
				Status: http.StatusBadRequest, 
				Message:  "error", 
				Data: &fiber.Map{"data": err.Error()}})
		}

		// Check credentials - In real world, you should check against a database

		err := UserCollection.FindOne(ctx, bson.M{"username" : request.Username}).Decode(&user)

		if err != nil{
			return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
				Status: http.StatusUnauthorized, 
				Message:  "error", 
				Data: &fiber.Map{"data": fiber.ErrUnauthorized}})
		
		}

		
		match := CheckPasswordHash(request.Password, user.Password)
		// fmt.Println(match)
		if ((request.Username != user.Username) || (match != true)) {
			return fiber.ErrUnauthorized
		}

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Username
		claims["role"] = user.Role
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	
		// Generate encoded token
		t, err := token.SignedString([]byte(secretKey))
		if err != nil {
		  return c.SendStatus(fiber.StatusInternalServerError)
		}

		// c.Cookie(&fiber.Cookie{
		// 	Name:     "jwt",
		// 	Value:    t,
		// 	Expires:  time.Now().Add(time.Hour * 24),
		// 	HTTPOnly: true,
		// })

		jsonJWT, err := json.Marshal(t)
		return c.Status(200).SendString(string(jsonJWT))
	}
}





