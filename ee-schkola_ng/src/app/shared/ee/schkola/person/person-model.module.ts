import { NgModule } from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

import {PersonRoutingModules} from './person-routing.module';
import {CommonModule} from '@angular/common';
import {TemplateModule} from '../../template/template.module';
import {MaterialModule} from '../../template/material.module';

import {PersonViewComponent} from './components/view/person-module-view.component';
import {ChurchViewComponent} from './church/components/view/church-entity-view.component';
import {ChurchListComponent} from './church/components/list/church-entity-list.component';
import {GraduationViewComponent} from './graduation/components/view/graduation-entity-view.component';
import {GraduationListComponent} from './graduation/components/list/graduation-entity-list.component';
import {ProfileViewComponent} from './profile/components/view/profile-entity-view.component';
import {ProfileListComponent} from './profile/components/list/profile-entity-list.component';

import {AddressComponent} from './basics/address/address-basic.component';
import {ChurchInfoComponent} from './basics/churchinfo/churchinfo-basic.component';
import {ContactComponent} from './basics/contact/contact-basic.component';
import {EducationComponent} from './basics/education/education-basic.component';
import {FamilyComponent} from './basics/family/family-basic.component';
import {PersonNameComponent} from './basics/personname/personname-basic.component';

@NgModule({
    declarations: [
        PersonViewComponent,
        ChurchViewComponent,
        ChurchListComponent,
        GraduationViewComponent,
        GraduationListComponent,
        ProfileViewComponent,
        ProfileListComponent,
        AddressComponent,
        ChurchInfoComponent,
        ContactComponent,
        EducationComponent,
        FamilyComponent,
        PersonNameComponent
    ],
    imports: [
        PersonRoutingModules,
        TemplateModule,
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        MaterialModule
    ],
    providers: [],
    exports: [
        ChurchViewComponent,
        GraduationViewComponent,
        ProfileViewComponent,
        AddressComponent,
        ChurchInfoComponent,
        ContactComponent,
        EducationComponent,
        FamilyComponent,
        PersonNameComponent
    ]
})
export class PersonModule { }




