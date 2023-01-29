package main

import (
	"context"
	"fmt"
	"log"

	"github.com/felipdc/goilerpi/controllers"
	"github.com/felipdc/goilerpi/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	us          services.UserService
	uc          controllers.UserController
	ctx         context.Context
	userc       *mongo.Collection
	mongoclient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI("mongodb://host.docker.internal:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("Error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error while trying to ping mongo", err)
	}

	fmt.Println("Mongo connection established")

	userc = mongoclient.Database("userdb").Collection("users")
	us = services.NewUserService(userc, ctx)
	uc = controllers.New(us)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	uc.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":4000"))

}
