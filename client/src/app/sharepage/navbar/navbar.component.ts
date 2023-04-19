import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from 'app/auth.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {
  searchQuery: string = " ";
  
  constructor(public authService: AuthService, private router: Router) { }
  
  onSearch(): void {
    console.log(this.searchQuery);
    this.router.navigate(['/usersearch', this.searchQuery]);
  }
}