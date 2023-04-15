import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

// Creating interfaces for recipes and ingredients
interface Recipe {
  title: string;
  image: string;
}

interface Ingredient {
  ingredientlist: string;
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
  ingredientlist = ["Chicken"]; 
  title: string = " ";
  image: string = " ";


  apiUrl = 'http://localhost:8080/api/recipes';
  
  //backend requests
  fetchChicken() {
    const data = {
      ingredientlist: this.ingredientlist
    }
    this.http.post(this.apiUrl, data).subscribe(
      (response: any) => {
        this.title = response.Title;
        this.image = response.Image;
        console.log(response);
      },
      (error: any) => {
        console.error('Error fetching chicken data:', error);
        
      }
    );
  }
}
