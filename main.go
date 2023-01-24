package main

import (
	"context"
	"fmt"
	"log"

	"example.com/gin-api/controllers"
	"example.com/gin-api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	us          services.Userservices
	uc          controllers.UserController
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("mongodb+srv://raj:mc6x*BHVvy_LWf5@cluster0.3gbrm.mongodb.net/test")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	userc = mongoclient.Database("userdb").Collection("users")
	us = services.NewUserservice(userc, ctx)
	uc = controllers.New(us)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("v1")
	uc.RegisterUserRouter(basepath)

	log.Fatal(server.Run(":9090"))

}
