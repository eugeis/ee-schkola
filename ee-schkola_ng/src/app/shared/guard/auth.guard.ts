import {Injectable} from '@angular/core';
import {CanActivate, Router} from '@angular/router';
import {AppSettings} from '../services/common.service';

@Injectable()
export class AuthGuard implements CanActivate {

    constructor(private router: Router) {
    }

    canActivate() {
        if (localStorage.getItem(AppSettings.CURRENT_ACCOUNT)) {
            return true;
        }

        this.router.navigate(['/login']);
        return false;
    }
}
