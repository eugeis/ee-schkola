import {SchoolApplication} from '@schkola/student/StudentApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';

@Injectable()
export class SchoolApplicationDataService extends TableDataService {
    itemName = 'schoolapplication';

    pageName = 'SchoolApplicationComponent';

    getFirst() {
        return new SchoolApplication();
    }
}




