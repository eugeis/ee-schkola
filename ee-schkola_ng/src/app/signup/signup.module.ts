import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {SignupRoutingModule} from './signup-routing.module';
import {SignupComponent} from './signup.component';
import {AccountService} from "../shared/services/account.service";
import {FormsModule} from "@angular/forms";
import {AlertComponent} from "../shared/components/alert/alert.component";
import {AlertService} from "../shared/services/alert.service";

@NgModule({
    imports: [
        CommonModule,
        SignupRoutingModule,
        FormsModule,
    ],
    declarations: [SignupComponent],
    providers: [AccountService],
})
export class SignupModule {
}
