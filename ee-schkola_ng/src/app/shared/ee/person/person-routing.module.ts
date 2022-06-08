import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import {ProfileViewComponent} from './profile/components/profile-view/profile-view.component';
import {PersonViewComponent} from './components/person-view/person-view.component';
import {ChurchViewComponent} from './church/components/church-view/church-view.component';
import {ProfileListComponent} from './profile/components/profile-list/profile-list.component';
import {ChurchListComponent} from './church/components/church-list/church-list.component';

const routes: Routes = [
    { path: '', component: PersonViewComponent },
    { path: 'profile', component: ProfileListComponent },
    { path: 'profile/new', component: ProfileViewComponent },
    { path: 'church', component: ChurchListComponent },
    { path: 'church/new', component: ChurchViewComponent },
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
})
export class PersonRoutingModules {}
