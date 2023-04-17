package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"foodcraft/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	//"github.com/dgrijalva/jwt-go"
)

type Post struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type (
	a []interface{}
	m map[string]interface{}
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestSignInHandlerFail(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "abcd",
		Password: "password",
	}

	raw, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("%s/api/signin", ts.URL), "application/json", bytes.NewBuffer(raw))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode) // check that the status code is 401 (invalid user/password)
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	assert.Equal(t, payload["error"], "Invalid username or password")
}

func TestSignUpHandlerFail(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "admin",
		Password: "password",
	}

	raw, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("%s/api/signup", ts.URL), "application/json", bytes.NewBuffer(raw))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode) // check that the status code is 500 (user exists already)
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	assert.Equal(t, payload["error"], "Username already taken")
}

func TestSignUpHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "testSignUp",
		Password: "password",
	}

	raw, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("%s/api/signup", ts.URL), "application/json", bytes.NewBuffer(raw))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusAccepted, resp.StatusCode) // check that the status code is 202
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	assert.Equal(t, payload["message"], "Account has been created")

	// delete user so test works next time
	raw, _ = json.Marshal(user)
	res, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/users", ts.URL), bytes.NewBuffer(raw))

	client := &http.Client{}
	resp, err = client.Do(res)

	defer res.Body.Close()
	defer resp.Body.Close()
}

func TestRefreshHandlerUnauthorized(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "admin",
		Password: "password",
	}

	raw, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("%s/api/refresh", ts.URL), "application/json", bytes.NewBuffer(raw))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode) // check that the status code is 401 (no sign in yet)
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	assert.Equal(t, payload["error"], "token contains an invalid number of segments")
}

func TestNewRecipeHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "newRecipe",
		Password: "password",
	}
	// sign in
	raw1, _ := json.Marshal(user)
	resp, _ := http.Post(fmt.Sprintf("%s/api/signin", ts.URL), "application/json", bytes.NewBuffer(raw1))
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	JWTtoken := payload["token"]

	// add recipe
	var ingredients models.Ingredients
	ingredients.IngredientList = []string{"milk"}

	raw, _ := json.Marshal(ingredients)

	r, err := http.NewRequest("POST", fmt.Sprintf("%s/api/recipes", ts.URL), bytes.NewBuffer(raw))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Authorization", JWTtoken)

	client := &http.Client{}
	res, err := client.Do(r)

	defer res.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode) // check that the status code is 200
}

func TestRefreshHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "admin",
		Password: "password",
	}
	// sign in
	raw1, _ := json.Marshal(user)
	resp, _ := http.Post(fmt.Sprintf("%s/api/signin", ts.URL), "application/json", bytes.NewBuffer(raw1))
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	JWTtoken := payload["token"]

	// refresh

	raw, _ := json.Marshal(user)

	r, err := http.NewRequest("POST", fmt.Sprintf("%s/api/refresh", ts.URL), bytes.NewBuffer(raw))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Authorization", JWTtoken)

	client := &http.Client{}
	res, err := client.Do(r)

	defer res.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode) // check that the status code is 400

	post := &Post{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}

	assert.Equal(t, post.Error, "Token is not expired yet")
}

func TestUserSpecificRecipeList(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "recipeList1",
		Password: "unitTester1",
	}
	// sign in
	raw1, _ := json.Marshal(user)
	resp, _ := http.Post(fmt.Sprintf("%s/api/signin", ts.URL), "application/json", bytes.NewBuffer(raw1))
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	JWTtoken := payload["token"]

	//check that there are 2 recipes associated to user1
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/recipes", ts.URL), nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", JWTtoken)
	clientGet := &http.Client{}
	resGet, err := clientGet.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resGet.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resGet.StatusCode) // check that the status code is 200

	dataGet, _ := ioutil.ReadAll(resGet.Body)

	var recipes1 []models.Recipe
	json.Unmarshal(dataGet, &recipes1)
	assert.Equal(t, len(recipes1), 2) // check that there are 2 recipes associated to user1

	// sign in second user
	user2 := models.User{
		Username: "recipeList2",
		Password: "unitTester2",
	}

	raw3, _ := json.Marshal(user2)
	resp1, _ := http.Post(fmt.Sprintf("%s/api/signin", ts.URL), "application/json", bytes.NewBuffer(raw3))
	defer resp1.Body.Close()
	data1, _ := ioutil.ReadAll(resp1.Body)

	var payload1 map[string]string
	json.Unmarshal(data1, &payload1)

	JWTtoken1 := payload1["token"]

	//check that there is only 1 recipe associated to user2
	req2, err := http.NewRequest("GET", fmt.Sprintf("%s/api/recipes", ts.URL), nil)
	if err != nil {
		t.Fatal(err)
	}

	req2.Header.Set("Authorization", JWTtoken1)
	clientGet2 := &http.Client{}
	resGet2, err := clientGet2.Do(req2)
	if err != nil {
		t.Fatal(err)
	}
	defer resGet2.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resGet2.StatusCode) // check that the status code is 200

	dataGet2, _ := ioutil.ReadAll(resGet2.Body)

	var recipes2 []models.Recipe
	json.Unmarshal(dataGet2, &recipes2)
	assert.Equal(t, len(recipes2), 1) // check that there is only 1 recipe associated to user2
}

func TestDeleteRecipes(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "deleteRecipes",
		Password: "password",
	}
	// sign in
	raw1, _ := json.Marshal(user)
	resp, _ := http.Post(fmt.Sprintf("%s/api/signin", ts.URL), "application/json", bytes.NewBuffer(raw1))
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	JWTtoken := payload["token"]

	raw, _ := json.Marshal(user)

	// add recipe
	var ingredients models.Ingredients
	ingredients.IngredientList = []string{"milk"}

	raw, _ = json.Marshal(ingredients)

	r, err := http.NewRequest("POST", fmt.Sprintf("%s/api/recipes/666439", ts.URL), bytes.NewBuffer(raw))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Authorization", JWTtoken)

	client := &http.Client{}
	res, err := client.Do(r)

	defer res.Body.Close()

	r, err = http.NewRequest("DELETE", fmt.Sprintf("%s/api/recipes/", ts.URL), bytes.NewBuffer(raw))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Authorization", JWTtoken)

	client = &http.Client{}
	res, err = client.Do(r)

	defer res.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode) // check that the status code is 200

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/recipes", ts.URL), nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", JWTtoken)
	clientGet := &http.Client{}
	resGet, err := clientGet.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resGet.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resGet.StatusCode) // check that the status code is 200

	dataGet, _ := ioutil.ReadAll(resGet.Body)

	var recipes1 []models.Recipe
	json.Unmarshal(dataGet, &recipes1)
	assert.Equal(t, len(recipes1), 0) // check that there are 0 recipes associated to user1

	post := &Post{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}

	assert.Equal(t, post.Message, "All recipes deleted for user")
}

func TestDeleteUser(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "testDeleteUser",
		Password: "password",
	}

	raw, _ := json.Marshal(user)
	resp, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/users", ts.URL), bytes.NewBuffer(raw))

	client := &http.Client{}
	res, err := client.Do(resp)

	defer res.Body.Close()
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode) // check that the status code is 202
	data, _ := ioutil.ReadAll(res.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	assert.Equal(t, payload["message"], "User has been deleted")

	// remake user so test works next time
	raw, _ = json.Marshal(user)
	res, err = http.Post(fmt.Sprintf("%s/api/signup", ts.URL), "application/json", bytes.NewBuffer(raw))
	defer res.Body.Close()
}

func TestGetUsername(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "admin",
		Password: "password",
	}
	// sign in
	raw1, _ := json.Marshal(user)
	resp, _ := http.Post(fmt.Sprintf("%s/api/signin", ts.URL), "application/json", bytes.NewBuffer(raw1))
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	JWTtoken := payload["token"]

	// get username

	raw, _ := json.Marshal(user)

	r, err := http.NewRequest("GET", fmt.Sprintf("%s/api/user", ts.URL), bytes.NewBuffer(raw))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Authorization", JWTtoken)

	client := &http.Client{}
	res, err := client.Do(r)

	defer res.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode) // check that the status code is 200

	dataGet, _ := ioutil.ReadAll(res.Body)

	var pay map[string]string
	json.Unmarshal(dataGet, &pay)

	assert.Equal(t, pay["username"], "admin")
}

func TestAddAndDeleteOneRecipe(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "deleteOne",
		Password: "password",
	}
	// sign in
	raw1, _ := json.Marshal(user)
	resp, _ := http.Post(fmt.Sprintf("%s/api/signin", ts.URL), "application/json", bytes.NewBuffer(raw1))
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	JWTtoken := payload["token"]

	raw, _ := json.Marshal(user)

	// add recipe to user
	/*var ingredients models.Ingredients
	ingredients.IngredientList = []string{"milk"}

	raw, _ = json.Marshal(ingredients)*/

	r, err := http.NewRequest("POST", fmt.Sprintf("%s/api/recipes/1090966", ts.URL), bytes.NewBuffer(raw))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Authorization", JWTtoken)

	client := &http.Client{}
	res, err := client.Do(r)

	defer res.Body.Close()

	r, err = http.NewRequest("DELETE", fmt.Sprintf("%s/api/recipes/1090966", ts.URL), bytes.NewBuffer(raw))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Authorization", JWTtoken)

	client = &http.Client{}
	res, err = client.Do(r)

	defer res.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode) // check that the status code is 200

	dataGet, _ := ioutil.ReadAll(res.Body)

	// check if recipe still exists for user
	r, err = http.NewRequest("GET", fmt.Sprintf("%s/api/userRecipe/666439", ts.URL), bytes.NewBuffer(raw))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Authorization", JWTtoken)

	client = &http.Client{}
	res, err = client.Do(r)

	defer res.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode) // check that the status code is 500

	var pay map[string]string
	json.Unmarshal(dataGet, &pay)

	assert.Equal(t, pay["message"], "Recipe 1090966 deleted for user")

}

func TestFindOneRecipe(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "findOne",
		Password: "password",
	}
	// sign in
	raw1, _ := json.Marshal(user)
	resp, _ := http.Post(fmt.Sprintf("%s/api/signin", ts.URL), "application/json", bytes.NewBuffer(raw1))
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	JWTtoken := payload["token"]

	raw, _ := json.Marshal(user)

	// check if recipe still exists for user
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/api/userRecipe/666439", ts.URL), bytes.NewBuffer(raw))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Authorization", JWTtoken)

	client := &http.Client{}
	res, err := client.Do(r)

	defer res.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode) // check that the status code is 200

	dataGet, _ := ioutil.ReadAll(res.Body)

	var pay map[string]string
	json.Unmarshal(dataGet, &pay)
	// "id": recipeID, "userID": userID
	assert.Equal(t, pay["id"], "666439")
	assert.Equal(t, pay["userID"], "643b1df79091eb7c7e371c64")
}

func TestGetRecipeInstructions(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "admin",
		Password: "password",
	}
	// sign in
	raw1, _ := json.Marshal(user)
	resp, _ := http.Post(fmt.Sprintf("%s/api/signin", ts.URL), "application/json", bytes.NewBuffer(raw1))
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	JWTtoken := payload["token"]

	//request recipe instructions
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/api/recipes/1040243/instructions", ts.URL), bytes.NewBuffer(raw1))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Authorization", JWTtoken)

	client := &http.Client{}
	res, err := client.Do(r)

	defer res.Body.Close()

	assert.Nil(t, err)

	dataGet, _ := ioutil.ReadAll(res.Body)
	var pay map[string]string
	json.Unmarshal(dataGet, &pay)
	//verify instructions are properly retrieved
	assert.Equal(t, pay["instructions"], "Instructions\n\n\n\nPlace white candy melts in a microwave safe bowl, and melt in the microwave according to package directions. \n\n\nStir until smooth. \n\n\nDip rice krispie treats into the candy melts, to fully coat the top of the treats. (You can cover the sides if you'd like as well) Place onto a baking sheet lined with parchment paper. Repeat with remaining rice krispie treats.\n\n\n Allow to set until candy hardens, or place in the fridge for 2-3 minutes if needed. \n\n\nSpoon remaining chocolate into a small ziplock bag. \n\n\nSnip the corner of the bag off with scissors. \n\n\nSqueeze a little white chocolate onto the bottom of each of the eye balls, to place 2 onto each rice krispie treat. \n\n\nDrizzle chocolate across the top of the rice krispie treats in a random pattern to create the look of a mummy wrap. Repeat with additional rice krispie treats. \n\n\nAllow chocolate to fully set. \n\n\nServe immediately. Store in an airtight container for up to 5 or 6 days.")
}
