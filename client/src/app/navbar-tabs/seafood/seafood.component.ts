import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { RecipeDialogComponent } from 'app/recipe-dialog/recipe-dialog.component';
import { MatDialog, MAT_DIALOG_DATA, MatDialogConfig, MatDialogRef } from '@angular/material/dialog';
// Creating interfaces for recipes and ingredients
interface Recipe {
  ID: number;
  Title: string;
  Image: string;
  Likes: number;
}
interface recipeClick {
  ID: number;
  Title: string;
  Image: string;
  Likes: number;
  instructions: string;
}
@Component({
  selector: 'app-pork',
  templateUrl: './seafood.component.html',
  styleUrls: ['./seafood.component.css']
})
export class SeafoodComponent implements OnInit {
  
  accountData: any;
    isDetailsDialogOpen: boolean = false;
  constructor(private http: HttpClient,private dialog: MatDialog) {
    this.recipeClick = {
      ID: 0,
      Title: "",
      Image: "",
      Likes: 0,
      instructions: ""
  };
   }
   ngOnInit() {
    this.fetchSeafood(); 
  }

  //update to single object instead of an array
  ingredientlist = ["Salmon"]; 
  ingredientArray: Recipe[] = [];
  recipeClick: recipeClick;
  apiUrl = 'http://localhost:8080/api/recipes';
  
  //backend requests
  fetchSeafood() {
    const data = {
      ingredientlist: this.ingredientlist
    }
    this.http.post<Recipe[]>(this.apiUrl, data).subscribe(
      (response: Recipe[]) => {
        this.ingredientArray = response;
        
        console.log(this.ingredientArray);
        
      },
      (error: any) => {
        console.error('Error fetching chicken data:', error);
        
      }
    );
  }
  openDetails(): void {
    if (!this.isDetailsDialogOpen) { // check if the details dialog is already open
      const dialogConfig = new MatDialogConfig();
      dialogConfig.data = this.recipeClick; // Pass the recipe data to the dialog component

      const dialogRef = this.dialog.open(RecipeDialogComponent, dialogConfig);

      dialogRef.afterClosed().subscribe(() => {
        this.isDetailsDialogOpen = false; // set the flag to false when the dialog is closed
      });

      this.isDetailsDialogOpen = true; // set the flag to true when the dialog is opened
    }
}

onRecipeCardClick(recipe: Recipe): void {
    if (!this.isDetailsDialogOpen) { // check if the details dialog is already open
        const dialogConfig = new MatDialogConfig();
        recipeClicked: this.recipeClick;

        this.recipeClick.Title = recipe.Title;
        this.recipeClick.Likes = recipe.Likes;
        this.recipeClick.Image = recipe.Image;
        this.recipeClick.ID = recipe.ID;
        this.getInstructions();

        dialogConfig.data = this.recipeClick; // Pass the recipe data to the dialog component
  
        const dialogRef = this.dialog.open(RecipeDialogComponent, dialogConfig);
  
        dialogRef.afterClosed().subscribe(() => {
          this.isDetailsDialogOpen = false; // set the flag to false when the dialog is closed
        });
  
        this.isDetailsDialogOpen = true; // set the flag to true when the dialog is opened
      }
}

getInstructions() {
  this.http.get('http://localhost:8080/api/recipes/' + this.recipeClick.ID + '/instructions').subscribe(
    (res: any) => {      
      this.recipeClick.instructions = res.instructions;
    }
  )
}

removeRecipe(ID: number) {
  this.http.delete('http://localhost:8080/api/recipes/' + ID).subscribe(
  (res: any) => {
      window.location.reload();
  })
}
}
