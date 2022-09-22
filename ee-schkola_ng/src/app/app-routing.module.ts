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
    { path: 'finance', loadChildren: () => import('./shared/ee/schkola/finance/finance-model.module').then(x => x.FinanceModule) },
    { path: 'library', loadChildren: () => import('./shared/ee/schkola/library/library-model.module').then(x => x.LibraryModule) },
    { path: 'person', loadChildren: () => import('./shared/ee/schkola/person/person-model.module').then(x => x.PersonModule) },
    { path: 'student', loadChildren: () => import('./shared/ee/schkola/student/student-model.module').then(x => x.StudentModule) },
    { path: 'template', loadChildren: () => import('./shared/ee/template/template.module').then(x => x.TemplateModule) },
    { path: 'not-found', loadChildren: () => import('./not-found/not-found.module').then(x => x.NotFoundModule) },
    { path: '**', redirectTo: 'not-found' }
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule { }
