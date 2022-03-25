import {Injectable} from '@angular/core';
import {Http, RequestOptions, Response} from '@angular/http';
import {AppSettings} from './common.service';
import {map} from 'rxjs/operators';

@Injectable()
export class AuthService {
    loginUrl = AppSettings.HOST_API + '/login'
    logoutUrl = AppSettings.HOST_API + '/logout'

    constructor(private http: Http) {
    }

    login(username: string, password: string) {
        const myParams = new URLSearchParams()
        const options = new RequestOptions({params: myParams});
        options.params.set('username', username)
        options.params.set('password', password)
        return this.http.post(this.loginUrl,
            {}, options).pipe(map(
            (response: Response) => {
                return response.json();
            }));
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
