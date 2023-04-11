
# Detail work you've completed in Sprint 4
**Frontend**
- 

**Backend**
- added DELETE for user accounts

# List unit tests for frontend
- Used Cypress component tests for unit testing
- Search component isolated and tested buttons, increments, visual texts, and type function
- Login component isolated and tested filling out form
- End-to-end testing by checking url extensions and moving through pages
- About component isolated to test specific aspects of the page, button
- About e2e test to determine if webiste if visitable.

# List unit tests for backend
- GET for user-specific recipe lists
- POST for adding a recipe to the database
- POST for signing into an existing and nonexistant user
- POST for signing up with an existing and nonexistant user
- DELETE 

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

`POST /api/recipes`: Generates a new recipe for user.

`DELETE /api/recipes`: Clears a user's recipe list.


## Authentication and Authorization
Authentication is performed using JWT access tokens. When a user successfully authenticates using their credentials, a JWT access token is generated and returned in the response. This access token can be used to authenticate future requests by including it in the Authorization header of the request.

Authorization is enforced using middleware that checks the validity of the access token included in the Authorization header of the request. 
Passwords are hashed and salted using the HS256 hashing algorithm before being stored in the MONGODB database to protect personal data.
## Data Models
The backend stores data in two collections:

`recipes`: Stores recipe data, including the recipe name, ingredients,etc.

`users`: Stores user data as the username and password hash.
## Testing
Automated tests can be run using the `go test` command. Tests are located in the main_test.go file and cover both REST API handlers and authentication middleware.
## Bugs

