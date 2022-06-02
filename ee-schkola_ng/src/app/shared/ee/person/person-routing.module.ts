import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import {ProfileViewComponent} from './profile/components/profile-view/profile-view.component';
import {PersonViewComponent} from './profile/components/person-view/person-view.component';
import {ChurchViewComponent} from './profile/components/church-view/church-view.component';

const routes: Routes = [
    { path: '', component: PersonViewComponent },
    { path: 'profile', component: ProfileViewComponent },
    { path: 'church', component: ChurchViewComponent },
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
})
export class PersonRoutingModules {}
