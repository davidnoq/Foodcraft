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

func NewRecipesHandler(ctx context.Context, collection *mongo.Collection, userCollection *mongo.Collection) *RecipesHandler {
	return &RecipesHandler{
		collection: collection,
		//userCollection: userCollection,
		ctx: ctx,
	}
}

// can access all variables of struct since it has RecipesHandler type
func (handler *RecipesHandler) ListRecipesHandler(c *gin.Context) {
	cur, err := handler.collection.Find(handler.ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(handler.ctx)
	var recipes []models.Recipe

	// decode one at a time into recipe struct then append to list of recipes
	for cur.Next(handler.ctx) {
		var recipe models.Recipe
		cur.Decode(&recipe)

		recipes = append(recipes, recipe)
	}

	c.JSON(http.StatusOK, recipes)
}

// won't need after 2 handlers after are done
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
	var recipes []models.Recipe
	_ = json.Unmarshal(body, &recipes)
	newRecipe := recipes[0]

	_, err := handler.collection.InsertOne(handler.ctx, newRecipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//print new recipe struct that contains the recipe corresponding to the input ingredients
	c.JSON(http.StatusOK, newRecipe)
}

// add a recipe to a user's list of recipes - not finished
func (handler *RecipesHandler) NewUserRecipeHandler(c *gin.Context) {

	cur, err := handler.collection.Find(handler.ctx, bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(handler.ctx)
	var newRecipe models.Recipe
	// take in desired ID from user and store in variable
	var ID int
	if err := c.BindJSON(&ID); err != nil {
		return
	}

	// decode one at a time into recipe struct and if found append to list of recipes
	for cur.Next(handler.ctx) {
		var recipe models.Recipe
		cur.Decode(&recipe)

		if recipe.ID == ID {
			newRecipe = recipe
			break
		}

	}
	//print new recipe struct that contains the recipe corresponding to the input ID
	c.JSON(http.StatusOK, newRecipe)

	// add to user's struct
	/*claims := &Claims{}
	username := claims.Username

	idk, err := handler.userCollection.UpdateOne({username: username}, {$addToSet: {Recipes: newRecipe}})*/

}

// insert all recipes from api to recipes collection - only use once
func (handler *RecipesHandler) AllRecipesHandler(c *gin.Context) {

}
