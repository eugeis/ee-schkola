import {Injectable} from '@angular/core';
import {FormService} from '../../../template/services/form.service';
import {Profile} from '../../../schkola/person/PersonApiBase';

@Injectable()
export class ProfileViewService extends FormService {

    data: Map<String, Profile>;

    profile: Profile = new Profile();

    /*formArrayName = ['Actions', 'gender', 'firstname', 'lastname', 'birthName', 'birthday', 'street', 'suite', 'code', 'city', 'country'];
*/
    pageElement = ['Person'];

    tabElement = ['Profile', 'Church'];

    pageName = 'ProfileComponent';
}
