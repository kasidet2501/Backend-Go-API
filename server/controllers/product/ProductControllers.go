package product

import (
	"backend-ecom/configs"
	"backend-ecom/models"
	"backend-ecom/responses"
	"context"
	"encoding/json"
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ProductCollection *mongo.Collection = configs.GetCollection(configs.DB,"products")
var UserCollection *mongo.Collection = configs.GetCollection(configs.DB,"users")
var OrderCollection *mongo.Collection = configs.GetCollection(configs.DB,"orders")

var validate = validator.New()

// -----------------> Create Product <----------------- //
func CreateProduct(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var product models.Product
	defer cancel()

	//validate the request body
	if err:= c.BodyParser(&product); err != nil{
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
			Status: http.StatusBadRequest, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	}
	
	// // Save file image
	// file,err := c.FormFile("image")

	// if err != nil{
	// 	return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
	// 		Status: http.StatusBadRequest, 
	// 		Message:  "error", 
	// 		Data: &fiber.Map{"data": err.Error()}})
	// }

	// // Generate a new random filename
	// newFilename := uuid.New().String() + filepath.Ext(file.Filename)
	// var fullPath = "../client/src/images/" + newFilename

	// errSaveFile := 	c.SaveFile(file,fullPath)
	// if errSaveFile != nil{
	// 	return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
	// 		Status: http.StatusInternalServerError, 
	// 		Message:  "error", 
	// 		Data: &fiber.Map{"data": errSaveFile.Error()}})
	// }

	// product.Image = fullPath
	// // End save file image

	
	// // Test save Image

	// // full,err := Uploadfile(c);
	// // fmt.Print(full)

	// // End Test save Image


	// //use the validate library to validate required fields
	// if validationErr := validate.Struct(&product); validationErr != nil{
	// 	return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
	// 		Status: http.StatusBadRequest, 
	// 		Message:  "error", 
	// 		Data: &fiber.Map{"data": validationErr.Error()}})
	// }


	newProduct := models.Product{
		Id:			primitive.NewObjectID(),
		Title: 		product.Title,		
		Price: 		product.Price,		
		Description: product.Description,
		Category: 	product.Category,
		Image: 		product.Image,
	}

	result, err := ProductCollection.InsertOne(ctx, newProduct)
	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
			Status: http.StatusInternalServerError, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(200).JSON(responses.ResponseData{
		Status: http.StatusCreated, 
		Message:  "success", 
		Data: &fiber.Map{"data": result}})
}


// -----------------> Get  All Products <----------------- //
func GetProduct(c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var products []models.Product
	defer cancel()

	// Find all documents in the collection
	cursor, err := ProductCollection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode each document into the Product struct
	for cursor.Next(ctx) {
		var product models.Product
		if err := cursor.Decode(&product); err != nil {
			return err
		}
		products = append(products, product)
	}

	jsonProducts, err := json.Marshal(products)

	return c.Status(200).SendString(string(jsonProducts))
	 
}

// -----------------> Get Product by ID <----------------- //
func GetProductById (c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ProductId := c.Params("id")
	var product models.Product
	defer cancel()

	ObjId,_ := primitive.ObjectIDFromHex(ProductId)

	err := ProductCollection.FindOne(ctx, bson.M{"id" : ObjId}).Decode(&product)

	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
			Status: http.StatusCreated, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	
	}
	jsonProduct, err := json.Marshal(product)

	return c.Status(200).SendString(string(jsonProduct))
}

// -----------------> Edit Product <----------------- //
func EditProduct (c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ProductId := c.Params("id")
	var product models.Product
	defer cancel()

	ObjId,_ := primitive.ObjectIDFromHex(ProductId)

	if err := c.BodyParser(&product) ; err != nil{
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
			Status: http.StatusBadRequest, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validate library to validate required fields
	if validationErr := validate.Struct(&product); validationErr != nil{
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
			Status: http.StatusBadRequest, 
			Message:  "error", 
			Data: &fiber.Map{"data": validationErr.Error()}})
	}

	update := bson.M{
		"title" : product.Title,
		"price" : product.Price,
		"description" : product.Description,
		"category" : product.Category,
		"image" : product.Image, 
	}

	result,err := ProductCollection.UpdateOne(ctx, bson.M{"id" : ObjId}, bson.M{"$set": update})

	if err != nil{
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
			Status: http.StatusInternalServerError, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	}

	//get updated product details
	var UpdateProduct models.Product
	if result.MatchedCount == 1 {
        err := ProductCollection.FindOne(ctx, bson.M{"id": ObjId}).Decode(&UpdateProduct)

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
		Data: &fiber.Map{"data" : UpdateProduct},
	})
}


// -----------------> Delete Product <----------------- //
func DeleteProduct(c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ProductId := c.Params("id")
	// var product models.Product
	defer cancel()

	ObjId,_ := primitive.ObjectIDFromHex(ProductId)


	result,err := ProductCollection.DeleteOne(ctx, bson.M{"id" : ObjId})
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
			Data: &fiber.Map{"data": "Product with specified ID not found!"}})
	}

	return c.Status(200).JSON(responses.ResponseData{
		Status: http.StatusOK, 
		Message:  "success", 
		Data: &fiber.Map{"data": "Product successfully deleted!"}})
}

func Uploadfile(c *fiber.Ctx) (string ,error){
	// Parse form data
	file,err := c.FormFile("image")

	if err != nil{
		return "",c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
			Status: http.StatusBadRequest, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	}

	// Generate a new random filename
	newFilename := uuid.New().String() + filepath.Ext(file.Filename)
	var fullPath = "../public/uploads/" + newFilename

	errSaveFile := 	c.SaveFile(file,fullPath)
	if errSaveFile != nil{
		return "",c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
			Status: http.StatusInternalServerError, 
			Message:  "error", 
			Data: &fiber.Map{"data": errSaveFile.Error()}})
	}

	// return fullPath
	return fullPath, c.SendString(fullPath);
}

