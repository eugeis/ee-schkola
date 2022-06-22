import { NgModule } from '@angular/core';
import {CommonModule} from '@angular/common';

import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatNativeDateModule, MatOptionModule} from '@angular/material/core';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {ReactiveFormsModule} from '@angular/forms';
import {MatExpansionModule} from '@angular/material/expansion';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatSidenavModule} from '@angular/material/sidenav';
import {MatDividerModule} from '@angular/material/divider';
import {MatListModule} from '@angular/material/list';
import {MatTableModule} from '@angular/material/table';
import {MatPaginatorModule} from '@angular/material/paginator';
import {MatTabsModule} from '@angular/material/tabs';
import {MatDatepickerModule} from '@angular/material/datepicker';
import {MatMenuModule} from '@angular/material/menu';

import {FormComponent} from './components/form/form.component';
import {ButtonComponent} from './components/button/button.component';
import {PageComponent} from './components/page/page.component';
import {TableComponent} from './components/table/table.component';

import {TemplateRoutingModules} from './template-routing.modules';
import {ProfileViewService} from '../person/profile/services/profile-view.service';
import {ProfileDataService} from '../person/profile/services/profile-data.service';
import {ChurchViewService} from '../person/church/services/church-view.service';
import {ChurchDataService} from '../person/church/services/church-data.service'
import {TemplateService} from './services/template.service';
import {TableDataService} from './services/data.service';
import {ButtonService} from './services/button.service';

@NgModule({
    declarations: [FormComponent,
        ButtonComponent,
        PageComponent,
        TableComponent,
    ],
    imports: [
        TemplateRoutingModules,
        CommonModule,
        MatInputModule,
        MatFormFieldModule,
        MatSelectModule,
        MatOptionModule,
        MatIconModule,
        MatButtonModule,
        ReactiveFormsModule,
        MatExpansionModule,
        MatGridListModule,
        MatToolbarModule,
        MatSidenavModule,
        MatDividerModule,
        MatListModule,
        MatTableModule,
        MatPaginatorModule,
        MatTabsModule,
        MatDatepickerModule,
        MatNativeDateModule,
        MatMenuModule,
    ],
    providers: [
        ProfileViewService, ProfileDataService, ChurchViewService, ChurchDataService, TableDataService, ButtonService, TemplateService
    ],
    exports: [
        FormComponent,
        ButtonComponent,
        PageComponent,
        TableComponent
    ]
})
export class TemplateModule { }
