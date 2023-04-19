/// <reference types="cypress"/>

describe('Search for Recipe', () => {
    it('find recipe', () => {
        cy.visit('/login')
        cy.viewport(1200, 800)
        cy.get('input[id=signinUsername]').type('admin');
        cy.get('input[id=signinPassword]').type('password');
        cy.get('[id=loginButton]').contains('Login')
        cy.get('[id=loginButton]').click()
        cy.contains('Profile')
        cy.visit('/search')
        cy.contains('Salt').click()
        cy.get('[id^=clear]').click()
        cy.contains('Sugar').click()
        cy.get('[id^=craft]').click()
        cy.contains('Creamy Lemon Popsicles').click()
        cy.contains('Likes')
        cy.contains('Instructions')
        cy.contains('Close').click()
        cy.get('[id^=like]').click()

        //check if the profile contains recipe
        cy.contains('Profile').click()
        cy.contains('Creamy Lemon Popsicles')

        //check clear all recipes button
        cy.get('[id^=myProfile]').click()
        cy.contains('Clear All Recipes').click()
        cy.contains('All recipes have been cleared')
        cy.contains('My Recipes').click()
        cy.contains('There are no recipes to display')
    })
})