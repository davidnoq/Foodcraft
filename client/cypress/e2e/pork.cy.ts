/// <reference types="cypress"/>

describe('Pork page', () => {
    it('Should visit the Pork page', () => {
      cy.visit('/pork');
      cy.url().should('includes','pork');
      cy.contains('Generated Recipes');
      
    })
  })