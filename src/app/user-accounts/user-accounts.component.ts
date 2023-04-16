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
    
    ngOnInit() {
        this.http.get('http://localhost:8080/api/user').subscribe(
        (res: any) => {
            this.username = res.username;
        })
    }

    showRecipes = true;
    showProfile = false;

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


}
