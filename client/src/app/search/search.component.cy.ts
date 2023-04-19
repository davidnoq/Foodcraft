import { TestBed } from '@angular/core/testing';
import { SearchComponent } from './search.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { AuthService } from 'app/auth.service';
import { MatDialog, MatDialogModule, MatDialogConfig } from '@angular/material/dialog';
import { RecipeDialogComponent } from 'app/recipe-dialog/recipe-dialog.component';

beforeEach(() => TestBed.configureTestingModule({
  imports: [HttpClientTestingModule, MatDialogModule], 
  providers: [SearchComponent,
    AuthService, 
    MatDialog,
    RecipeDialogComponent,
    MatDialogConfig]
}));

describe('SearchComponent', () => {
  it('mounts', () => {
    // change the viewport of the test
    cy.viewport(1300, 700)
    // mount the component to test
    cy.mount(SearchComponent)
    // check text
    cy.contains('Type in Ingredients')
    // type ingredient thats not within list
    cy.get('input[id=search]').type('bread')
    // click ingredients listed
    cy.contains('Salt').click()
    // check all
    cy.get('[id^=checkbox]').check()
    // uncheck all
    cy.get('[id^=checkbox]').uncheck()
    // contains a craft recipe button
    cy.contains('Craft Recipe')
  })
})