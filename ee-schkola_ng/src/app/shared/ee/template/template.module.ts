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

import {FormComponent} from './components/form/form.component';
import {ButtonComponent} from './components/button/button.component';
import {PageComponent} from './components/page/page.component';
import {TableComponent} from './components/table/table.component';

import {FamilyComponent} from '../component/family/family.component';
import {EducationComponent} from '../component/education/education.component';
import {LocationComponent} from '../component/location/location.component';
import {PersonNameComponent} from '../component/personname/personname.component';
import {ContactComponent} from '../component/contact/contact.component';
import {ChurchInfoComponent} from '../component/churchinfo/churchinfo.component';
import {AddressComponent} from '../component/address/address.component';

import {TemplateRoutingModules} from './template-routing.modules';
import {ProfileViewService} from '../person/profile/services/profile-view.service';
import {ProfileDataService} from '../person/profile/services/profile-data.service';
import {ChurchViewService} from '../person/church/services/church-view.service';
import {ChurchDataService} from '../person/church/services/church-data.service';

@NgModule({
    declarations: [FormComponent,
        ButtonComponent,
        PageComponent,
        TableComponent,
        FamilyComponent,
        EducationComponent,
        LocationComponent,
        PersonNameComponent,
        ContactComponent,
        ChurchInfoComponent,
        AddressComponent,
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
    ],
    providers: [
        ProfileViewService, ProfileDataService, ChurchViewService, ChurchDataService
    ],
    exports: [
        FormComponent,
        ButtonComponent,
        PageComponent,
        TableComponent
    ]
})
export class TemplateModule { }
