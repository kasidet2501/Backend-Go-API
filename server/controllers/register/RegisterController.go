package register

import (
	"backend-ecom/configs"
	"backend-ecom/models"
	"backend-ecom/responses"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = configs.GetCollection(configs.DB,"users")
var validate = validator.New()

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func Register(c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	//validate the request body
	if err:= c.BodyParser(&user); err != nil{
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
			Status: http.StatusBadRequest, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	}

	user.Role = "user"

	// Example: Check if username is unique before inserting or updating
	err := UserCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&user)
	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
			Status: http.StatusBadRequest, 
			Message:  "error",
			Data: &fiber.Map{"data": "Username is already used"}})
	} 

	// Example: Check if email is unique before inserting or updating
	err = UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&user);
	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
			Status: http.StatusBadRequest, 
			Message:  "error",
			Data: &fiber.Map{"data": "Email is already used"}})
	} 


	//use the validate library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil{
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
			Status: http.StatusBadRequest, 
			Message:  "error", 
			Data: &fiber.Map{"data": validationErr.Error()}})
	}

	user.Id = primitive.NewObjectID()

	hashPass,err := HashPassword(user.Password)
	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
			Status: http.StatusInternalServerError, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	}
	user.Password = hashPass

	result, err := UserCollection.InsertOne(ctx, user)
	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
			Status: http.StatusInternalServerError, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	}
	json, err := json.Marshal(result)
	fmt.Println(result)
	return c.Status(200).SendString(string(json))
	// return c.Status(200).JSON(responses.ResponseData{
	// 	Status: http.StatusCreated, 
	// 	Message:  "success", 
	// 	Data: &fiber.Map{"data": result}})
}