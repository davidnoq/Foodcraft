import { HttpClient, HttpHeaders, HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';


interface Ingredient
{
  Amount: number;
  Name: string;
  Unit: string;
}

@Component({
  selector: 'app-featured',
  templateUrl: './featured.component.html',
  styleUrls: ['./featured.component.css']
})
export class FeaturedComponent implements OnInit {
  title= " ";
  image= " ";
  ID = " ";
  likes= 0;
  instructions= " ";
  ingredientsArray:Ingredient[] = [];
  
  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.fetchRecipe().subscribe(
      (response: any) => {
        this.title = response.Title;
        this.image = response.Image;
        this.likes = response.Likes
        this.ID = response.ID;
        console.log(response);
        this.ingredientsArray = response.UsedIngredients;
        console.log(this.ingredientsArray);
        this.fetchInstructions(response.ID).subscribe((response:any) => {
          console.log(response);
          this.instructions = response.instructions;
          
        }, (error: any) => {
          console.error('Error fetching instructions:', error);
        });
      },
      (error: any) => {
        console.error('Error fetching chicken data:', error);
        
      }
    );

    
  }

  apiUrl = 'http://localhost:8080/api/recipes/featured'; 
  instrucURL = 'http://localhost:8080/api/recipes/:ID/instructions';
  httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json' 
    })
  };

  fetchRecipe() {
    
    return this.http.post<FeaturedRecipe>(this.apiUrl, {}, this.httpOptions);
  }
  
  fetchInstructions(recipeID: number)
  {
    const url = this.instrucURL.replace(':ID', String(recipeID));
    return this.http.get(url,this.httpOptions);
  }
}

interface FeaturedRecipe {
  id: number;
  title: string;
  image: string;
  likes: number;
}
