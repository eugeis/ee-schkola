import { SbAdminCliUpdatePage } from './app.po';
import 'jasmine'

describe('sb-admin-cli-update App', () => {
    let page: SbAdminCliUpdatePage;

    beforeEach(async () => {
        page = new SbAdminCliUpdatePage();
    });

    it('should display page name', async () => {
        await page.navigateToProfilePage();
        expect(await page.getPageName().getWebElement().getText()).toEqual('PERSON');
    });
});
