package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() {
	// กำหนด URL ของ MongoDB
	mongoURL := "mongodb://localhost:27017"

	// กำหนดตัวเลือกการเชื่อมต่อ
	clientOptions := options.Client().ApplyURI(mongoURL)

	// เชื่อมต่อกับ MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// ตรวจสอบว่าเชื่อมต่อได้
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// ปิดการเชื่อมต่อเมื่อเสร็จสิ้น
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
}

