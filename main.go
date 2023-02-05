package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Ingredients struct {
	IngredientList []string `json:"ingredientlist"`
}
type Recipe []struct {
	ID                    int    `json:"id"`
	Title                 string `json:"title"`
	Image                 string `json:"image"`
	ImageType             string `json:"imageType"`
	UsedIngredientCount   int    `json:"usedIngredientCount"`
	MissedIngredientCount int    `json:"missedIngredientCount"`
	MissedIngredients     []struct {
		ID           int      `json:"id"`
		Amount       float64  `json:"amount"`
		Unit         string   `json:"unit"`
		UnitLong     string   `json:"unitLong"`
		UnitShort    string   `json:"unitShort"`
		Aisle        string   `json:"aisle"`
		Name         string   `json:"name"`
		Original     string   `json:"original"`
		OriginalName string   `json:"originalName"`
		Meta         []string `json:"meta"`
		Image        string   `json:"image"`
	} `json:"missedIngredients"`
	UsedIngredients []struct {
		ID           int      `json:"id"`
		Amount       int      `json:"amount"`
		Unit         string   `json:"unit"`
		UnitLong     string   `json:"unitLong"`
		UnitShort    string   `json:"unitShort"`
		Aisle        string   `json:"aisle"`
		Name         string   `json:"name"`
		Original     string   `json:"original"`
		OriginalName string   `json:"originalName"`
		Meta         []string `json:"meta"`
		Image        string   `json:"image"`
	} `json:"usedIngredients"`
	UnusedIngredients []interface{} `json:"unusedIngredients"`
	Likes             int           `json:"likes"`
}

var recipes []Recipe
var ctx context.Context
var err error
var client *mongo.Client
var collection *mongo.Collection

func init() {
	recipes = make([]Recipe, 0)
	ctx = context.Background()
	client, err = mongo.Connect(ctx,
		options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(),
		readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
}
func NewRecipeHandler(c *gin.Context) {
	// take in desired ingredients from user and store in variable
	var ingredients Ingredients
	if err := c.BindJSON(&ingredients); err != nil {
		return
	}
	// convert array of ingredients to string so that it can be in proper format for the url for api call
	ingredientsString := strings.Join(ingredients.IngredientList, "%2c")
	url := "https://spoonacular-recipe-food-nutrition-v1.p.rapidapi.com/recipes/findByIngredients?ingredients=" + ingredientsString + "&number=1&ignorePantry=true&ranking=1"
	// make api GET request to spoonacular API to search for recipes by ingredients
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", "0e2d3a4b52msh4f7ca3d8295bc0ap1374f1jsnaa5308ae1f95")
	req.Header.Add("X-RapidAPI-Host", "spoonacular-recipe-food-nutrition-v1.p.rapidapi.com")
	res, _ := http.DefaultClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}
	//body stores the result of the API call
	body, _ := ioutil.ReadAll(res.Body)
	//create a new recipe struct, then put the result of the API call into recipe
	var recipe Recipe
	_ = json.Unmarshal(body, &recipe)
	_, err = collection.InsertOne(ctx, recipe)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Error while inserting a new recipe"})
		return
	}
	recipes = append(recipes, recipe)
	//print new recipe struct that contains the recipe corresponding to the input ingredients
	c.JSON(http.StatusOK, recipe)
}

func ListRecipesHandler(c *gin.Context) {
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(ctx)

	recipes := make([]Recipe, 0)
	for cur.Next(ctx) {
		var recipe Recipe
		cur.Decode(&recipe)
		recipes = append(recipes, recipe)
	}

	c.JSON(http.StatusOK, recipes)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.Run()
}
