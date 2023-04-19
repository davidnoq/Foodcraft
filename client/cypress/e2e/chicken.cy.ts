/// <reference types="cypress"/>

describe('Chicken page', () => {
    it('Should visit the featured page', () => {
      cy.visit('/chicken');
      cy.url().should('includes','chicken');
      cy.contains('Generated Recipes');
      
    })
  })