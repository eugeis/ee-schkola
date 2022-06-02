import { NgModule } from '@angular/core';

import {PersonRoutingModules} from './person-routing.module';
import {CommonModule} from '@angular/common';
import {TemplateModule} from '../template/template.module';

import {ProfileViewComponent} from './profile/components/profile-view/profile-view.component';
import {PersonViewComponent} from './profile/components/person-view/person-view.component';
import {ChurchViewComponent} from './profile/components/church-view/church-view.component';

@NgModule({
    declarations: [ProfileViewComponent, PersonViewComponent, ChurchViewComponent],
    imports: [
        PersonRoutingModules,
        TemplateModule,
        CommonModule,
    ],
    providers: [

    ],
    exports: [
        ProfileViewComponent
    ]
})
export class PersonModule { }
