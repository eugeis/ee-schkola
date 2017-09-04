import {Injectable} from '@angular/core';
import {Http, Response} from '@angular/http';
import 'rxjs/add/operator/map'

@Injectable()
export class AuthenticationService {
    constructor(private http: Http) {
    }

    login(username: string, password: string) {
        let p = new URLSearchParams()
        p.set("command", "login")
        p.set("password", password)
        let test = p.get("command")
        return this.http.post("http://localhost:8080/auth/accounts/" + username,
            {}, {params: p}).map(
            (response: Response) => {
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
