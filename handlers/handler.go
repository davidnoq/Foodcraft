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
	//userID, _ := c.Get("userID")
	// convert array of ingredients to string so that it can be in proper format for the url for api call
	ingredientsString := strings.Join(ingredients.IngredientList, "%2c")
	url := "https://spoonacular-recipe-food-nutrition-v1.p.rapidapi.com/recipes/findByIngredients?ingredients=" + ingredientsString + "&number=5&ignorePantry=true&ranking=1"
	// make api GET request to spoonacular API to search for recipes by ingredients
	// can edit url parameters to change number of recipes returned
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

	//recipe to database commented out for now

	/*newRecipe.UserID = userID.(string)

	// check if it already exists for user
	recipeInt := newRecipe.ID

	err := handler.collection.FindOne(handler.ctx, bson.M{"userId": userID, "id": recipeInt}).Decode(&newRecipe)

	if err != mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Recipe already in user's list"})
		return
	}*/

	for i := 0; i < len(recipes); i++ {
		_, err := handler.collection.InsertOne(handler.ctx, recipes[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	//print new recipe struct that contains the recipe corresponding to the input ingredients
	c.JSON(http.StatusOK, recipes)
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

func (handler *RecipesHandler) FeaturedRecipeHandler(c *gin.Context) {
	//spoonaular API call for generating random recipe
	url := "https://spoonacular-recipe-food-nutrition-v1.p.rapidapi.com/recipes/random"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", "0e2d3a4b52msh4f7ca3d8295bc0ap1374f1jsnaa5308ae1f95")
	req.Header.Add("X-RapidAPI-Host", "spoonacular-recipe-food-nutrition-v1.p.rapidapi.com")
	res, _ := http.DefaultClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}
	userID, _ := c.Get("userID")
	body, _ := ioutil.ReadAll(res.Body)
	//create a new recipe struct, then put the result of the API call into recipe
	var recipes models.FeaturedRecipe
	_ = json.Unmarshal(body, &recipes)
	//need to convert recipes.Recipes[0] into models.Recipe struct
	newRecipe := ConvertFeaturedRecipeToRecipe(recipes)
	newRecipe.UserID = userID.(string)

	recipeInt := newRecipe.ID

	//check for duplicates
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

func ConvertFeaturedRecipeToRecipe(featuredRecipe models.FeaturedRecipe) models.Recipe {
	//the spoonacular API endpoint for generating a random recipe returns a struct of a 
	// different format from the one we use, so we need to convert it
	recipe := featuredRecipe.Recipes[0]
	var newRecipe models.Recipe
	newRecipe.ID = recipe.ID
	newRecipe.Title = recipe.Title
	newRecipe.Image = recipe.Image
	newRecipe.ImageType = recipe.ImageType
	newRecipe.UsedIngredientCount = 0
	newRecipe.MissedIngredientCount = 0

	newRecipe.UsedIngredients = make([]struct {
		ID           int      `bson:"id"`
		Amount       int      `bson:"amount"`
		Unit         string   `bson:"unit"`
		UnitLong     string   `bson:"unitLong"`
		UnitShort    string   `bson:"unitShort"`
		Aisle        string   `bson:"aisle"`
		Name         string   `bson:"name"`
		Original     string   `bson:"original"`
		OriginalName string   `bson:"originalName"`
		Meta         []string `bson:"meta"`
		Image        string   `bson:"image"`
	}, len(recipe.ExtendedIngredients))

	// set the list of ingredients from featured recipe to the list of used ingredients in recipe struct
	for i, ingredient := range recipe.ExtendedIngredients {
		newRecipe.UsedIngredients[i].ID = ingredient.ID
		newRecipe.UsedIngredients[i].Amount = int(ingredient.Amount)
		newRecipe.UsedIngredients[i].Unit = ingredient.Unit
		newRecipe.UsedIngredients[i].UnitLong = ingredient.UnitLong
		newRecipe.UsedIngredients[i].UnitShort = ingredient.UnitShort
		newRecipe.UsedIngredients[i].Aisle = ingredient.Aisle
		newRecipe.UsedIngredients[i].Name = ingredient.Name
		newRecipe.UsedIngredients[i].Original = ingredient.OriginalString
		newRecipe.UsedIngredients[i].OriginalName = ingredient.Name
		newRecipe.UsedIngredients[i].Meta = ingredient.MetaInformation
		newRecipe.UsedIngredients[i].Image = ingredient.Image
	}

	newRecipe.Likes = recipe.AggregateLikes

	return newRecipe
}

func (handler *RecipesHandler) AddOneRecipeHandler(c *gin.Context) {
	recipeID := c.Param("ID")
	recipeInt, err := strconv.Atoi(recipeID)

	userID, _ := c.MustGet("userID").(string)

	// check if recipe and user combo exists
	var result models.Recipe
	err = handler.collection.FindOne(handler.ctx, bson.M{"userId": userID, "id": recipeInt}).Decode(&result)
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Recipe already exists for user"})
		return

	}

	err = handler.collection.FindOne(handler.ctx, bson.M{"id": recipeInt}).Decode(&result)

	var newRecipe models.Recipe
	newRecipe = result

	newRecipe.UserID = userID

	_, err = handler.collection.InsertOne(handler.ctx, newRecipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// check if recipe & user combo  now exists
	err = handler.collection.FindOne(handler.ctx, bson.M{"userId": userID, "id": recipeInt}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Recipe " + recipeID + " is not in user's collection"})
			return
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recipe " + recipeID + " added for user"})
}
