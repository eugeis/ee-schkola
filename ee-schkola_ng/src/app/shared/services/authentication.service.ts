import {Injectable} from '@angular/core';
import {Http, RequestOptions, Response} from '@angular/http';
import 'rxjs/add/operator/map'
import {printLine} from "tslint/lib/test/lines";

@Injectable()
export class AuthenticationService {
    constructor(private http: Http) {
    }

    login(username: string, password: string) {
        let myParams = new URLSearchParams()
        let options = new RequestOptions({params: myParams});
        options.params.set("username", username)
        options.params.set("password", password)
        return this.http.post("http://localhost:8080/login",
            {}, options).map(
            (response: Response) => {
                return response.json();
            });
    }

    logout() {
        // remove account from local storage to log account out
        localStorage.removeItem('currentAccount');
        return this.http.post("http://localhost:8080/logout", {}).subscribe();
    }
}
