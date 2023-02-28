import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

interface User {
  name: string;
  email: string;
  password: string;
}

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  loginEmail: string = '';
  loginPassword: string = '';
  signupName: string = '';
  signupEmail: string = '';
  signupPassword: string = '';

  constructor(private http: HttpClient, private router: Router) {}

  login() {
    const body = { email: this.loginEmail, password: this.loginPassword };

    this.http.post('/signin', body).subscribe((response: any) => {
      // store the user information in local storage
      localStorage.setItem('user', JSON.stringify(response.user));

      // navigate to the home page
      this.router.navigate(['/']);
    });
  }

  signup() {
    const user: User = {
      name: this.signupName,
      email: this.signupEmail,
      password: this.signupPassword,
    };

    this.http.post('/signup', user).subscribe((response: any) => {
      // store the user information in local storage
      localStorage.setItem('user', JSON.stringify(response.user));

      // navigate to the home page
      this.router.navigate(['/']);
    });
  }
}
