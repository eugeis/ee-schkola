import {Component, Input, OnInit} from '@angular/core';
import {TranslateService} from '@ngx-translate/core';

@Component({
    selector: 'app-page',
    templateUrl: './page.component.html',
    styleUrls: ['./page.component.scss'],
})

export class PageComponent implements OnInit {
    @Input() pageName: string;
    @Input() pageElement: any[];
    @Input() tabElement: any[];

    constructor(public translate: TranslateService) {
        translate.addLangs(['en', 'de', 'es' , 'fa', 'fr', 'ur']);
        translate.setDefaultLang('en');
    }

    ngOnInit(): void {
    }

    changeLanguage(language: string) {
        this.translate.use(language);

    }
}
