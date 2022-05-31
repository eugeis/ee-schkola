import { NgModule } from '@angular/core';
import {CommonModule} from '@angular/common';

import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatOptionModule} from '@angular/material/core';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {ReactiveFormsModule} from '@angular/forms';
import {MatExpansionModule} from '@angular/material/expansion';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatSidenavModule} from '@angular/material/sidenav';
import {MatDividerModule} from '@angular/material/divider';
import {MatListModule} from '@angular/material/list';

import {FormComponent} from './components/form/form.component';
import {ButtonComponent} from './components/button/button.component';
import {PageComponent} from './components/page/page.component';

import {FamilyComponent} from '../component/family/family.component';
import {EducationComponent} from '../component/education/education.component';
import {LocationComponent} from '../component/location/location.component';
import {PersonNameComponent} from '../component/personname/personname.component';
import {ContactComponent} from '../component/contact/contact.component';
import {ChurchInfoComponent} from '../component/churchinfo/churchinfo.component';
import {AddressComponent} from '../component/address/address.component';

import {TemplateRoutingModules} from './template-routing.modules';
import {ProfileViewService} from '../profile/services/profile-view.service';

@NgModule({
    declarations: [FormComponent,
        ButtonComponent,
        PageComponent,
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
        MatListModule
    ],
    providers: [
        ProfileViewService
    ],
    exports: [
        FormComponent,
        ButtonComponent,
        PageComponent
    ]
})
export class TemplateModule { }
