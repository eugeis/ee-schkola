import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';

import { LoginRoutingModule } from './login-routing.module';
import { LoginComponent } from './login.component';
import {AuthenticationService} from "../shared/services/authentication.service";

@NgModule({
    imports: [
        CommonModule,
        LoginRoutingModule
    ],
    declarations: [LoginComponent],
    providers: [AuthenticationService]
})
export class LoginModule {
}
