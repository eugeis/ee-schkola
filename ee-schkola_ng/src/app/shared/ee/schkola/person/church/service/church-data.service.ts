import {Church} from '@schkola/person/PersonApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';

@Injectable()
export class ChurchDataService extends TableDataService {
    itemName = 'church';

    pageName = 'ChurchComponent';

    getFirst() {
        return new Church();
    }
}




