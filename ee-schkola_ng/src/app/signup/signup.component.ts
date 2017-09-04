import {Component, OnInit} from '@angular/core';
import {routerTransition} from '../router.animations';
import {Router} from "@angular/router";
import {AccountService} from "../shared/services/account.service";
import {AlertService} from "../shared/services/alert.service";
import {Account} from "../shared/ee/schkola/auth/AuthApiBase";
import {PersonName} from "../shared/ee/schkola/SharedApiBase";

@Component({
    selector: 'app-signup',
    templateUrl: './signup.component.html',
    styleUrls: ['./signup.component.scss'],
    animations: [routerTransition()]
})
export class SignupComponent implements OnInit {

    model: Account = new Account;
    passwordRepeat: String;
    loading = false;

    ngOnInit() {
        this.model.name = new PersonName
    }

    constructor(private router: Router,
                private accountService: AccountService,
                private alertService: AlertService) {
    }

    register() {
        this.loading = true;
        this.model.id = this.model.username
        this.accountService.create(this.model)
            .subscribe(
                data => {
                    this.alertService.success('Registration successful', true);
                    this.router.navigate(['/login']);
                },
                error => {
                    this.alertService.error(error);
                    this.loading = false;
                });
    }
}
