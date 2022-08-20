package main

import (
	"context"
	"fmt"
	"log"

	"book.com/sina-apis/controllers"
	"book.com/sina-apis/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server                   *gin.Engine
	ctx                      context.Context
	mongoclint               *mongo.Client
	err                       error
)
var (
	autherService            services.AutherService
	AutherController         controllers.AutherController
	authercollection         *mongo.Collection
)
var (
	userservice              services.UserService
	UserController           controllers.UserController
	usercollection           *mongo.Collection

)
func init(){
	ctx:=context.TODO()
	mongoconn:=options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclint,err=mongo.Connect(ctx,mongoconn)
	if err !=nil {
		log.Fatal(err)
	}
	err=mongoclint.Ping(ctx,readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongo connection established")
	usercollection=mongoclint.Database("autherdb").Collection("authers")
	userservice=services.Newuserservice(usercollection,ctx)
	UserController=controllers.New1(userservice)
	/////////////////////////////////////////////////
	authercollection=mongoclint.Database("autherdb").Collection("authers")
	autherService=services.NewAutherservice(authercollection,ctx)
	AutherController=controllers.New(autherService)

	server=gin.Default()

    

}

func main(){
	defer mongoclint.Disconnect(ctx)

	AbasePath:=server.Group("/v1")
	UbasePath:=server.Group("/v1")
	AutherController.RegisterAutherRouter(AbasePath)
	UserController.RegisterUserRoutes(UbasePath)
	log.Fatal(server.Run(":8080")) 



}