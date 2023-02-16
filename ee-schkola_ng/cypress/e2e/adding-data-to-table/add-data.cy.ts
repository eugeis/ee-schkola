///<reference path="../../../src/app/shared/ee/schkola/person/profile/service/profile-data.service.ts"/>
///<reference path="../../../src/app/shared/ee/template/services/data.service.ts"/>

describe('should measure performance', () => {
    const element = {
        address: {code: 'a', suite: 'a', street: 'a', country: 'a', city: 'a'},
        birthName: 'a',
        birthday: new Date(Date.now()),
        church: {church: 'a', services: 'a', member: true},
        contact: {phone: 'a', email: 'a', cellphone: 'a'},
        education: {graduation: {name: 'a', level: 0}, profession: 'a', other: 'a'},
        family: {maritalState: 0, childrenCount: 0, partner: {first: 'a', last: 'a'}},
        gender: 0,
        name: {first: 'a', last: 'a'},
        photo: 'a',
        photoData: undefined,
        findByEmail(email?: string) {
            return undefined;
        },
        findByPhone(phone?: string) {
            return undefined;
        }
    }

    beforeEach(() => {
        // Measure Performance Time Before Having Items on Table
        const beforeHavingItemsOnTable = performance.now();
        cy.visit('http://localhost:4200/person/profile');
        cy.get('#pageName').first().should('have.text', 'PERSON');
        cy.wrap(performance.now()).then(afterReload => {
            cy.log(`Page load before having items on table took ${afterReload - beforeHavingItemsOnTable} ms`);
        })
    })

    it('should measure performance time before and after adding 1000 items to table', () => {

        // Adding 1000 Items on Table and Measure Time Needed
        cy.window().then(win => {
            let items: any = win.profileDataService.retrieveItemsFromCache('profile');
            cy.wrap(items).should('have.length', 0);
            const timeBeforeAddingItems = performance.now();
            for (let i = 0; i < 1000; i++) {
                element.birthName = 'a' + i;
                win.profileDataService.inputElement(element);
            }
            cy.wrap(performance.now()).then(timeAfterAddingItems => {
                cy.log(`It needs ${timeAfterAddingItems - timeBeforeAddingItems} ms to add 1000 items to table`);
            })
            items = win.profileDataService.retrieveItemsFromCache('profile');
            cy.wrap(items).should('have.length', 1000);
        })

        // Visiting Other Page With No Data
        const visitingOtherPage = performance.now();
        cy.visit('http://localhost:4200/finance');
        cy.get('#pageName').first().should('have.text', 'FINANCE');
        cy.wrap(performance.now()).then(afterReload => {
            cy.log(`Page load after having items on table took ${afterReload - visitingOtherPage} ms`);
        })

        // Visiting Back the Page with 1000 Data to Measure
        const afterHavingItemsOnTable = performance.now();
        cy.visit('http://localhost:4200/person/profile');
        cy.get('#pageName').first().should('have.text', 'PERSON');
        cy.wrap(performance.now()).then(afterReload => {
            cy.log(`Page load after having items on table took ${afterReload - afterHavingItemsOnTable} ms`);
        })
    })
})
