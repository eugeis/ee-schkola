import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import {TemplateComponent} from './components/example/template.component';

const routes: Routes = [
    { path: '', component: TemplateComponent }
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class TemplateRoutingModules { }
