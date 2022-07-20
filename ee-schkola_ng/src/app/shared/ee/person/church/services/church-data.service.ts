import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';
import {Church} from '../../../schkola/person/PersonApiBase';

@Injectable()
export class ChurchDataService extends TableDataService {
    itemName = 'church';

    formArrayType = ['string', 'string', 'string'];

    formArrayName = ['Actions', 'name', 'address', 'association'];

    pageElement = ['Person'];

    tabElement = ['Profile', 'Church'];

    pageName = 'ChurchComponent';

    getFirst() {
        return new Church();
    }
}
