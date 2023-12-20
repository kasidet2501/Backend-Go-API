package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product Struct
type Product struct{
	Id 			primitive.ObjectID `json:"id,omitempty"`
	Title 		string             `json:"title,omitempty" validate:"required"`
	Price		float64			   `json:"price,omitempty" validate:"required"`
	Description	string             `json:"description,omitempty" validate:"required"`
	Category	string             `json:"category,omitempty" validate:"required"`
	Image		string             `json:"image,omitempty" validate:"required"`
}


// User Struct
type Address struct {
	City        string      `json:"city"`
	Street      string      `json:"street"`
	Number      int         `json:"number"`
	Zipcode     string      `json:"zipcode"`
}

type Name struct {
	Firstname 	string 		`json:"firstname" validate:"required"`
	Lastname  	string 		`json:"lastname"  validate:"required"`
}

type User struct {
	Address  	Address 	`json:"address"`
	Id       	primitive.ObjectID     	`json:"id"`
	Email    	string  	`json:"email" validate:"required"`
	Username 	string  	`json:"username" validate:"required"`
	Password 	string  	`json:"password" validate:"required"`
	Name     	Name    	`json:"name" validate:"required"`
	Phone    	string  	`json:"phone" validate:"required"`
}


//Cart
type Cart struct {
	ProductID int 		`json:"productId"`
	Quantity  int 		`json:"quantity"`
}

type Order struct {
	ID       int       `json:"id"`
	UserID   int       `json:"userId"`
	Date     time.Time `json:"date"`
	Carts 	 []Cart    `json:"carts"`
}