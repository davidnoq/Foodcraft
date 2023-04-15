import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { retry, catchError } from 'rxjs/operators';
import { LoginComponent } from './pages/login/login.component';
import { Form, FormGroup } from '@angular/forms';
import { NavbarComponent } from './sharepage/navbar/navbar.component';
import { ChickenComponent } from './navbar-tabs/chicken/chicken.component';

@Injectable()
export class AuthService {

    API_URL = 'http://localhost:8080/api';
    TOKEN_KEY = 'token';

    constructor(private http: HttpClient, private router: Router) { }

    get token() {
        return localStorage.getItem(this.TOKEN_KEY);
    }

    get isAuthenticated() {
        return !!localStorage.getItem(this.TOKEN_KEY);
    }

    logout() {
        localStorage.removeItem(this.TOKEN_KEY);
        this.router.navigateByUrl('');
    }

    headers = {
        headers: new HttpHeaders({ 'Content-Type': 'application/json', 'Cache-Control': 'no-cache' })
    };

    login(username: string, password: string) {
        const data = {
            username: username,
            password: password
        };

        this.http.post(this.API_URL + '/signin', data, this.headers).subscribe(
            (res: any) => {
                console.log(res.token);
                localStorage.setItem(this.TOKEN_KEY, res.token);
                // navigate to profile page when token is returned
                this.router.navigateByUrl('');
            }, 
            (error) => {
                console.error
            }
        );
    }

    signup(email: string, username: string, password: string) {
        const data = {
            username: username,
            password: password
        };

        this.http.post(this.API_URL + '/signup', data, this.headers).subscribe(
            (res: any) => {
                localStorage.setItem(this.TOKEN_KEY, res.token);
                // navigate to profile page when token is returned
                this.router.navigateByUrl('');
            }, 
            (error) => {
                console.error
            }
        );
    }

    getAccount() {
        
    }

    
}
