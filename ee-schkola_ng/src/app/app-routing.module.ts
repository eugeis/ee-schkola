import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { AuthGuard } from './shared';

const routes: Routes = [
    {
        path: '',
        loadChildren: () => import('./layout/layout.module').then(x => x.LayoutModule),
        canActivate: [AuthGuard]
    },
    { path: 'login', loadChildren: () => import('./login/login.module').then(x => x.LoginModule) },
    { path: 'signup', loadChildren: () => import('./signup/signup.module').then(x => x.SignupModule) },
    { path: 'not-found', loadChildren: () => import('./not-found/not-found.module').then(x => x.NotFoundModule) },
    { path: '**', redirectTo: 'not-found' }
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule { }
