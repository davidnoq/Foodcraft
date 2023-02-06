package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"foodcraft/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type RecipesHandler struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewRecipesHandler(ctx context.Context, collection *mongo.Collection) *RecipesHandler {
	return &RecipesHandler{
		collection: collection,
		ctx:        ctx,
	}
}

func (handler *RecipesHandler) ListRecipesHandler(c *gin.Context) {
	cur, err := handler.collection.Find(handler.ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(handler.ctx)
	//var recipes []models.Recipe
	var recipes []interface{}
	for cur.Next(handler.ctx) {
		/*
		   var recipe models.Recipe
		   err := cur.Decode(&recipe)
		   if err != nil {
		       c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		   }
		   recipes = append(recipes, recipe)
		*/
		var recipe interface{}
		cur.Decode(&recipe)

		recipes = append(recipes, recipe)
	}

	c.JSON(http.StatusOK, recipes)
}

func (handler *RecipesHandler) NewRecipeHandler(c *gin.Context) {
	// take in desired ingredients from user and store in variable
	var ingredients models.Ingredients
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
	var recipe models.Recipe
	_ = json.Unmarshal(body, &recipe)
	/*
		_ = json.Unmarshal(body, &recipe)
		var datas []interface{}
		for _, r := range recipe {
			datas = append(datas, r)
		}
	*/
	_, err := handler.collection.InsertOne(handler.ctx, recipe[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		//c.JSON(http.StatusInternalServerError, datas[0])
		return
	}
	//print new recipe struct that contains the recipe corresponding to the input ingredients
	c.JSON(http.StatusOK, recipe)
}
