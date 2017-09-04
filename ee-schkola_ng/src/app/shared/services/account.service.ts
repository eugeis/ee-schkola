import {Injectable} from '@angular/core';
import {Headers, Http, RequestOptions, Response} from '@angular/http';

import {Account} from '../ee/schkola/auth/AuthApiBase';

@Injectable()
export class AccountService {
    constructor(private http: Http) {

    }

    getAll() {
        return this.http.get('http://localhost:8080/auth/accounts', this.jwt()).map((response: Response) => response.json());
    }

    getById(id: number) {
        return this.http.get('http://localhost:8080/auth/accounts/' + id, this.jwt()).map((response: Response) => response.json());
    }

    create(account: Account) {
        let observable = this.http.post('http://localhost:8080/auth/accounts/' + account.id, account, this.jwt());
        return observable.map((response: Response) => {
            return response.json();
        } )
    }

    update(account: Account) {
        return this.http.put('http://localhost:8080/auth/accounts/' + account.id, account, this.jwt()).map((response: Response) => response.json());
    }

    delete(id: number) {
        return this.http.delete('http://localhost:8080/auth/accounts/' + id, this.jwt()).map((response: Response) => response.json());
    }

    // private helper methods

    private jwt() {
        // create authorization header with jwt token
        let currentAccount = JSON.parse(localStorage.getItem('currentAccount'));
        if (currentAccount && currentAccount.token) {
            let headers = new Headers({'Authorization': 'Bearer ' + currentAccount.token});
            return new RequestOptions({headers: headers});
        }
    }
}
