import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
import { LoginComponent } from 'app/pages/login/login.component';
import { HttpClient } from "@angular/common/http";

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
        private http: HttpClient) { }

    ngOnInit() { }

    username : string = " ";

    getUserName() {
        this.http.get('http://localhost:8080/api/userRecipe').subscribe(
        (res: any) => {
            console.log(res)
        })
    }
}
