import { TestBed } from '@angular/core/testing';
import { SearchComponent } from './search.component';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

beforeEach(() => TestBed.configureTestingModule({
  imports: [HttpClientTestingModule], 
  providers: [SearchComponent]
}));

describe('SearchComponent', () => {
  it('mounts', () => {
    // change the viewport of the test
    cy.viewport(1300, 700)
    // mount the component to test
    cy.mount(SearchComponent)
    // check visual text
    cy.contains('Meal Size');
    cy.contains('Meal Type');
    // testing whether there is a button with the test 'Breakfast'
    cy.get('button').should('contains.text', 'Breakfast')
    // click the breakfast button
    cy.get('[id^=Breakfast]').click()
    // check if the button changed
    cy.get('[id^=Breakfast]').should('have.css', 'background-color', 'rgb(234, 192, 134)')
    // type in the search field by typing tomato
    cy.get('input[type=text]').type('tomato')
    // test increment for meal size
    cy.get('#increment').click()
    cy.get('#mealSize').should('contains.text', '2')
  })
})