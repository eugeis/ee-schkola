import { NgModule } from '@angular/core';

import {PersonRoutingModules} from './person-routing.module';
import {CommonModule} from '@angular/common';
import {TemplateModule} from '../template/template.module';

import {ProfileViewComponent} from './profile/components/profile-view/profile-view.component';
import {PersonViewComponent} from './components/person-view/person-view.component';
import {ChurchViewComponent} from './church/components/church-view/church-view.component';
import {ProfileListComponent} from './profile/components/profile-list/profile-list.component';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {ChurchListComponent} from './church/components/church-list/church-list.component';

@NgModule({
    declarations: [ProfileViewComponent,
        ProfileListComponent,
        PersonViewComponent,
        ChurchViewComponent,
        ChurchListComponent
    ],
    imports: [
        PersonRoutingModules,
        TemplateModule,
        CommonModule,
        MatIconModule,
        MatButtonModule,
    ],
    providers: [

    ],
    exports: [
        ProfileViewComponent
    ]
})
export class PersonModule { }
