import { Component, Input, OnInit } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from "@angular/common/http";
import { AuthService } from 'app/auth.service';
import { catchError } from 'rxjs/operators';
import { throwError } from 'rxjs';

// creating item for recipes and declared variables
interface recipes {
  name: string;
  ingredients: string;
}

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
  ) {
    this.generateTicks();
  }

  number: number = 1;

  increment() {
    this.number++;
  }

  decrement() {
    if (this.number > 1) {
      this.number--;
    }
  }

  min = 0;
  max = 100;
  step = 10;
  value = 50;
  ticks = [];

  generateTicks() {
    const numTicks = (this.max - this.min) / this.step + 1;
    const tickSize = 100 / (numTicks - 1);
    for (let i = 0; i < numTicks; i++) {
      const tickValue = this.min + i * this.step;
      const tickLeft = i === 0 ? 0 : (i === numTicks - 1 ? 100 : i * tickSize);
    }
  }

  onChange(event: any) {
    console.log(this.value);
  }

  breakfastClicked = false;
  lunchClicked = false;
  dinnerClicked = false;
  snackClicked = false;
  dessertClicked = false;

  mealType(buttonId: string) {
    const button = document.getElementById(buttonId);
    if (button) {
      if (button.classList.contains("clicked")) {
        button.classList.remove("clicked");
        switch (buttonId) {
          case "Breakfast":
            this.breakfastClicked = false;
            break;
          case "Lunch":
            this.lunchClicked = false;
            break;
          case "Dinner":
            this.dinnerClicked = false;
            break;
          case "Snack":
            this.snackClicked = false;
            break;
          case "Dessert":
            this.dessertClicked = false;
            break;
        }
      } else {
        button.classList.add("clicked");
        switch (buttonId) {
          case "Breakfast":
            this.breakfastClicked = true;
            break;
          case "Lunch":
            this.lunchClicked = true;
            break;
          case "Dinner":
            this.dinnerClicked = true;
            break;
          case "Snack":
            this.snackClicked = true;
            break;
          case "Dessert":
            this.dessertClicked = true;
            break;
        }
      }
    }
  }

  ngOnInit() {
    const slider = document.querySelector('.slider') as HTMLElement;
    const range = slider.querySelector('.slider-range') as HTMLInputElement;
    const tickMarks = slider.querySelectorAll('.tick-mark') as NodeListOf<HTMLElement>;

    range.addEventListener('input', () => {
      const value = parseInt(range.value);
      const position = value * slider.clientWidth;
      const ticks = Array.from(tickMarks);

      ticks.forEach((tick) => {
        tick.classList.toggle('active', parseInt(tick.innerText) === value);
      });

      range.style.setProperty('--thumb-position', `${position}px`);
    });

    range.dispatchEvent(new Event('input'));
  }

  OntheGoClicked = false;
  MinClicked = false;
  HrClicked = false;
  HrsClicked = false;

  mealTime(buttonId: string) {
    const button = document.getElementById(buttonId);
    if (button) {
      if (button.classList.contains("clicked")) {
        button.classList.remove("clicked");
        switch (buttonId) {
          case "On the Go":
            this.OntheGoClicked = false;
            break;
          case "~30 min":
            this.MinClicked = false;
            break;
          case "~1 Hr":
            this.HrClicked = false;
            break;
          case "~2 Hrs":
            this.HrsClicked = false;
            break;
        }
      } else {
        button.classList.add("clicked");
        switch (buttonId) {
          case "On the Go":
            this.OntheGoClicked = true;
            break;
          case "~30 min":
            this.MinClicked = true;
            break;
          case "~1 Hr":
            this.HrClicked = true;
            break;
          case "~2 Hrs":
            this.HrsClicked = true;
            break;
        }
      }
    }
  }

  selectedMealTypes: string[] = [];
  selectedCuisines: string[] = [];
  selectedIngredients: Ingredient[] = [];
  ingredientlist: string[] = [];
  recipes: string[] = [];

  // filter ingredients and retrieve what the user selects
  ingredient: Ingredient[] = [
    { name: 'Salt', selected: false },
    { name: 'Sugar', selected: false },
    { name: 'Flour', selected: false },
    { name: 'Tomato', selected: false },
    { name: 'Chicken', selected: false },
    { name: 'Rice', selected: false }
  ];

  searchTerm: string = '';
  searchIngredients() {
    return this.ingredient.filter(ingredient =>
      ingredient.name.toLowerCase().includes(this.searchTerm.toLowerCase())
    );
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
    console.log('Selected ingredient names:', this.ingredientlist);
  }

  headers = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json', 'Cache-Control': 'no-cache' })
  };

  // backend requests
  apiUrl = 'http://localhost:8080/api/recipes';
  searchRecipes() {
    const data = {
      ingredientlist: this.ingredientlist
    };
    //const url = `${this.apiUrl}?ingredients=${this.ingredientlist.join(',').toLowerCase()}`;
    this.httpClient.post(this.apiUrl, data).subscribe(
      (res: any) => {
        console.log(res);
      },
      (error) => {
        console.log(error.message); // Make sure to use the correct property name
      })
  }
}
