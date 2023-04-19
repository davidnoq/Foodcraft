import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
  selector: 'app-recipe-dialog',
  templateUrl: './recipe-dialog.component.html',
  styleUrls: ['./recipe-dialog.component.css']
})
export class RecipeDialogComponent {
  recipe: any; // Define recipe as any type to allow dynamic properties

  constructor(@Inject(MAT_DIALOG_DATA) 
    public data: any) {
    // Access the recipe data from the parent component
    this.recipe = this.data;
  }
}
