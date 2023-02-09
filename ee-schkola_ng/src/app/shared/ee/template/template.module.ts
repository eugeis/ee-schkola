import { NgModule } from '@angular/core';
import {CommonModule} from '@angular/common';

import {PageComponent} from './components/page/page.component';

import {TemplateRoutingModules} from './template-routing.modules';
import {TableDataService} from './services/data.service';
import {MaterialModule} from './material.module';
import {TranslateModule, TranslateService} from '@ngx-translate/core';
import {TemplateTranslateService} from '@template/services/translate.service';
import { registerLocaleData } from '@angular/common';
import localeEn from '@angular/common/locales/en';
import localeDe from '@angular/common/locales/de';
import localeEs from '@angular/common/locales/es';
import localeFa from '@angular/common/locales/fa';
import localeFr from '@angular/common/locales/fr';
import localeUr from '@angular/common/locales/ur';
import {DateTimeTranslationPipe} from '../../pipes/date.pipe';

registerLocaleData(localeEn, 'en');
registerLocaleData(localeDe, 'de');
registerLocaleData(localeEs, 'es');
registerLocaleData(localeFa, 'fa');
registerLocaleData(localeFr, 'fr');
registerLocaleData(localeUr, 'ur');

@NgModule({
    declarations: [
        PageComponent,
        DateTimeTranslationPipe
    ],
    imports: [
        TemplateRoutingModules,
        CommonModule,
        MaterialModule,
        TranslateModule
    ],
    providers: [
        {provide: TranslateService, useClass: TemplateTranslateService },
        TableDataService,
        DateTimeTranslationPipe
    ],
    exports: [
        PageComponent,
        DateTimeTranslationPipe
    ]
})
export class TemplateModule { }
