import {SchoolApplication} from '../student/StudentApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class SchoolApplicationViewComponent extends TableDataService {
    itemName = 'schoolapplication';

    pageName = 'SchoolApplicationComponent';

    getFirst() {
        return new SchoolApplication();
    }
}




