import { Component, OnInit, ViewChild  } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router, NavigationEnd } from '@angular/router';
import { LoginComponent } from 'app/pages/login/login.component';
import { HttpClient } from "@angular/common/http";
import { Injectable } from '@angular/core';
import { BreakpointObserver } from '@angular/cdk/layout';
import { MatSidenav } from '@angular/material/sidenav';
import { delay, filter } from 'rxjs/operators';
import { MatDividerModule } from '@angular/material/divider';

interface Recipe {
    ID: number;
    Title: string;
    Image: string;
    Likes: number;
}

@Component({
    selector: 'app-members',
    templateUrl: './user-accounts.component.html',
    styleUrls: ['./user-accounts.component.css']
})
export class userAccounts implements OnInit {
    accountData: any;
    constructor(
        private authService: AuthService, 
        private router: Router,
        private http: HttpClient,
        private observer: BreakpointObserver) { }

    username : string = " ";
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
}
