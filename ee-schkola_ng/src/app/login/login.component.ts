import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {routerTransition} from '../router.animations';
import {AuthenticationService} from "../shared/services/authentication.service";
import {AlertService} from "../shared/services/alert.service";

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
                private authenticationService: AuthenticationService,
                private alertService: AlertService) {
    }

    ngOnInit() {
        // reset login status
        this.authenticationService.logout();

        // get return url from route parameters or default to '/'
        this.returnUrl = this.route.snapshot.queryParams['returnUrl'] || '/dashboard';
    }

    login() {
        this.loading = true;
        this.authenticationService.login(this.model.username, this.model.password)
            .subscribe(function (data) {
                    this.onLoggedin()
                    this.router.navigate([this.returnUrl]);
                }.bind(this), function (error) {
                    this.alertService.error(JSON.stringify(error.json()));
                    this.loading = false;
                }.bind(this)
            );
    }

    onLoggedin() {
        localStorage.setItem('isLoggedin', 'true');
    }
}
