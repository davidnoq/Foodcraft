/// <reference types="cypress"/>

describe('Beef page', () => {
    it('Should visit the beef page', () => {
      cy.visit('/beef');
      cy.url().should('includes','beef');
      cy.contains('Generated Recipes');
      
    })
  })