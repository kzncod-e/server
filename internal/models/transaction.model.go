package models

import "go.mongodb.org/mongo-driver/bson/primitive"
type Transactions struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	Amount      float64            `bson:"amount" json:"amount"`
	Category    string           `bson:"category" json:"category"`
	Description string             `bson:"description" json:"description"`
	Date        primitive.DateTime `bson:"date" json:"date"`
	Type        string             `bson:"type" json:"type"` // "income" or "expense"
}




type TransactionInput struct {
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id" binding:"required"`
	Amount      float64            `bson:"amount" json:"amount" binding:"required"`
	Category    string             `bson:"category" json:"category" binding:"required"`
	Description string             `bson:"description" json:"description"`
	Date        string             `bson:"date" json:"date" binding:"required"` // Expecting date in string format, e.g., "2023-10-01"
	Type        string             `bson:"type" json:"type" binding:"required,oneof=income expense"` // "income" or "expense"
}


type wallets struct {
UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
Balance     float64            `bson:"balance" json:"balance"`
Date        primitive.DateTime `bson:"date" json:"date"`
}