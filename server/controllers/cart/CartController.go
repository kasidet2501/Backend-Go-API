package cart

import (
	"backend-ecom/configs"
	"backend-ecom/models"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var CartCollection *mongo.Collection = configs.GetCollection(configs.DB,"carts")
var ProductCollection *mongo.Collection = configs.GetCollection(configs.DB,"products")


func GetCarts(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var carts []models.CartItem
	cursor, err := CartCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer cursor.Close(ctx)

	// Iterate over the cursor and decode each document into the User struct
	for cursor.Next(ctx) {
		var cart models.CartItem
		if err := cursor.Decode(&cart); err != nil {
			log.Fatal(err.Error())
		}
		carts = append(carts, cart)
	}

	jsonProducts, err := json.Marshal(carts)

	return c.Status(200).SendString(string(jsonProducts))
}

// func AddToCart(c *fiber.Ctx) error{
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	var cart models.CartItem

// 	//validate the request body
// 	if err:= c.BodyParser(&cart); err != nil{
// 		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
// 			Status: http.StatusBadRequest, 
// 			Message:  "error", 
// 			Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	ProductPrice,err := CheckPrice(cart.ProductID,c);
// 	if err != nil{
// 		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
// 			Status: http.StatusInternalServerError, 
// 			Message:  "error", 
// 			Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	// Check to see if that product is in the cart.
// 	if err = CartCollection.FindOne(ctx, bson.M{"productId": cart.ProductID}).Decode(&cart); err == nil {
// 		err := increaseQuantity(cart.ProductID, ProductPrice, c)
// 		if err != nil{
// 			return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
// 				Status: http.StatusInternalServerError, 
// 				Message:  "error", 
// 				Data: &fiber.Map{"data": err.Error()}})
// 		}
// 	}else{ //กรณีเพิ่ม product ใหม่เข้าตระกร้า
// 		// กำหนดเงื่อนไขใน BSON filter
// 		filter := bson.M{
// 			"productId": cart.ProductID,
// 			"quantity" : 1,
// 			"price"	   : math.Round(ProductPrice*100)/100,
// 		}

// 		// Add product to cart
// 		if _ , err := CartCollection.InsertOne(ctx, filter); err != nil{
// 			return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
// 				Status: http.StatusInternalServerError, 
// 				Message:  "error", 
// 				Data: &fiber.Map{"data": err.Error()}})
// 		}
// 	}


// 	return c.Status(http.StatusOK).JSON(responses.ResponseData{
// 		Status: http.StatusOK, 
// 		Message:  "OK", 
// 		Data: &fiber.Map{"data": "Add product to cart successful"}})
// }

// func CheckPrice(productId string, c *fiber.Ctx) (float64, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	//find one product for get quantity
// 	ObjId,_ := primitive.ObjectIDFromHex(productId)
// 	var product models.Product
// 	if err := ProductCollection.FindOne(ctx, bson.M{"id": ObjId}).Decode(&product); err!= nil{
// 		return 0 , c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
// 			Status: http.StatusInternalServerError, 
// 			Message:  "error", 
// 			Data: &fiber.Map{"data": err.Error()}})
// 	}
// 	return float64(product.Price), nil
// }

// func DeleteItem(c *fiber.Ctx) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	CartId := c.Params("id")
// 	defer cancel()

// 	result,err := CartCollection.DeleteOne(ctx, bson.M{"productId" : CartId})
// 	if err != nil{
// 		return c.Status(http.StatusNotFound).JSON(responses.ResponseData{
// 			Status: http.StatusInternalServerError, 
// 			Message:  "error", 
// 			Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	if result.DeletedCount < 1 {
// 		messageError := "Cart Item ID : "+CartId+" not found!"
// 		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
// 			Status: http.StatusInternalServerError, 
// 			Message:  "error", 
// 			Data: &fiber.Map{"data": messageError}})
// 	}

// 	messageSuccess := "Cart Item ID : " + CartId + " successfully deleted!"
// 	return c.Status(http.StatusOK).JSON(responses.ResponseData{
// 		Status: http.StatusOK, 
// 		Message:  "success", 
// 		Data: &fiber.Map{"data": messageSuccess}})

// }

// func increaseQuantity(productId string,price float64, c *fiber.Ctx) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	//find one product for get quantity
// 	var result models.CartItem
// 	err := CartCollection.FindOne(ctx, bson.M{"productId": productId}).Decode(&result)
// 	if err != nil{
// 		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
// 			Status: http.StatusInternalServerError, 
// 			Message:  "error", 
// 			Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	newQuantity := result.Quantity + 1
	
// 	// กำหนดเงื่อนไขใน BSON filter
// 	update := bson.M{
// 		"productId": productId,
// 		"quantity" : newQuantity,
// 		"price"    : math.Round((price * float64(newQuantity))*100)/100,
// 	}

// 	if _, err := CartCollection.UpdateOne(ctx, bson.M{"productId" : productId}, bson.M{"$set": update}); err != nil{
// 		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
// 			Status: http.StatusInternalServerError, 
// 			Message:  "error", 
// 			Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	return nil
// }

// func ReduceQuantity(c *fiber.Ctx) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	var cart models.CartItem

// 	//validate the request body
// 	if err:= c.BodyParser(&cart); err != nil{
// 		return c.Status(http.StatusBadRequest).JSON(responses.ResponseData{
// 			Status: http.StatusBadRequest, 
// 			Message:  "error", 
// 			Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	ProductPrice,err := CheckPrice(cart.ProductID,c);
// 	if err != nil{
// 		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
// 			Status: http.StatusInternalServerError, 
// 			Message:  "error", 
// 			Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	//find one product for get old quantity
// 	var result models.CartItem
// 	if err = CartCollection.FindOne(ctx, bson.M{"productId": cart.ProductID}).Decode(&result); err != nil{
// 		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
// 			Status: http.StatusInternalServerError, 
// 			Message:  "error", 
// 			Data: &fiber.Map{"data": err.Error()}})
// 	}

// 	newQuantity := result.Quantity - 1

// 	// กำหนดเงื่อนไขใน BSON filter
// 	update := bson.M{}
// 	if(result.Quantity > 1){
// 		update = bson.M{
// 			"productId": cart.ProductID,
// 			"quantity" : newQuantity,
// 			"price"    : math.Round((ProductPrice * float64(newQuantity))*100)/100,
// 		}
// 	}else if(result.Quantity == 1){
// 		result,err := CartCollection.DeleteOne(ctx, bson.M{"productId" : cart.ProductID})
// 		if err != nil{
// 			return c.Status(http.StatusNotFound).JSON(responses.ResponseData{
// 				Status: http.StatusInternalServerError, 
// 				Message:  "error", 
// 				Data: &fiber.Map{"data": err.Error()}})
// 		}

// 		if result.DeletedCount < 1 {
// 			return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
// 				Status: http.StatusInternalServerError, 
// 				Message:  "error", 
// 				Data: &fiber.Map{"data": "Product in cart with specified ID not found!"}})
// 		}
// 	}

// 	if _, err := CartCollection.UpdateOne(ctx, bson.M{"productId" : cart.ProductID}, bson.M{"$set": update}); err != nil{
// 		return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
// 			Status: http.StatusInternalServerError, 
// 			Message:  "error", 
// 			Data: &fiber.Map{"data": err.Error()}})
// 	}


// 	return c.Status(http.StatusOK).JSON(responses.ResponseData{
// 		Status: http.StatusOK, 
// 		Message:  "OK", 
// 		Data: &fiber.Map{"data": "Reduce quantity product successful"}})

// }