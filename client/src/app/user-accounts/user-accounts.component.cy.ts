import { TestBed } from '@angular/core/testing';
import { userAccounts } from './user-accounts.component';
import { LoginComponent } from 'app/pages/login/login.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { AuthService } from 'app/auth.service';
import { MatDialog, MatDialogModule, MatDialogConfig } from '@angular/material/dialog';
import { RecipeDialogComponent } from 'app/recipe-dialog/recipe-dialog.component';

beforeEach(() => TestBed.configureTestingModule({
  imports: [HttpClientTestingModule, MatDialogModule], 
  providers: [userAccounts, 
    LoginComponent, 
    AuthService, 
    MatDialog,
    RecipeDialogComponent,
    MatDialogConfig],
}));

describe('userProfile', () => {
    it('delete all recipes', () => {
        cy.mount(userAccounts)
        // change the viewport of the test
        cy.viewport(1200, 800)
        // find button 'Profile' and click
        cy.get('[id^=myProfile]').click();
        // check to see if it changes content on page
        cy.contains('My Profile');
        cy.contains('Username:');
        cy.contains('Delete Account');
        // find button 'clear all recipes' and click
        cy.get('[id^=clearAll]').click();
    })

    it('my recipes', () => {
        cy.mount(userAccounts)
        // change the viewport of the test
        cy.viewport(1200, 800)
        // find button 'Profile' and click
        cy.get('[id^=myRecipes]').click();
        // check to see if it changes content on page
        cy.contains('My Recipes');
    })
})