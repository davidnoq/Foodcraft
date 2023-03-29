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
	Error string `json:"error"`
}

type (
	a []interface{}
	m map[string]interface{}
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestListRecipesHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/api/recipes", ts.URL))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode) // check that the status code is 200
	data, _ := ioutil.ReadAll(resp.Body)

	var recipes []models.Recipe
	json.Unmarshal(data, &recipes)
	assert.Equal(t, len(recipes), 25) // check that there are 24 recipes in the database
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
		Username: "admin10",
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

	/*post := &Post{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}

	assert.Equal(t, post.Error, "Token is not expired yet")*/
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

func TestUserSpecificRecipeList(t *testing.T){
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "unitTester1",
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

	// add 2 recipes to user
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
	
	var ingredients2 models.Ingredients
    ingredients2.IngredientList = []string{"eggs", "flour"}
	raw2, _ := json.Marshal(ingredients2)

	r2, err := http.NewRequest("POST", fmt.Sprintf("%s/api/recipes", ts.URL), bytes.NewBuffer(raw2))
    if err != nil {
        panic(err)
    }

	r2.Header.Add("Authorization", JWTtoken)
	res2, err := client.Do(r2)
	defer res2.Body.Close()
	assert.Nil(t, err)
    assert.Equal(t, http.StatusOK, res2.StatusCode) // check that the status code is 200

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
		Username: "UnitTester2",
		Password: "UnitTester2",
	}
	
	raw3, _ := json.Marshal(user2)
	resp1, _ := http.Post(fmt.Sprintf("%s/api/signin", ts.URL), "application/json", bytes.NewBuffer(raw3))
	defer resp1.Body.Close()
	data1, _ := ioutil.ReadAll(resp1.Body)

	var payload1 map[string]string
	json.Unmarshal(data1, &payload1)

	JWTtoken1 := payload1["token"]

	// add recipe to second user
	var ingredients4 models.Ingredients
	ingredients4.IngredientList = []string{"bacon"}

	raw4, _ := json.Marshal(ingredients4)

	r4, err := http.NewRequest("POST", fmt.Sprintf("%s/api/recipes", ts.URL), bytes.NewBuffer(raw4))
	if err != nil {
		panic(err)
	}

	r4.Header.Add("Authorization", JWTtoken1)

	client = &http.Client{}
	res4, err := client.Do(r4)

	defer res.Body.Close()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res4.StatusCode) // check that the status code is 200
	
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

