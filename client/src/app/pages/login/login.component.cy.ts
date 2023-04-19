import { TestBed } from '@angular/core/testing';
import { LoginComponent } from './login.component';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AuthService } from 'app/auth.service';

beforeEach(() => TestBed.configureTestingModule({
  imports: [HttpClientTestingModule, RouterTestingModule], 
  providers: [LoginComponent, AuthService]
}));

beforeEach(() => {
  cy.mount(LoginComponent)
})

describe('LoginComponent', () => {
  it('signup test', () => {
    cy.contains('Signup')
    cy.get('input[id=email]').type('testing@email.com');
    cy.get('input[id=signupUsername]').type('testUser');
    cy.get('input[id=signupPassword]').type('testingPass');
    cy.get('[id=signupButton]').contains('Signup')
  })

  it('signin test', () => {
    cy.contains('Login')
    cy.get('input[id=signinUsername]').type('testUser');
    cy.get('input[id=signinPassword]').type('testingPass');
    cy.get('[id=loginButton]').contains('Login')
  })
})

