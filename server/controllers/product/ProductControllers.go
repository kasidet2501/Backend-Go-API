package product

import (
	"backend-ecom/configs"
	"backend-ecom/models"
	"backend-ecom/responses"
	"context"
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
var CartCollection *mongo.Collection = configs.GetCollection(configs.DB,"carts")

var validate = validator.New()

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
	
	// Save file image
	file,err := c.FormFile("image")

	if err != nil{
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
			Status: http.StatusBadRequest, 
			Message:  "error", 
			Data: &fiber.Map{"data": err.Error()}})
	}

	// Generate a new random filename
	newFilename := uuid.New().String() + filepath.Ext(file.Filename)
	var fullPath = "../public/uploads/" + newFilename

	errSaveFile := 	c.SaveFile(file,fullPath)
	if errSaveFile != nil{
		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
			Status: http.StatusInternalServerError, 
			Message:  "error", 
			Data: &fiber.Map{"data": errSaveFile.Error()}})
	}

	product.Image = fullPath
	// End save file image

	
	// Test save Image

	// full,err := Uploadfile(c);
	// fmt.Print(full)

	// End Test save Image


	//use the validate library to validate required fields
	if validationErr := validate.Struct(&product); validationErr != nil{
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
			Status: http.StatusBadRequest, 
			Message:  "error", 
			Data: &fiber.Map{"data": validationErr.Error()}})
	}


	newProduct := models.Product{
		Id:			primitive.NewObjectID(),
		Title: 		product.Title,		
		Price: 		product.Price,		
		Description: product.Description,
		Category: 	product.Category,
		Image: 		fullPath,
	}

	result, err := ProductCollection.InsertOne(ctx, newProduct)
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


func GetProduct (c *fiber.Ctx) error{
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
	return c.Status(http.StatusOK).JSON(responses.ResponseData{
		Status: http.StatusCreated, 
		Message:  "success", 
		Data: &fiber.Map{"data": product}})
}

func EditProduct (c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ProductId := c.Params("id")
	var product models.Product
	defer cancel()

	ObjId,_ := primitive.ObjectIDFromHex(ProductId)

	if err := c.BodyParser(&product) ; err != nil{
		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
			Status: http.StatusCreated, 
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

func DeleteProduct(c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ProductId := c.Params("id")
	// var product models.Product
	defer cancel()

	ObjId,_ := primitive.ObjectIDFromHex(ProductId)

	// Delete image from folder

	// err := ProductCollection.FindOne(ctx, bson.M{"id" : ObjId}).Decode(&product)

	// if err != nil{
	// 	return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
	// 		Status: http.StatusCreated, 
	// 		Message:  "error", 
	// 		Data: &fiber.Map{"data": err.Error()}})
	
	// }

	// // ตรวจสอบว่าไฟล์หรือโฟลเดอร์นั้นมีอยู่จริงหรือไม่
	// if _, err := os.Stat(product.Image); os.IsNotExist(err) {
	// 	// ถ้าไม่มี, ส่ง HTTP response กลับไปว่าไม่พบไฟล์หรือโฟลเดอร์
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	// 		"error": "File or directory not found",
	// 	})
	// }

	// // ลบไฟล์หรือโฟลเดอร์
	// // err := os.Remove(product.Image)
	// if err = os.Remove(product.Image); err != nil {
	// 	// ถ้าเกิดข้อผิดพลาดในการลบ, ส่ง HTTP response กลับไปว่าเกิดข้อผิดพลาด
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": fmt.Sprintf("Error deleting file: %s", err.Error()),
	// 	})
	// }

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

	return c.Status(http.StatusOK).JSON(responses.ResponseData{
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

