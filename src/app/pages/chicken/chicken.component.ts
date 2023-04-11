import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-chicken',
  templateUrl: './chicken.component.html',
  styleUrls: ['./chicken.component.css']
})
export class ChickenComponent implements OnInit {
  chickenData: any[] = []; // Define a variable to hold the chicken data

  constructor(private http: HttpClient) { }

  ngOnInit() {
    this.fetchChicken(); // Call fetchChicken() method when the component initializes
  }

  fetchChicken() {
    // Make an HTTP GET request to your backend API endpoint to fetch chicken data
    this.http.get<any[]>('/recipes').subscribe(
      (response: any) => {
        this.chickenData = response; // Assign the fetched data to the chickenData variable
      },
      (error: any) => {
        console.error('Error fetching chicken data:', error);
      }
    );
  }
}