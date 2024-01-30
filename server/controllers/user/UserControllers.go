package user

import (
	"backend-ecom/configs"
	"backend-ecom/models"
	"backend-ecom/responses"
	"context"
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

// -----------------> Get  All users <----------------- //
func GetUser(c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.User
	defer cancel()

	// Find all documents in the collection
	cursor, err := UserCollection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode each document into the User struct
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return err
		}
		users = append(users, user)
	}


	return c.Status(http.StatusOK).JSON(responses.ResponseData{
		Status: http.StatusOK, 
		Message:  "success", 
		Data: &fiber.Map{"data": users}})
	 
}

// -----------------> Get User By ID <----------------- //
func GetUserByID(c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	UserId := c.Params("id")
	var user models.User
	defer cancel()

	ObjId,_ := primitive.ObjectIDFromHex(UserId)

	// err := UserCollection.FindOne(ctx, bson.M{"username" : UserId}).Decode(&user)
	err := UserCollection.FindOne(ctx, bson.M{"id" : ObjId}).Decode(&user)

	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
			Status: http.StatusInternalServerError, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	
	}
	return c.Status(http.StatusOK).JSON(responses.ResponseData{
		Status: http.StatusCreated, 
		Message:  "success", 
		Data: &fiber.Map{"data": user}})
	
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


// -----------------> Create User <----------------- //
func CreateUser(c *fiber.Ctx) error{
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

	return c.Status(http.StatusCreated).JSON(responses.ResponseData{
		Status: http.StatusCreated, 
		Message:  "success", 
		Data: &fiber.Map{"data": result}})
	
}

// -----------------> Update User <----------------- //
func EditUser(c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	UserId := c.Params("id")
	var user models.User
	defer cancel()

	ObjId,_ := primitive.ObjectIDFromHex(UserId)

	if err := c.BodyParser(&user) ; err != nil{
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
			Status: http.StatusBadRequest, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validate library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil{
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
			Status: http.StatusBadRequest, 
			Message:  "error", 
			Data: &fiber.Map{"data": validationErr.Error()}})
	}

	//new feature : Check Username and password if they are duplicate in another user then return a badrequest

	update := bson.M{
		"role" : user.Role,
		"email": user.Email,
		"username": user.Username,
		"password": user.Password,
		"name": bson.M{
			"firstname": user.Name.Firstname,
			"lastname": user.Name.Lastname,
		},
		"phone": user.Phone,
	}

	result,err := UserCollection.UpdateOne(ctx, bson.M{"id" : ObjId}, bson.M{"$set": update})

	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
			Status: http.StatusInternalServerError, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated product details
	var UpdateUser models.User
	if result.MatchedCount == 1 {
        err := UserCollection.FindOne(ctx, bson.M{"id": ObjId}).Decode(&UpdateUser)

        if err != nil {
            return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
				Status: http.StatusInternalServerError, 
				Message: "error", 
				Data: &fiber.Map{"data": err.Error()}})
        }
    }
	return c.Status(http.StatusOK).JSON(responses.ResponseData{
		Status: http.StatusOK,
		Message: "success",
		Data: &fiber.Map{"data" : UpdateUser},
	})
	
}


// -----------------> Delete User <----------------- //
func DeleteUser(c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	UserId := c.Params("id")
	defer cancel()

	ObjId,_ := primitive.ObjectIDFromHex(UserId)

	result,err := UserCollection.DeleteOne(ctx, bson.M{"id" : ObjId})
	if err != nil{
		return c.Status(http.StatusNotFound).JSON(responses.ResponseData{
			Status: http.StatusInternalServerError, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
			Status: http.StatusInternalServerError, 
			Message:  "error", 
			Data: &fiber.Map{"data": "User with specified ID not found!"}})
	}

	return c.Status(http.StatusOK).JSON(responses.ResponseData{
		Status: http.StatusOK, 
		Message:  "success", 
		Data: &fiber.Map{"data": "User successfully deleted!"}})
}
