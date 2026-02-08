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

func userCollection() *mongo.Collection {
	return database.DB.Collection("users")
}
// c *gin.context itu mirip req dan res di expressjs
func GetUsers(c *gin.Context) {
	// harus di cance biar memori ga bocor
	// 2 var ini menyimpan nilai context dan cancel function
	// merupkan multiple return assignment
	ctx, cancel:= context.WithTimeout(c.Request.Context(),10*time.Second)
	defer cancel() 
	// cursir itu mirip async await di
	cursor,err:=userCollection().Find(ctx,bson.M{})
	if err!=nil{
		c.JSON(500,gin.H{"message":"error get users"})
		return
	}
		defer cursor.Close(ctx)
	var users []models.User
if err := cursor.All(ctx,&users); err!=nil{
		c.JSON(500,gin.H{"message":"error get users"})
		return
	}
	c.JSON(200,gin.H{"data":users})
}

func GetUserByID(c *gin.Context) {
	ctx,cancel:=context.WithTimeout(c.Request.Context(),10*time.Second)
	// cancel biar memori ga bocor
	defer cancel()
	idParam:=c.Param("id")
	var user models.User
	err:=userCollection().FindOne(ctx,bson.M{"id":idParam}).Decode(&user)
	if err!=nil{
		if err==mongo.ErrNoDocuments{
			c.JSON(404,gin.H{"message":"user not found"})
			return
		}
		c.JSON(500,gin.H{"message":"error getting user"})
		return
	}
	c.JSON(200,gin.H{"data":user})
}

func CreateUser(c *gin.Context) {
	ctx,cancel:=context.WithTimeout(c.Request.Context(),10*time.Second)
	defer cancel()
	var payload models.UserRegistration
	
	if err:=c.ShouldBindJSON(&payload); err!=nil{
		c.JSON(403,gin.H{
			"message":"Invalid Request body",
			"error":err.Error(),	
			"payload":payload,
		})
		return
	}

// check if email already exists
if	err:=userCollection().FindOne(ctx,bson.M{"email":payload.Email}); err==nil{
	c.JSON(409,gin.H{"message":"email already exists"})
	return
}
_,err:=userCollection().InsertOne(ctx,payload)
if err!=nil{
	c.JSON(500,gin.H{"message":"error creating user"})
	return
}
c.JSON(201,gin.H{"message":"user created successfully"})
}