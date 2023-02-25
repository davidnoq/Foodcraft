package main

import (
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

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestListRecipesHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/recipes", ts.URL))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode) // check that the status code is 200
	data, _ := ioutil.ReadAll(resp.Body)

	var recipes []models.Recipe
	json.Unmarshal(data, &recipes)
	assert.Equal(t, len(recipes), 21) // check that there are 20 recipes in the database
}