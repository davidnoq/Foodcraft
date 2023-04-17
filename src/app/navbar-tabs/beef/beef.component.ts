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
  selector: 'app-beef',
  templateUrl: './beef.component.html',
  styleUrls: ['./beef.component.css']
})
export class BeefComponent implements OnInit {
  
  constructor(private http: HttpClient) { }

  //call fetchChicken() method when the component initializes
  ngOnInit() {
    this.fetchBeef(); 
  }

  //update to single object instead of an array
  ingredientlist = ["Beef"]; 
  title: string = " ";
  image: string = " ";
  likes: string = " ";


  apiUrl = 'http://localhost:8080/api/recipes';
  
  //backend requests
  fetchBeef() {
    const data = {
      ingredientlist: this.ingredientlist
    }
    this.http.post(this.apiUrl, data).subscribe(
      (response: any) => {
        this.title = response.Title;
        this.image = response.Image;
        this.likes = response.Likes
        console.log(response);
      },
      (error: any) => {
        console.error('Error fetching steak data:', error);
        
      }
    );
  }
}
