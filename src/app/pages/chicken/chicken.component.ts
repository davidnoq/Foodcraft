import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

// Creating interfaces for recipes and ingredients
interface Recipe {
  title: string;
  image: string;
}

interface Ingredient {
  name: string;
}

@Component({
  selector: 'app-chicken',
  templateUrl: './chicken.component.html',
  styleUrls: ['./chicken.component.css']
})
export class ChickenComponent implements OnInit {
  
  constructor(private http: HttpClient) { }

  //call fetchChicken() method when the component initializes
  ngOnInit() {
    this.fetchChicken(); 
  }

  //update to single object instead of an array
  ingredient: Ingredient = { name: 'Chicken' }; 
  recipes: Recipe[] = []; 

  apiUrl = 'http://localhost:8080/api/recipes';
  
  //backend requests
  fetchChicken() {
    this.http.post<any[]>(this.apiUrl, [this.ingredient]).subscribe(
      (response: any[]) => {
        this.recipes = response;
        console.log(this.recipes);
      },
      (error: any) => {
        console.error('Error fetching chicken data:', error);
        console.log(this.ingredient);
      }
    );
  }
}
