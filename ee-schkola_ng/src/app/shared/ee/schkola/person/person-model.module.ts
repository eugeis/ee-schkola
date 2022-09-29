import { NgModule } from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

import {PersonRoutingModules} from './person-routing.module';
import {CommonModule} from '@angular/common';
import {TemplateModule} from '../../template/template.module';
import {MaterialModule} from '../../template/material.module';

import {PersonViewComponent} from './components/view/person-module-view.component';
import {ChurchViewComponent} from '@schkola/person/church/components/view/church-entity-view.component';
import {ChurchListComponent} from '@schkola/person/church/components/list/church-entity-list.component';
import {ChurchFormComponent} from '@schkola/person/church/components/form/church-form.component';
import {GraduationViewComponent} from '@schkola/person/graduation/components/view/graduation-entity-view.component';
import {GraduationListComponent} from '@schkola/person/graduation/components/list/graduation-entity-list.component';
import {GraduationFormComponent} from '@schkola/person/graduation/components/form/graduation-form.component';
import {ProfileViewComponent} from '@schkola/person/profile/components/view/profile-entity-view.component';
import {ProfileListComponent} from '@schkola/person/profile/components/list/profile-entity-list.component';
import {ProfileFormComponent} from '@schkola/person/profile/components/form/profile-form.component';
import {AddressComponent} from '@schkola/person/basics/address/address-basic.component';
import {ChurchInfoComponent} from '@schkola/person/basics/churchinfo/churchinfo-basic.component';
import {ContactComponent} from '@schkola/person/basics/contact/contact-basic.component';
import {EducationComponent} from '@schkola/person/basics/education/education-basic.component';
import {FamilyComponent} from '@schkola/person/basics/family/family-basic.component';
import {PersonNameComponent} from '@schkola/person/basics/personname/personname-basic.component';

@NgModule({
    declarations: [
        PersonViewComponent,
        ChurchViewComponent,
        ChurchListComponent,
        ChurchFormComponent,
        GraduationViewComponent,
        GraduationListComponent,
        GraduationFormComponent,
        ProfileViewComponent,
        ProfileListComponent,
        ProfileFormComponent,
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
        MaterialModule,
    ],
    providers: [],
    exports: [
        ChurchFormComponent,
        GraduationFormComponent,
        ProfileFormComponent,
        AddressComponent,
        ChurchInfoComponent,
        ContactComponent,
        EducationComponent,
        FamilyComponent,
        PersonNameComponent
    ]
})
export class PersonModule {}



