import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AuthService } from 'app/auth.service';
import { Router } from '@angular/router';


@Component({
  selector: 'app-chicken',
  templateUrl: './chicken.component.html',
  styleUrls: ['./chicken.component.css']
})
export class ChickenComponent implements OnInit {
  form: FormGroup;
  chickenError: string = '';

  constructor(private fb: FormBuilder, private authService: AuthService, private router: Router) {
    this.form = this.fb.group({
      // sign up info
      chicken: ['', Validators.required],
      
  });
}

ngOnInit() { }

chickenRecipes() {
  const val = this.form.value;

  if (val.chicken)
  {
    this.authService.chickenRecipes(val.chicken);
    
    if (this.router.url != 'http://localhost:4200/') 
    {
      this.chickenError = 'Invalid credentials. Please try again.';
    }
  }

}

}
