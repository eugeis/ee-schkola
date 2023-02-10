describe('should test displayed items', () => {
    beforeEach(() => {
        cy.visit('http://localhost:4200/person/profile')
    })

    it('displays pagename of the page', () => {
        cy.get('#pageName').first().should('have.text', 'PERSON')
    })
})
