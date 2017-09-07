import {environment} from '../../../environments/environment';

export class AppSettings {
    public static CURRENT_ACCOUNT = "CURRENT_ACCOUNT"
    public static HOST_API = environment.hostApi
    public static HTTP_COOKIE = environment.httpCookie
}
