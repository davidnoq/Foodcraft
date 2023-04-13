import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';
import { LoginComponent } from 'app/pages/login/login.component';

@Component({
    selector: 'app-members',
    templateUrl: './user-accounts.component.html',
    styleUrls: ['./user-accounts.component.css']
})
export class userAccounts implements OnInit {
    accountData: any;
    constructor(private authService: AuthService, private router: Router) { }

    ngOnInit() { }

    save(email: string, username: string) {
        username = username;
        email = email;
    }
}
