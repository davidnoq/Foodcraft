import { Component, OnInit } from '@angular/core';
import { AuthService } from '../auth.service';
import { Router } from '@angular/router';

@Component({
    selector: 'app-members',
    templateUrl: './user-accounts.component.html',
    styleUrls: ['./user-accounts.component.css']
})
export class MembersComponent implements OnInit {
    accountData: any;
    constructor(private authService: AuthService, private router: Router) { }

    ngOnInit() {
        this.authService.getAccount().subscribe(
            (res: any) => {
                this.accountData = res;
            }, (err: any) => {
                this.router.navigateByUrl('/login');
            }
        );
    }
}
