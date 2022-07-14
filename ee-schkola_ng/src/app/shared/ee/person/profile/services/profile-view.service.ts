import {Injectable} from '@angular/core';
import {FormService} from '../../../template/services/form.service';
import {Profile} from '../../../schkola/person/PersonApiBase';

@Injectable()
export class ProfileViewService extends FormService {

    // TODO: Implement Mapping
    data: Map<String, Profile>;

    profile: Profile = new Profile();

    /*Static
    formArrayName = ['Actions', 'gender', 'firstname', 'lastname', 'birthName', 'birthday', 'street', 'suite', 'code', 'city', 'country'];*/

    pageElement = ['Person'];

    tabElement = ['Profile', 'Church'];

    pageName = 'ProfileComponent';

    // Dynamic
    generateTableHeader() {
        const formArrayName = ['Actions'];
        Object.keys(this.profile).map((profileIndex) => {
            typeof this.profile[profileIndex] === 'object' ?
                this.profile[profileIndex] instanceof Date ?
                    formArrayName.push(profileIndex) :
                    Object.keys(this.profile[profileIndex]).forEach((element) => {
                    formArrayName.push(element);
                }) : formArrayName.push(profileIndex);
        })
        return formArrayName;
    }
}
