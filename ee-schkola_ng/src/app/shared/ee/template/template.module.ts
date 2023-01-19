import { NgModule } from '@angular/core';
import {CommonModule} from '@angular/common';

import {FormComponent} from './components/form/form.component';
import {ButtonComponent} from './components/button/button.component';
import {PageComponent} from './components/page/page.component';
import {TableComponent} from './components/table/table.component';

import {TemplateRoutingModules} from './template-routing.modules';
import {FormService} from './services/form.service';
import {TableDataService} from './services/data.service';
import {ButtonService} from './services/button.service';
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
import {DateTimeTranslationPipePipe} from '../../pipes/date.pipe';

registerLocaleData(localeEn, 'en');
registerLocaleData(localeDe, 'de');
registerLocaleData(localeEs, 'es');
registerLocaleData(localeFa, 'fa');
registerLocaleData(localeFr, 'fr');
registerLocaleData(localeUr, 'ur');

@NgModule({
    declarations: [FormComponent,
        ButtonComponent,
        PageComponent,
        TableComponent,
        DateTimeTranslationPipePipe
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
        ButtonService,
        FormService,
        DateTimeTranslationPipePipe
    ],
    exports: [
        FormComponent,
        ButtonComponent,
        PageComponent,
        TableComponent,
        DateTimeTranslationPipePipe
    ]
})
export class TemplateModule { }
