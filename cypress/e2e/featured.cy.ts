/// <reference types="cypress"/>

describe('Featured page', () => {
    it('Should visit the featured page', () => {
      cy.visit('/featured');
      cy.url().should('includes','featured');
      cy.contains('Likes');
      
    })
  })