import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

// creating item for recipes and declared variables
interface Recipes {
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

  ngOnInit() {
    this.fetchChicken(); // Call fetchChicken() method when the component initializes
  }

  ingredient: Ingredient[] = [{ name: 'Chicken'}]
  recipes: string[] = [];

  apiUrl = 'http://localhost:8080/api/recipes';
 

// backend requests
  fetchChicken() {
    
    const url = `${this.apiUrl}?ingredients=${this.ingredient.join(',').toLowerCase()}`;
    this.http.post<any>(this.apiUrl, this.ingredient).subscribe(
      (response: any) => {
        this.recipes = response; 
        console.log(this.recipes);
      },
      (error: any) => {
        console.error('Error fetching chicken data:', error);
      }
    );
  }
}