package controllers

import (
	"context"
	"server/server/internal/database"
	"server/server/internal/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


func transactionCollection() *mongo.Collection {
	return database.DB.Collection("transactions")
}

func getTransactionsbyUserID(c *gin.Context)  {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	userIDParam := c.Param("user_id")
	cursor, err := transactionCollection().Find(ctx, bson.M{"user_id": userIDParam})
	if err != nil {
		c.JSON(500, gin.H{"message": "error getting transactions"})
		return 
	}
	var transactions []models.Transactions
	if err := cursor.All(ctx, &transactions); err != nil {
		c.JSON(500, gin.H{"message": "error getting transactions"})
		return
	}
	c.JSON(200, gin.H{"data": transactions})
}

func CreateTransaction(c *gin.Context) {
	ctx,cancel:=context.WithTimeout(c.Request.Context(),10*time.Second)
	defer cancel()
	var payload models.TransactionInput
	if err:=c.BindJSON(&payload);err!=nil{
		c.JSON(403,gin.H{
			"message":"Invalid Request body",
			"error":err.Error(),	
		})
		return
	}
	_,err:=transactionCollection().InsertOne(ctx,payload)
	if err!=nil{
		c.JSON(500,gin.H{"message":"error creating transaction"})
		return
	}
	c.JSON(201,gin.H{"message":"transaction created successfully"})
}