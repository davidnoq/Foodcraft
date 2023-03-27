package main

import (
	"context"
	"log"

	handlers "foodcraft/handlers"

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
	collectionUsers := client.Database("foodcraft").Collection("users")

	recipesHandler = handlers.NewRecipesHandler(ctx, collection, collectionUsers)
	authHandler = handlers.NewAuthHandler(ctx, collectionUsers)
}

func SetupServer() *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	api := router.Group("/api")
	//api.GET("/recipes", recipesHandler.ListRecipesHandler)
	api.POST("/signin", authHandler.SignInHandler)
	api.POST("/refresh", authHandler.RefreshHandler)
	api.POST("/signup", authHandler.SignUpHandler)

	authorized := api.Group("/")
	authorized.Use(authHandler.AuthMiddleware())
	{
		authorized.POST("/recipes", recipesHandler.NewRecipeHandler)
		authorized.GET("/recipes", recipesHandler.ListRecipesHandler)
	}

	return router
}
func main() {
	SetupServer().Run()
}
