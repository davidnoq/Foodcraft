
# Detail work you've completed in Sprint 2
**Frontend**
- 

**Backend**
- 

# List unit tests and Cypress test for frontend
- 

# List unit tests for backend
- GET for retrieving recipes in database
- POST for adding a recipe to the database
- POST for signing into an existing and nonexistant user
- POST for signing up with an existing and nonexistant user
- POST for refreshing JWT token after signin

# Foodcraft Backend Documentation
## Overview
The Foodcraft backend is a RESTful API server built with Go and Gin that provides endpoints for managing recipes, getting recipe recommendations and user authentication. It is designed to connect to a MongoDB database to store and retrieve data.
## Architecture
The backend is built using the Gin web framework and connects to a MongoDB database using the official Go driver. It consists of two main components:
Recipe API handlers: These are responsible for implementing functionality of our recipe application 
Authentication middleware: This middleware enables user authentication and provides a layer of user security
## APIs
`GET /api/recipes`: Returns a list of recipes.

`POST /api/signin`: Authenticates a user and generates a JWT access token.

`POST /api/refresh`: Refreshes a JWT access token.
`POST /api/signup`: Registers a new user.
`POST /api/recipes`: Creates a new recipe. (requires authentication)
## Authentication and Authorization
Authentication is performed using JWT access tokens. When a user successfully authenticates using their credentials, a JWT access token is generated and returned in the response. This access token can be used to authenticate future requests by including it in the Authorization header of the request.

Authorization is enforced using middleware that checks the validity of the access token included in the Authorization header of the request. 
Passwords are hashed and salted using the HS256 hashing algorithm before being stored in the MONGODB database to protect personal data.
## Data Models
The backend stores data in two collections:
recipes: Stores recipe data, including the recipe name, ingredients, and instructions.
users: Stores user data, including the user email and password hash.
## Testing
Automated tests can be run using the `go test` command. Tests are located in the handlers/handlers_test.go file and cover both REST API handlers and authentication middleware.
## Bugs
Currently were are having trouble with resolving CORS issues within our backend.
