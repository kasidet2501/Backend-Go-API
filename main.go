package main

import (
	"log"

	"github.com/backend-go-ecom/repository/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// ... (โค้ดเชื่อมต่อ MongoDB)

	db := db.New()

	app := fiber.New()

	app.Get("/users", func(c *fiber.Ctx) error {
		// ดึงข้อมูลจาก MongoDB
		collection,err := db.client.Database("go-ecom").Collection("products")
		if err != nil{
			log.Fatal(err)
		}
		// ทำสิ่งที่คุณต้องการกับ collection

		return c.SendString("List of users")
	})

	// สร้างเซิร์ฟเวอร์
	db.err = app.Listen(":3000")
	if db.err != nil {
		log.Fatal(db.err)
	}
}