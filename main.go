package main

import (
	"context"
	"log"

	handlers "foodcraft/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var authHandler *handlers.AuthHandler
var recipesHandler *handlers.RecipesHandler

func init() {
	ctx := context.Background()

	// apply URI so as to not have to pass in commandline
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://fullstack:fullstack@foodcraft.p5l8kww.mongodb.net/?retryWrites=true&w=majority"))

	// check for successful connection
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	// global variable to access endpoint handlers
	collection := client.Database("foodcraft").Collection("recipes")
	recipesHandler = handlers.NewRecipesHandler(ctx, collection)

	collectionUsers := client.Database("foodcraft").Collection("users")
	authHandler = handlers.NewAuthHandler(ctx, collectionUsers)
}
func SetupServer() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.POST("/signin", authHandler.SignInHandler)
	router.POST("/refresh", authHandler.RefreshHandler)
	router.POST("/signup", authHandler.SignUpHandler)

	authorized := router.Group("/")
	authorized.Use(authHandler.AuthMiddleware())
	{
		authorized.POST("/recipes", recipesHandler.NewRecipeHandler)
	}

	return router
}
func main() {
	SetupServer().Run()
}
