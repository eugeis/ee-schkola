import {Injectable} from '@angular/core';
import {Headers, Http, RequestOptions, Response} from '@angular/http';

import {Account} from '../ee/schkola/auth/AuthApiBase';
import {AppSettings} from "./common.service";

@Injectable()
export class AccountService {
    emptyOptions = new RequestOptions()
    accounts = `${AppSettings.HOST_API}/auth/accounts/`

    constructor(private http: Http) {
    }

    getAll() {
        return this.http.get(this.accounts, this.jwt()).map((response: Response) => response.json());
    }

    getById(id: number) {
        return this.http.get(this.accounts + id, this.jwt()).map((response: Response) => response.json());
    }

    create(account: Account) {
        let observable = this.http.post(this.accounts + account.id, account, this.jwt());
        return observable.map((response: Response) => {
            return response.json();
        })
    }

    update(account: Account) {
        return this.http.put(this.accounts + account.id, account, this.jwt()).map((response: Response) => response.json());
    }

    delete(id: number) {
        return this.http.delete(this.accounts + id, this.jwt()).map((response: Response) => response.json());
    }

    // private helper methods

    private jwt() {
        if (!AppSettings.HTTP_COOKIE) {
            // create authorization header with jwt token
            let currentAccount = JSON.parse(localStorage.getItem(AppSettings.CURRENT_ACCOUNT));
            if (currentAccount && currentAccount.token) {
                let headers = new Headers({'Authorization': 'Bearer ' + currentAccount.token});
                return new RequestOptions({headers: headers});
            }
        }
        return this.emptyOptions
    }
}
