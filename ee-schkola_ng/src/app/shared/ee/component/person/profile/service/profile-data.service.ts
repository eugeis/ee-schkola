import {Profile} from '../person/PersonApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class ProfileViewComponent extends TableDataService {
    itemName = 'profile';

    pageName = 'ProfileComponent';

    getFirst() {
        return new Profile();
    }
}




