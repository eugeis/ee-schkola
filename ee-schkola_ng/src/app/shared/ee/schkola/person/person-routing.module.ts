import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import {PersonViewComponent} from './components/view/person-module-view.component';
import {ChurchViewComponent} from './church/components/view/church-entity-view.component';
import {ChurchListComponent} from './church/components/list/church-entity-list.component';
import {GraduationViewComponent} from './graduation/components/view/graduation-entity-view.component';
import {GraduationListComponent} from './graduation/components/list/graduation-entity-list.component';
import {ProfileViewComponent} from './profile/components/view/profile-entity-view.component';
import {ProfileListComponent} from './profile/components/list/profile-entity-list.component';

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





