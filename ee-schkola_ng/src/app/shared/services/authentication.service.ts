﻿import { Injectable } from '@angular/core';
import { Http, Headers, Response } from '@angular/http';
import { Observable } from 'rxjs/Observable';
import 'rxjs/add/operator/map'

@Injectable()
export class AuthenticationService {
    constructor(private http: Http) { }

    login(usernameOrEmail: string, password: string) {
        return this.http.post('/api/authenticate', JSON.stringify({ usernameOrEmail: usernameOrEmail, password: password }))
            .map((response: Response) => {
                // login successful if there's a jwt token in the response
                let account = response.json();
                if (account && account.token) {
                    // store account details and jwt token in local storage to keep account logged in between page refreshes
                    localStorage.setItem('currentAccount', JSON.stringify(account));
                }

                return account;
            });
    }

    logout() {
        // remove account from local storage to log account out
        localStorage.removeItem('currentAccount');
    }
}
