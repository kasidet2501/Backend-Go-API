package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Login Struct
type LoginRequest struct {
	Username    string `json:"username"`
	Password 	string `json:"password"`
  }

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


type Name struct {
	Firstname 	string 		`json:"firstname,omitempty" validate:"required"`
	Lastname  	string 		`json:"lastname,omitempty"  validate:"required"`
}

type User struct {
	Id       	primitive.ObjectID     	`json:"id,omitempty"`
	Role		string		`json:"role,omitempty" validate:"required"`
	Email    	string  	`json:"email,omitempty" validate:"required"`
	Username 	string  	`json:"username,omitempty" validate:"required"`
	Password 	string  	`json:"password,omitempty" validate:"required"`
	Name     	Name    	`json:"name,omitempty" validate:"required"`
	Phone    	string  	`json:"phone,omitempty" validate:"required"`
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