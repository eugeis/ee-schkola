import {Graduation} from '@schkola/person/PersonApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';

@Injectable()
export class GraduationDataService extends TableDataService {
    itemName = 'graduation';

    pageName = 'GraduationComponent';

    getFirst() {
        return new Graduation();
    }
}




