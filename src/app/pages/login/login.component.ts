import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AuthService } from 'app/auth.service';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
    form: FormGroup;

    constructor(private fb: FormBuilder, private authService: AuthService) {
        this.form = this.fb.group({
            // sign up info
            email: ['', Validators.required],
            signupUsername: ['', Validators.required],
            signupPassword: ['', Validators.required],
            // sign in info
            signinUsername: ['', Validators.required],
            signinPassword: ['', Validators.required],
        });
    }

    ngOnInit() { }

    login() {
        const val = this.form.value;

        if (val.signinUsername && val.signinPassword) {
            this.authService.login(val.signinUsername, val.signinPassword);
        }
    }

    signup() {
        const vals = this.form.value;

        if (vals.email && vals.signupUsername && vals.signupPassword) {
            this.authService.signup(vals.email, vals.signupUsername, vals.signupPassword);
        }
    }
}