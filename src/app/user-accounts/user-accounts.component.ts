import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
    selector: 'app-members',
    templateUrl: './user-accounts.component.html',
    styleUrls: ['./user-accounts.component.css']
})
export class userAccounts implements OnInit {
    accountData: any;
    constructor(private authService: AuthService, private router: Router) { }

    ngOnInit() { }
}
