import { NgModule } from '@angular/core';
import {ReactiveFormsModule} from '@angular/forms';

import {PersonRoutingModules} from './person-routing.module';
import {CommonModule} from '@angular/common';
import {TemplateModule} from '../template/template.module';

import {ProfileViewComponent} from './profile/components/profile-view/profile-view.component';
import {PersonViewComponent} from './components/person-view/person-view.component';
import {ChurchViewComponent} from './church/components/church-view/church-view.component';
import {ProfileListComponent} from './profile/components/profile-list/profile-list.component';
import {ChurchListComponent} from './church/components/church-list/church-list.component';

import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatExpansionModule} from '@angular/material/expansion';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatInputModule} from '@angular/material/input';

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
        MatExpansionModule,
        MatFormFieldModule,
        MatSelectModule,
        MatInputModule,
        ReactiveFormsModule,
    ],
    providers: [

    ],
    exports: [
        ProfileViewComponent
    ]
})
export class PersonModule { }
