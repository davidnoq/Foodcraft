import { Component, Input, OnInit } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from "@angular/common/http";
import { AuthService } from 'app/auth.service';
import { catchError } from 'rxjs/operators';
import { throwError } from 'rxjs';

// creating item for recipes and declared variables

interface Ingredient {
  name: string;
  selected: boolean;
}

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent {

  constructor(
    private httpClient: HttpClient
  ) {}

  ngOnInit() {}

  selectedIngredients: Ingredient[] = [];
  ingredientlist: string[] = [];
  recipes: string[] = [];
  hasMatch = true;

  // filter ingredients and retrieve what the user selects
  ingredient: Ingredient[] = [
    { name: 'Salt', selected: false },
    { name: 'Sugar', selected: false },
    { name: 'Flour', selected: false },
    { name: 'Tomato', selected: false },
    { name: 'Chicken', selected: false },
    { name: 'Rice', selected: false },
    { name: 'Beef', selected: false},
    { name: 'Pork', selected: false},
    { name: 'Pepper', selected: false},
    { name: 'Orange', selected: false},
    { name: 'Apple', selected: false}
  ];

  searchTerm: string = '';
  isLoading = false;

  searchIngredients() {
    const ingredientList = this.ingredient.filter(ingredient =>
      ingredient.name.toLowerCase().includes(this.searchTerm.toLowerCase())
    );

    return ingredientList.sort((a, b) => {
      if (a.selected && !b.selected) {
        return -1;
      } else if (!a.selected && b.selected) {
        return 1;
      } else {
        return 0;
      }
    });
  }

  onIngredientSelected(ingredient: Ingredient) {
    ingredient.selected = !ingredient.selected;
    this.recipeError = false;
    if (ingredient.selected) {
      this.selectedIngredients.push(ingredient);
    } else {
      const index = this.selectedIngredients.findIndex(selected => selected.name === ingredient.name);
      this.selectedIngredients.splice(index, 1);
    }
    this.ingredientlist = this.selectedIngredients.map(selected => selected.name);
  }

  headers = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json', 'Cache-Control': 'no-cache' })
  };

  title: string = "";
  likes: number = 0;
  imageSrc: string = "";
  recipeFound = false;
  recipeError = false;

  // backend requests
  apiUrl = 'http://localhost:8080/api/recipes';
  searchRecipes() {
    const data = {
      ingredientlist: this.ingredientlist
    };
    this.isLoading = true;
    //const url = `${this.apiUrl}?ingredients=${this.ingredientlist.join(',').toLowerCase()}`;
    this.httpClient.post(this.apiUrl, data).subscribe(
      (res: any) => {
        this.title = res.Title;
        this.likes = res.Likes;
        this.imageSrc = res.Image;
        this.recipeFound = true;
        this.isLoading = false;
      },
      (error) => {
        this.recipeError = true;
        this.isLoading = false;
      })
  }

  addIngredient(ingredient: string) {
    const ingredientName = ingredient.charAt(0).toUpperCase() + ingredient.slice(1);
    this.ingredient.push({ name: ingredientName, selected: true });
    this.ingredientlist.push(ingredientName);
    this.ingredient = [...this.ingredient];
    this.searchTerm = '';
  }

  get sortedIngredients() {
    // Sort the ingredients based on whether they are checked or not
    return this.ingredient.sort((a, b) => {
      if (a.selected && !b.selected) {
        return -1;
      } else if (!a.selected && b.selected) {
        return 1;
      } else {
        return 0;
      }
    });
  }

  clearSelected() {
    this.ingredient.forEach((ingredient) => {
      ingredient.selected = false;
      this.recipeError = false;
    });
  }
}
