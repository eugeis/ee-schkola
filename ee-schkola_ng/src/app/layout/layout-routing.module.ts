import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LayoutComponent } from './layout.component';

const routes: Routes = [
    {
        path: '', component: LayoutComponent,
        children: [
            { path: 'dashboard', loadChildren: () => import('./dashboard/dashboard.module').then(x => x.DashboardModule) },
            { path: 'charts', loadChildren: () => import('./charts/charts.module').then(x => x.ChartsModule) },
            { path: 'tables', loadChildren: () => import('./tables/tables.module').then(x => x.TablesModule) },
            { path: 'forms', loadChildren: () => import('./form/form.module').then(x => x.FormModule) },
            { path: 'bs-element', loadChildren: () => import('./bs-element/bs-element.module').then(x => x.BsElementModule) },
            { path: 'grid', loadChildren: () => import('./grid/grid.module').then(x => x.GridModule) },
            { path: 'components', loadChildren: () => import('./bs-component/bs-component.module').then(x => x.BsComponentModule) },
            { path: 'blank-page', loadChildren: () => import('./blank-page/blank-page.module').then(x => x.BlankPageModule) },
        ]
    }
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class LayoutRoutingModule { }
