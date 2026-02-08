package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRegistration struct {
	Name     string `bson:"name" json:"name" binding:"required"`
	Email    string `bson:"email" json:"email" binding:"required,email"`
	Password string `bson:"password" json:"password" binding:"required,min=6"`
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Avatar   string             `bson:"avatar" json:"avatar"`
	Password string             `bson:"password" json:"-"`
}



type Budget struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
	Category   string             `bson:"category" json:"category"`
	Amount     float64            `bson:"amount" json:"amount"`
	StartDate  primitive.DateTime `bson:"start_date" json:"start_date"`
	EndDate    primitive.DateTime `bson:"end_date" json:"end_date"`
}