package cart

import (
	"backend-ecom/configs"
	"backend-ecom/models"
	"backend-ecom/responses"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserData struct {
	Username string
	Role  string
}

var OrderCollection *mongo.Collection = configs.GetCollection(configs.DB,"orders")

func AddToCart(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// var cart []models.CartItem


	// ดึงข้อมูลผู้ใช้จาก JWT Token
	// token := c.Locals("user").(*jwt.Token)
	// claims := token.Claims.(jwt.MapClaims)

	// user := c.Locals("user").(*UserData)
	// username := user.Username
	// username := claims["username"].(string)


	// ดึงข้อมูลสินค้าที่ต้องการเพิ่ม
	var newItem models.CartItem
	if err := c.BodyParser(&newItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// สร้าง Order หรือดึง Order ที่มีอยู่แล้ว
	filter := bson.M{"username": "user"}
	var existingOrder models.Order
	err := OrderCollection.FindOne(ctx, filter).Decode(&existingOrder)
	if err == mongo.ErrNoDocuments {
		// กรณีไม่พบ Order ใน MongoDB
		existingOrder = models.Order{
			Id: primitive.NewObjectID(),
			Username:  "user",
			Carts: map[string]models.CartItem{ newItem.ProductID : newItem},
		}
		_, err := OrderCollection.InsertOne(ctx, existingOrder)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add item to cart"})
		}
	} else if err != nil {
		// กรณีเกิด error อื่น ๆ
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch order"})
	} else {
		// กรณีมี Order อยู่แล้ว
		// เช็คก่อนว่า user นั้นมี items ชิ้นนั้นอยู่หรือไม่
		check := CheckItem("user", newItem.ProductID)

		if (check){
			fmt.Println("case : true")

			// กำหนดเงื่อนไขใน BSON filter
			filter := bson.M{
				"username": "user",
			}

			// กำหนดการอัปเดต quantity ใน BSON update
			update := bson.M{
				"$set": bson.M{
					// ในกรณีนี้, เราจะใช้ `$set` เพื่อระบุ `carts.productId` เป็น key และ `newCartItem` เป็นค่า
					fmt.Sprintf("carts.%s", newItem.ProductID): bson.M{
						"productId" : newItem.ProductID,
						"quantity" : newItem.Quantity,
					},
				},
			}
			_,err := OrderCollection.UpdateOne(ctx, filter, update)
			// _,err := OrderCollection.UpdateOne(ctx, bson.M{"username" : "user"}, bson.M{"$set": update})

			if err != nil{
				return c.Status(http.StatusInternalServerError).JSON(responses.ResponseData{
					Status: http.StatusInternalServerError, 
					Message:  "error", 
					Data: &fiber.Map{"data": err.Error()}})
			}
		}else{
			// fmt.Println("case : false")
			// // กรณีมี Order อยู่แล้ว		
			// existingOrder.Carts = append(existingOrder.Carts, newItem)
			// _, err := OrderCollection.ReplaceOne(context.Background(), filter, existingOrder)
			// if err != nil {
			// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add item to cart"})
			// }
		}
	}
	return c.JSON(fiber.Map{"message": "Item added to cart successfully"})
}

func CheckItem(user string, productId string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// กำหนดเงื่อนไขใน BSON filter
	filter := bson.M{
		"username": user,
		"carts"  : bson.M{
			"productId" : productId,
		},
	}

	// filter := bson.M{
	// 	"username": user,
	// 	"carts": bson.M{
	// 		"productId": productid,
	// 	},
	// }

	fmt.Println(productId)

	var result models.Order
	err := OrderCollection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// ไม่มี item อยู่
		fmt.Println(err)
		return  false
	} else if err != nil {
		log.Fatal(err)
	} else {
		// มี item อยู่
		return true
	}
	fmt.Println(&fiber.Map{"result": result})
	return false
}

