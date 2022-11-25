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

@NgModule({
    declarations: [FormComponent,
        ButtonComponent,
        PageComponent,
        TableComponent,
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
        FormService
    ],
    exports: [
        FormComponent,
        ButtonComponent,
        PageComponent,
        TableComponent
    ]
})
export class TemplateModule { }
