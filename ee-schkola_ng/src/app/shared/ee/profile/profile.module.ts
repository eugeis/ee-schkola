import { NgModule } from '@angular/core';
import { ProfileViewComponent } from './components/profile-view/profile-view.component';
import { ProfilePageComponent } from './components/profile-page/profile-page.component';
import { ProfileListComponent } from './components/profile-list/profile-list.component';
import { TemplateModule } from '../template/template.module';
import {ProfileRoutingModules} from './profile-routing.module';
import {CommonModule} from '@angular/common';
import {MatSidenavModule} from '@angular/material/sidenav';

@NgModule({
    declarations: [ProfileViewComponent,
        ProfilePageComponent,
        ProfileListComponent],
    imports: [
        ProfileRoutingModules,
        TemplateModule,
        CommonModule,
        MatSidenavModule,
    ],
    providers: [

    ],
    exports: [
        ProfileViewComponent
    ]
})
export class ProfileModule { }
