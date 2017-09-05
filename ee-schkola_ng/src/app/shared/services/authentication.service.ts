import {Injectable} from '@angular/core';
import {Http, RequestOptions, Response} from '@angular/http';
import 'rxjs/add/operator/map'

@Injectable()
export class AuthenticationService {
    constructor(private http: Http) {
    }

    login(username: string, password: string) {
        let myParams = new URLSearchParams()
        let options = new RequestOptions({params: myParams});
        options.params.set("command", "login")
        options.params.set("password", password)
        return this.http.post("http://localhost:8080/auth/accounts/" + username,
            {}, options).map(
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
