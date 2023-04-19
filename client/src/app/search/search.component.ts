import { Component, Input, OnInit, Inject } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from "@angular/common/http";
import { AuthService } from 'app/auth.service';
import { catchError } from 'rxjs/operators';
import { throwError } from 'rxjs';
import { RecipeDialogComponent } from 'app/recipe-dialog/recipe-dialog.component';
import { MatDialog, MAT_DIALOG_DATA, MatDialogConfig, MatDialogRef } from '@angular/material/dialog';

// creating item for recipes and declared variables

interface Ingredient {
  name: string;
  selected: boolean;
}

interface Recipe {
  ID: number;
  Title: string;
  Image: string;
  Likes: number;
  instructions: string;
  liked: boolean;
}

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent {
  recipe: Recipe;
  isDetailsDialogOpen: boolean = false;

  constructor(
    private httpClient: HttpClient,
    private dialog: MatDialog
  ) {
    this.recipe = {
      ID: 0,
      Title: "",
      Image: "",
      Likes: 0,
      instructions: "",
      liked: false
    };
  }

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

  recipeFound = false;

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
        this.recipe.ID = res[1].ID;
        this.recipe.Title = res[1].Title;
        this.recipe.Likes = res[1].Likes;
        this.recipe.Image = res[1].Image;
        this.recipeFound = true;
        this.isLoading = false;
        this.getInstructions();
      },
      (error) => {
        this.isLoading = false;
      }
    )
  }

  getInstructions() {
    this.httpClient.get('http://localhost:8080/api/recipes/' + this.recipe.ID + '/instructions').subscribe(
      (res: any) => {      
        this.recipe.instructions = res.instructions;
      }
    )
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
    });
  }

  openDetails(): void {
    if (!this.isDetailsDialogOpen) { // check if the details dialog is already open
      const dialogConfig = new MatDialogConfig();
      dialogConfig.data = this.recipe; // Pass the recipe data to the dialog component

      const dialogRef = this.dialog.open(RecipeDialogComponent, dialogConfig);

      dialogRef.afterClosed().subscribe(() => {
        this.isDetailsDialogOpen = false; // set the flag to false when the dialog is closed
      });

      this.isDetailsDialogOpen = true; // set the flag to true when the dialog is opened
    }
  }

  closeDetails(): void {
    this.dialog.closeAll(); // Close all open dialogs
    this.isDetailsDialogOpen = false; // set the flag to false when the dialog is closed
  }

  onLikeButtonClick() {
    if(this.recipe.liked == false) {
      this.httpClient.get(this.apiUrl + '/' + this.recipe.ID).subscribe(
        (res: any) => {
          this.recipe.liked = true;
        }
      )
    } else {
      this.httpClient.delete(this.apiUrl + '/' + this.recipe.ID).subscribe(
        (res: any) => {
          this.recipe.liked = false;
        })
      }
  }
}
