/// <reference types="cypress"/>

describe('Search page', () => {
  it('Should visit the search page', () => {
    cy.visit('/search')
    cy.url().should('includes','search')
    cy.contains('Meal Size');
    cy.contains('Meal Type');
  })
})