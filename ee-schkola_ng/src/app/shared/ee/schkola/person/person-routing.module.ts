import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

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

const routes: Routes = [
    { path: '', component: PersonViewComponent },
    { path: 'church', component: ChurchListComponent },
    { path: 'church/new', component: ChurchViewComponent },
    { path: 'church/edit/:id', component: ChurchViewComponent },
    { path: 'graduation', component: GraduationListComponent },
    { path: 'graduation/new', component: GraduationViewComponent },
    { path: 'graduation/edit/:id', component: GraduationViewComponent },
    { path: 'profile', component: ProfileListComponent },
    { path: 'profile/new', component: ProfileViewComponent },
    { path: 'profile/edit/:id', component: ProfileViewComponent }
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
})
export class PersonRoutingModules {}





