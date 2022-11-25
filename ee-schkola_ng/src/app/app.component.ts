import { Component } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss']
})
export class AppComponent {
    constructor(private translate: TranslateService) {
        /*translate.addLangs(['de', 'ru', 'en']);
        translate.setDefaultLang('de');
        const browserLang = translate.getBrowserLang();
        translate.use(browserLang.match(/de|ru|en/) ? browserLang : 'de');*/
        translate.setDefaultLang('en');
        translate.use(translate.currentLang);
    }
}
