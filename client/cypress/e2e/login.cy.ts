/// <reference types="cypress"/>

beforeEach(() => {
    cy.visit('/login')
})

describe('Login/Signup page', () => {
    it('Visit the Login/Signup Page', () => {
        cy.visit('/login')
        cy.url().should('includes','login')
        cy.contains('Login');
        cy.contains('Signup');
    })

    it('testing invalid credentials for login', () => {
        cy.get('input[id=signinUsername]').type('testUser');
        cy.get('input[id=signinPassword]').type('testingPass');
        cy.get('[id=loginButton]').contains('Login')
        cy.get('[id=loginButton]').click()
        cy.contains('Invalid credentials. Please try again.')
    }) 

    it('testing navigation of login', () => {
        cy.get('input[id=signinUsername]').type('admin');
        cy.get('input[id=signinPassword]').type('password');
        cy.get('[id=loginButton]').contains('Login')
        cy.get('[id=loginButton]').click()
        cy.url().should('include', '')
        cy.contains('Craft Your Food')
    })

    it('testing signup and logout button', () => {
        cy.get('input[id=email]').type('testing2@email.com');
        cy.get('input[id=signupUsername]').type('testUser2');
        cy.get('input[id=signupPassword]').type('testingPass2');
        cy.get('[id=signupButton]').contains('Signup')
        cy.get('[id=signupButton]').click()
        cy.contains('Credentials already taken. Please try again.')
    })

    it('testing signin after signup', () => {
        cy.get('input[id=signinUsername]').type('testUser2');
        cy.get('input[id=signinPassword]').type('testingPass2');
        cy.get('[id=loginButton]').click()
        cy.url().should('include', '')
    })
})