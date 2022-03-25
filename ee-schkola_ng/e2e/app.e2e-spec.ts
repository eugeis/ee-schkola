import { SbAdminCliUpdatePage } from './app.po';
import '@types/jest';

describe('sb-admin-cli-update App', () => {
  let page: SbAdminCliUpdatePage;

  beforeEach(() => {
    page = new SbAdminCliUpdatePage();
  });

  it('should display welcome message', async () => {
    await page.navigateTo();
    await expect(page.getParagraphText().getText).toEqual('Welcome to app!!');
  });
});
