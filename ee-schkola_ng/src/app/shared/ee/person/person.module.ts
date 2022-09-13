import { NgModule } from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

import {PersonRoutingModules} from './person-routing.module';
import {CommonModule} from '@angular/common';
import {TemplateModule} from '../template/template.module';

import {ProfileViewComponent} from './profile/components/profile-view/profile-view.component';
import {PersonViewComponent} from './components/person-view/person-view.component';
import {ChurchViewComponent} from './church/components/church-view/church-view.component';
import {ProfileListComponent} from './profile/components/profile-list/profile-list.component';
import {ChurchListComponent} from './church/components/church-list/church-list.component';
import {PersonNameComponent} from './basics/components/personname/personname.component';
import {AddressComponent} from './basics/components/address/address.component';
import {MaterialModule} from '../template/material.module';

@NgModule({
    declarations: [ProfileViewComponent,
        ProfileListComponent,
        PersonViewComponent,
        ChurchViewComponent,
        ChurchListComponent,
        PersonNameComponent,
        AddressComponent
    ],
    imports: [
        PersonRoutingModules,
        TemplateModule,
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        MaterialModule
    ],
    providers: [

    ],
    exports: [
        ProfileViewComponent
    ]
})
export class PersonModule { }
