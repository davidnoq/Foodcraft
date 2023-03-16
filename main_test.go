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

// not working - 500 code, index out of range?
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
	recipe := m{
		"UsedIngredients": a{
			m{"name": "milk"},
		},
	}

	raw, _ := json.Marshal(recipe)

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

func TestSignInHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "admin",
		Password: "password",
	}

	raw, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("%s/api/signin", ts.URL), "application/json", bytes.NewBuffer(raw))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode) // check that the status code is 200

	// difficult to check token/expiration time since it changes each time
}
