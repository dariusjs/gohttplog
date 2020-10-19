describe('The Home Page', () => {
    it('successfully loads', () => {
      // cy.visit('/')
      cy.request('http://localhost:8080')
      .its('body').should('include', 'Hello')
      // cy.get('h1').should('contain', 'jane.lane')
    })
  })
