import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
import { HttpClient } from "@angular/common/http";
import { BreakpointObserver } from '@angular/cdk/layout';
import { RecipeDialogComponent } from 'app/recipe-dialog/recipe-dialog.component';
import { MatDialog, MAT_DIALOG_DATA, MatDialogConfig, MatDialogRef } from '@angular/material/dialog';

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
    selector: 'app-members',
    templateUrl: './user-accounts.component.html',
    styleUrls: ['./user-accounts.component.css']
})
export class userAccounts implements OnInit {
    accountData: any;
    isDetailsDialogOpen: boolean = false;
    constructor(
        private authService: AuthService, 
        private router: Router,
        private http: HttpClient,
        private observer: BreakpointObserver,
        private dialog: MatDialog) { 
            this.recipeClick = {
                ID: 0,
                Title: "",
                Image: "",
                Likes: 0,
                instructions: ""
            };
        }

    username : string = " ";
    recipeClick: recipeClick;
    recipeList: Recipe[] = [];
    noRecipes = false;
    
    ngOnInit() {
        this.http.get('http://localhost:8080/api/user').subscribe(
        (res: any) => {
            this.username = res.username;
        })

        this.http.get<Recipe[]>('http://localhost:8080/api/recipes').subscribe(recipes => {
            if(recipes != null) {
                this.recipeList = recipes.map(recipe => {
                    const { ID, Title, Image, Likes } = recipe;
                    return { ID, Title, Image, Likes };
                });
            } else if (recipes == null) {
                this.noRecipes = true;
                this.recipeList = [];
            }
        });
    }

    showRecipes = true;
    showProfile = false;

    getRecipes() {
        this.http.get('http://localhost:8080/api/recipes').subscribe(
        (res: any) => {
            if (res != null) {
                console.log(res)
            } else if (res == null) {
                this.noRecipes = true;
            }
        })
    }

    toggleProfile() {
        if (this.showRecipes == true) {
            this.showRecipes = false;
            this.showProfile = true;
        } else {
            this.showProfile = true;
        }
    }

    toggleRecipes() {
        if (this.showProfile == true) {
            this.showRecipes = true;
            this.showProfile = false;
        } else {
            this.showRecipes = true;
        }
    }

    clear = false;

    clearRecipes() {
        this.http.delete('http://localhost:8080/api/recipes').subscribe(
        (res: any) => {
            if (res.message == "All recipes deleted for user") {
                this.clear = true;
                this.recipeList = [];
                this.noRecipes = true;
            } else {
                this.clear = false;
            }
        })
    }

    removeRecipe(ID: number) {
        this.http.delete('http://localhost:8080/api/recipes/' + ID).subscribe(
        (res: any) => {
            window.location.reload();
        })
    }

    // open dialog window

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
}
