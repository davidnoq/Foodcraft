import { AboutComponent } from './about.component';
import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

beforeEach(() => TestBed.configureTestingModule({
  imports: [HttpClientTestingModule], 
  providers: [AboutComponent]
}));

describe('AboutComponent', () => {
  it('mounts', () => {
    // change the viewport of the test
    cy.viewport(1300, 700)
    // mount the component to test
    cy.mount(AboutComponent)
    // check visual text
    cy.contains('About');
    cy.contains('David Noguera');

    
  })

  it('should redirect user to a pop up contact page', () => {
    cy.viewport(1300, 700)
    // mount the component to test
    cy.mount(AboutComponent)
    cy.get('button.btn.btn-contact', {timeout: 30000,}).should('be.visible').click();
   

    
    });
  })
