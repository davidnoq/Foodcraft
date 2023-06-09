
# Detail work you've completed in Sprint 4
**Frontend**
- styled frontend to replicate theme, 'Minecraft'.
- added functionality to search recipe component to output recipe after user selects ingredients
- added loading animation when the frontend is waiting for response for recipe
- allow user to add ingredients that aren't listed already
- added functionality to clicking on a recipe, it displays a dialog with more details regarding recipe.
- allow user to 'like' a recipe, which in turn adds it to the user profile
- profile displays user liked recipes
- allow user to clear all recipes in profile
- allow user to delete account
- allow user to choose from preselected recipes on the dropdown menu
- allows user to get inspiration from the featured recipe of the day and see ingredients plus instructions
- fixed about page to look more straightforward
- fixed home page and styling of the entire website including the buttons

**Backend**
- added DELETE function for user accounts
- added DELETE function for removing single recipe from user's list
- added handler for retrieving instructions of a specific recipe through a GET request to Spoonacular endpoint
- added handler for generating a random recipe
- added GET function for frontend to retrieve username
- added GET function to find a specific recipe within a user's list
- refactored new recipe handler to generate ten recipe recommendations based on user inputted ingredients instead of just one
- added handler for adding recipe to a user's list
- created unit tests for new changes

# List unit tests for frontend
- Used Cypress component tests for unit testing
- Search component isolated and tested buttons, increments, visual texts, and type function
- Login component isolated and tested filling out form for both sign up and login
- End-to-end testing by checking url extensions and moving through pages
- About component isolated to test specific aspects of the page, button
- About e2e test to determine if website if visitable.
- Login e2e testing if profile shows up when user logs in or signs up
- Login e2e testing logout function when the user is logged in
- Checking whether a profile exists and making a new profile for login component testing
- check whether contact button is visible and can be accessed
- Check when the user input ingredients, user receives a recipe
- Test clear all recipes functions for profile
- Check if post to the backend is recieved from the e2e tests on beef,chicken,pork, etc
- completed new component tests for the new additions, checking if each component contains the generated recipes

# List unit tests for backend
- GET for retrieving recipes in database associated with a user (user-specific recipe lists)
- POST for adding a recipe to the database
- POST for signing into an existing and nonexistant user
- POST for signing up with an existing and nonexistant user
- POST for refreshing JWT token after signin
- DELETE for removing all recipes from a user's list
- DELETE for removing a user from the database
- GET and DELETE for adding and removing a single recipe from a user's list
- GET for recipe instructions
- POST for Featured Recipes
- GET for retrieving username
- GET for finding a recipe within a user's list

------------

# Foodcraft Backend Documentation
## Overview
The Foodcraft backend is a RESTful API server built with Go and Gin that provides endpoints for managing recipes, getting recipe recommendations and user authentication. It is designed to connect to a MongoDB database to store and retrieve data.
## Tech stack
- Go to build the backend server
- Gin to create the REST API
- MONGODB for database
- JWT for user authentication

## Architecture
The backend is built using the Gin web framework and connects to a MongoDB database using the official Go driver. It consists of two main components:

Recipe API handlers: These are responsible for implementing functionality of our recipe application 

Authentication middleware: This middleware enables user authentication and provides a layer of user security
## APIs
`GET /api/recipes`: Returns a user's list of recipes

`POST /api/signin`: Authenticates a user and generates a JWT access token.

`POST /api/refresh`: Refreshes a JWT access token.

`POST /api/signup`: Registers a new user.

`POST /api/recipes`: Generates 10 recipe recommendations for user.

`DELETE /api/recipes`: Clears a user's recipe list.

`GET /api/recipes/:ID/instructions`: returns instructions for the specificed recipe.

`GET /api/user`: returns username of currently logged-in user.

`DELETE /api/recipes/:ID`: deletes the specified recipe from the user's recipe list.

`GET /api/userRecipe/:ID`: returns a specific recipe within a user's list.

`POST /api/recipes/featured`: generates a random recipe using Spoonacular's custom endpoint and adds it to the user's list.

`GET /api/recipes/:ID`: adds the specified recipe to the database with the corresponding user ID.

## Authentication and Authorization
Authentication is performed using JWT access tokens. When a user successfully authenticates using their credentials, a JWT access token is generated and returned in the response. This access token can be used to authenticate future requests by including it in the Authorization header of the request.

Authorization is enforced using middleware that checks the validity of the access token included in the Authorization header of the request. 
Passwords are hashed and salted using the HS256 hashing algorithm before being stored in the MONGODB database to protect personal data.
## Data Models
The backend stores data in two main collections:

`recipes`: Stores recipe data, including the recipe name, ingredients, etc.

`users`: Stores user data as the username and password hash.

## Testing
Automated tests can be run using the `go test` command. Tests are located in the main_test.go file and cover both REST API handlers and authentication middleware.
## Bugs
