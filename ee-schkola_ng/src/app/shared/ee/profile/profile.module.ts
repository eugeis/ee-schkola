import { NgModule } from '@angular/core';
import { ProfileViewComponent } from './components/profile-view/profile-view.component';
import { TemplateModule } from '../template/template.module';
import {ProfileRoutingModules} from './profile-routing.module';
import {CommonModule} from '@angular/common';

@NgModule({
    declarations: [
        ProfileViewComponent,
    ],
    imports: [
        ProfileRoutingModules,
        TemplateModule,
        CommonModule
    ],
    providers: [

    ],
    exports: [
        ProfileViewComponent
    ]
})
export class ProfileModule { }
