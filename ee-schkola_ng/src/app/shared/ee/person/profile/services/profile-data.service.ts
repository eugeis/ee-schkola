import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';
import {Profile} from '../../../schkola/person/PersonApiBase';

@Injectable()
export class ProfileDataService extends TableDataService {
    itemName = 'profile';

    dataName = 'profileData';

    pageElement = ['Person'];

    tabElement = ['Profile', 'Church'];

    pageName = 'ProfileComponent';

    getFirst() {
        return new Profile();
    }
}
