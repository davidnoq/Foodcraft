package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
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
	userID, _ := c.MustGet("userID").(string)
	cur, err := handler.collection.Find(handler.ctx, bson.M{"userId": userID})
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

func (handler *RecipesHandler) NewRecipeHandler(c *gin.Context) {
	// take in desired ingredients from user and store in variable
	var ingredients models.Ingredients
	if err := c.BindJSON(&ingredients); err != nil {
		return
	}
	userID, _ := c.Get("userID")
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
	newRecipe.UserID = userID.(string)

	// check if it already exists for user
	recipeInt := newRecipe.ID

	err := handler.collection.FindOne(handler.ctx, bson.M{"userId": userID, "id": recipeInt}).Decode(&newRecipe)

	if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Recipe already in user's list"})
		return
	}

	_, err = handler.collection.InsertOne(handler.ctx, newRecipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//print new recipe struct that contains the recipe corresponding to the input ingredients
	c.JSON(http.StatusOK, newRecipe)
}

func (handler *RecipesHandler) DeleteAllRecipesHandler(c *gin.Context) {
	userID, _ := c.MustGet("userID").(string)
	_, err := handler.collection.DeleteMany(handler.ctx, bson.M{"userId": userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All recipes deleted for user"})
}

func (handler *RecipesHandler) InstructionsForRecipeHandler(c *gin.Context) {
	recipeID := c.Param("ID")

	// Set up the Spoonacular API URL
	url := "https://spoonacular-recipe-food-nutrition-v1.p.rapidapi.com/recipes/" + recipeID + "/information"

	// Make API GET request to Spoonacular API to get information for recipe
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", "0e2d3a4b52msh4f7ca3d8295bc0ap1374f1jsnaa5308ae1f95")
	req.Header.Add("X-RapidAPI-Host", "spoonacular-recipe-food-nutrition-v1.p.rapidapi.com")
	res, _ := http.DefaultClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}

	// Read the response body
	body, _ := ioutil.ReadAll(res.Body)

	var information models.Information
	_ = json.Unmarshal(body, &information)
	//return the instructions for the specified recipe
	c.JSON(http.StatusOK, gin.H{"instructions": information.Instructions})
}

func (handler *RecipesHandler) DeleteOneRecipesHandler(c *gin.Context) {
	recipeID := c.Param("ID")
	recipeInt, err := strconv.Atoi(recipeID)

	userID, _ := c.MustGet("userID").(string)

	// check if recipe & user combo exists
	var result models.Recipe
	err = handler.collection.FindOne(handler.ctx, bson.M{"userId": userID, "id": recipeInt}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Recipe " + recipeID + " is not in user's collection"})
			return
		}
		return
	}

	_, err = handler.collection.DeleteOne(handler.ctx, bson.M{"userId": userID, "id": recipeInt})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// check if it still exists
	err = handler.collection.FindOne(handler.ctx, bson.M{"userId": userID, "id": recipeInt}).Decode(&result)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe " + recipeID + " deleted for user"})
}

func (handler *RecipesHandler) FindRecipeHandler(c *gin.Context) {
	recipeID := c.Param("ID")
	recipeInt, err := strconv.Atoi(recipeID)

	userID, _ := c.MustGet("userID").(string)

	// check if recipe & user combo exists
	var result models.Recipe
	err = handler.collection.FindOne(handler.ctx, bson.M{"userId": userID, "id": recipeInt}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Recipe is not in user's collection"})
			return
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": recipeID, "userID": userID})
}

func (handler *RecipesHandler) FeaturedRecipeHandler(c *gin.Context){
	url := "https://spoonacular-recipe-food-nutrition-v1.p.rapidapi.com/recipes/random"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", "0e2d3a4b52msh4f7ca3d8295bc0ap1374f1jsnaa5308ae1f95")
	req.Header.Add("X-RapidAPI-Host", "spoonacular-recipe-food-nutrition-v1.p.rapidapi.com")
	res, _ := http.DefaultClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}
	
	body, _ := ioutil.ReadAll(res.Body)
	//create a new recipe struct, then put the result of the API call into recipe
	var recipes models.FeaturedRecipe
	_ = json.Unmarshal(body, &recipes)
	//print new recipe struct that contains the recipe corresponding to the input ingredients
	c.JSON(http.StatusOK, recipes)
}

