import {browser, by, element} from 'protractor';

export class SbAdminCliUpdatePage {
    navigateTo() {
      return browser.get('/');
    }

    navigateToProfilePage() {
        return browser.get('/person/profile');
    }

    getParagraphText() {
      return element(by.css('app-root h1'));
    }

    getPageName() {
        return element(by.id('pageName'));
    }
}
