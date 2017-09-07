import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {routerTransition} from '../router.animations';
import {AuthService} from "../shared/services/auth.service";
import {AlertService} from "../shared/services/alert.service";
import {AppSettings} from "../shared/services/common.service";

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrls: ['./login.component.scss'],
    animations: [routerTransition()]
})
export class LoginComponent implements OnInit {

    model: any = {};
    loading = false;
    returnUrl: string;

    constructor(private route: ActivatedRoute,
                private router: Router,
                private authService: AuthService,
                private alertService: AlertService) {
    }

    ngOnInit() {
        // reset login status
        this.authService.onLogout();

        // get return url from route parameters or default to '/'
        this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/dashboard';
    }

    login() {
        this.loading = true;
        this.authService.login(this.model.username, this.model.password)
            .subscribe(function (accountToken) {
                    if (accountToken && accountToken.token) {
                        // store account details and jwt token in local storage to keep account logged in between page refreshes
                        localStorage.setItem(AppSettings.CURRENT_ACCOUNT, JSON.stringify(accountToken));
                    }
                    this.router.navigate([this.returnUrl]);
                }.bind(this), function (error) {
                    this.alertService.error(JSON.stringify(error.json()));
                    this.loading = false;
                }.bind(this)
            );
    }
}
