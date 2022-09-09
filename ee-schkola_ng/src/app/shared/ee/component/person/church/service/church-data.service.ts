import {Church} from '../person/PersonApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class ChurchViewComponent extends TableDataService {
    itemName = 'church';

    pageName = 'ChurchComponent';

    getFirst() {
        return new Church();
    }
}




