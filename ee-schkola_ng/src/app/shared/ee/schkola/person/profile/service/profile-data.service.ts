import {Profile} from '@schkola/person/PersonApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';

@Injectable()
export class ProfileDataService extends TableDataService {
    itemName = 'profile';

    pageName = 'ProfileComponent';

    getFirst() {
        return new Profile();
    }
}




