﻿# User Stories
**Frontend**
- As a parent, I want to specify how large the meals I want to make, in order to have enough made for the family.
- As a user, I want to choose what type of meal I would like to find (breakfast, lunch, dinner), so that I can narrow down my searches. 
- As a chef, I want to utilize all extra ingredients in my storage so that nothing goes to waste.
- As a student, I want to find diverse meals, so that I can widen my palate

**Backend**
 - As a website visitor, I want to be able to get a recipe after inputting the ingredients I have.
 - As a user, I want to be able to have a recipe list and save new recipes to the list
 -    As someone with allergies, I want to know ahead of time if my allergen(s) are in something so that I can avoid it.
-   As a picky eater, I want to be able to save what ingredients I have so that I don’t have to list/check them all off each time.
- As a user, I want to be able to login to the website securely, with my password protected, and my data saved.
# Issues to address
**Frontend**
- Build skeleton of the user interface
- Implement page routing for navigation through the website
- Create the necessary user interaction for searching through the recipe database. This includes filters such as meal size and type.
- Design the user interface to be appealing and easy to navigate for the user.
- Create a login and registration interface for a user profile.

**Backend**
- Create recipe data model for backend
- implement Spoonacular API into our backend by making a GET request "Search Recipes by Ingredients." Create REST API for our recipe model
- Connect backend to MONGODB database for CRUD operations -- allow recipes to be saved on the database and be retrieved properly 
- implement user authentication with backend so that users can login and have password encryption
# Which issues were successfully completed
- For backend, we were able to complete all of the listed issues except for user authentication. This consisted of setting up a recipe data model and setting up a rest API for adding recipes and getting a list of recipes (also incorporated a 3rd-party API for specific feature). Then we connected MONGODB database to our backend to allow recipes to be saved on the database and be retrieved properly.
- For the frontend, we implemented the skeleton of the website layout. Implemented page routing and allowed easy access throughout the website. 
# Which issues weren't completed and why?
- For backend, we were unable to complete user login authentication because setting up the REST API and connecting with mongoDB took longer than expected. We ran into a lot of problems while connecting mongoDB with our recipe struct not being able to be read properly and had to do a lot of debugging. We will get it done by next sprint. 
- For the frontend, the three other issues, appealing to the user, search interaction for the recipe database, and login registration are all in their template forms, there isnt much functionality and styling done to the website yet. 
