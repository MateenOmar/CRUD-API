package main

import (
	"context"
	"example/CRUD-APIs/controllers"
	"example/CRUD-APIs/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	itemService    services.ItemService
	ItemController controllers.ItemController
	c              context.Context
	itemCollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	c = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(c, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(c, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}
	fmt.Println("mongo connection established")

	itemCollection = mongoclient.Database("itemdb").Collection("items")
	itemService = services.NewItemService(itemCollection, c)
	ItemController = controllers.New(itemService)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(c)

	basepath := server.Group("/")
	ItemController.RegisterItemRoutes(basepath)

	port := os.Getenv("PORT")

	log.Fatal(server.Run(port))
}
