import {Injectable} from '@angular/core';
import {Http, RequestOptions, Response} from '@angular/http';
import 'rxjs/add/operator/map'
import {AppSettings} from "./common.service";

@Injectable()
export class AuthService {
    loginUrl = AppSettings.HOST_API + "/login"
    logoutUrl = AppSettings.HOST_API + "/logout"

    constructor(private http: Http) {
    }

    login(username: string, password: string) {
        let myParams = new URLSearchParams()
        let options = new RequestOptions({params: myParams});
        options.params.set("username", username)
        options.params.set("password", password)
        return this.http.post(this.loginUrl,
            {}, options).map(
            (response: Response) => {
                return response.json();
            });
    }

    logout() {
        this.onLogout()
        // remove account from local storage to log account out
        localStorage.removeItem(AppSettings.CURRENT_ACCOUNT);
    }

    onLogout() {
        // remove account from local storage to log account out
        localStorage.removeItem(AppSettings.CURRENT_ACCOUNT);
    }
}
