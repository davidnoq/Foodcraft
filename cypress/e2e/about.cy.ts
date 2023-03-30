/// <reference types="cypress"/>

describe('About page', () => {
    it('Should visit the about page', () => {
      cy.visit('/about')
      cy.url().should('includes','about')
      cy.contains('FoodCraft');
      cy.contains('David Noguera');
    })
  })