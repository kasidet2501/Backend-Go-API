package order

import (
	"backend-ecom/configs"
	"backend-ecom/models"
	"context"
	"log"
	"os"
	"time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var CartCollection *mongo.Collection = configs.GetCollection(configs.DB,"carts")
var OrderCollection *mongo.Collection = configs.GetCollection(configs.DB,"orders")


// UserClaims คือ struct ที่ใช้เก็บข้อมูลใน JWT token
type UserClaims struct {
	// jwt.StandardClaims
	Username string `json:"username"`
	Role   string `json:"role"`
}

func GetUsername(c *fiber.Ctx) (string){
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

	var user models.UserData

	claims := token.Claims.(jwt.MapClaims)
    user.Username = claims["username"].(string)

	return user.Username

}


func ConfirmOrder(c *fiber.Ctx) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find all documents in the collection
	var carts []models.CartItem = getCarts()

	// get username from JWT toke
	usernameJWT := GetUsername(c)

	// สร้าง Order
	filter := bson.M{"username": usernameJWT}
	var order models.Order
	err := OrderCollection.FindOne(ctx, filter).Decode(&order)
	if err == mongo.ErrNoDocuments {
		// กรณีไม่พบ Order ใน MongoDB
		price:= GetSumPrice() // Sum total price
		order = models.Order{
			Id: primitive.NewObjectID(),
			Username:  usernameJWT,
			Carts: carts,
			Price: price,
		}
		_, err := OrderCollection.InsertOne(ctx, order)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to add item to cart"})
		}
	} else if err != nil {
		// กรณีเกิด error อื่น ๆ
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch order"})
	}
	
	return c.JSON(fiber.Map{"data": order})
}

func getCarts() []models.CartItem {
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
	return carts
}

func GetSumPrice() (float64) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ทำ aggregation สำหรับรวม price ทั้งหมด
	pipeline := bson.A{
		bson.M{
			"$group": bson.M{
				"_id":   nil,
				"total": bson.M{"$sum": "$price"},
			},
		},
	}

	// ทำ aggregation
	cursor, err := CartCollection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}

	// ปิด cursor เมื่อเสร็จสิ้น
	defer cursor.Close(context.Background())

	// สร้าง struct เพื่อเก็บผลลัพธ์
	var result struct {
		Total float64 `json:"total"`
	}

	// Decode ผลลัพธ์
	if cursor.Next(ctx) { //ตรวจสอบว่ามีเอกสารถัดไปในผลลัพธ์หรือไม่ 
		if err := cursor.Decode(&result); err != nil { //ดึงข้อมูลจาก cursor และ map ลงใน struct 
			log.Fatal(err.Error())		}
	}

	// ผลลัพธ์
	return float64(result.Total)
}


// func GetProductID(username string,c *fiber.Ctx) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// var order models.Order

// 	// Find the user document
// 	filter := bson.M{"username": username}
// 	var order models.Order

// 	err := OrderCollection.FindOne(ctx, filter).Decode(&order)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Extract productIds from user's carts
// 	// var productIds []string
// 	for _, cartEntry := range order.Carts {
// 		// productIds = append(productIds, cartEntry.ProductID)
// 	}

// 	// Print the result (you can use the productIds array as needed)
// 	// fmt.Println("Product IDs:", productIds)

// }