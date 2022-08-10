import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';
import {Profile} from '../../../schkola/person/PersonApiBase';

@Injectable()
export class ProfileDataService extends TableDataService {
    itemName = 'profile';

    pageName = 'ProfileComponent';

    getFirst() {
        return new Profile();
    }
}
