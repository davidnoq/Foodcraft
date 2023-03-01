import { TestBed } from '@angular/core/testing';
import { LoginComponent } from './login.component';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AuthService } from 'app/auth.service';

beforeEach(() => TestBed.configureTestingModule({
  imports: [HttpClientTestingModule, RouterTestingModule], 
  providers: [LoginComponent, AuthService]
}));

describe('LoginComponent', () => {
  it('mounts', () => {
    // change the viewport of the test
    cy.viewport(1300, 700)
    // mount the component to test
    cy.mount(LoginComponent)
    // fill out form
    cy.get('input[type="text"').should('be.visible').type('john-doe@example.com')
    cy.get('input[type="password"').should('be.visible').type('password')
  })
})