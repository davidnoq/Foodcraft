import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';


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
        this.router.navigateByUrl('/');
    }

    login(username: string, pass: string) {
        const headers = {
            headers: new HttpHeaders({ 'Content-Type': 'application/json', 'Cache-Control': 'no-cache' })
        };

        const data = {
            user: username,
            password: pass
        };

        this.http.post(this.API_URL + '/signin', data, headers).subscribe(
            (res: any) => {
                localStorage.setItem(this.TOKEN_KEY, res.token);
                // navigate to profile page when token is returned
                this.router.navigateByUrl('');
            }
        );
    }

    signup(email: string, username: string, pass: string) {
        const data = {
            email: email,
            user: username,
            password: pass
        };

        this.http.post(this.API_URL + '/signup', data).subscribe(
            (res: any) => {
                localStorage.setItem(this.TOKEN_KEY, res.token);
                // navigate to profile page when token is returned
                this.router.navigateByUrl('');
            }
        );
    }

    getAccount() {
        return this.http.get(this.API_URL + '/account');
    }
}
