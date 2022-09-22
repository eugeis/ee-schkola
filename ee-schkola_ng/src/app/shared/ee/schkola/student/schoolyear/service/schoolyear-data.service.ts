import {SchoolYear} from '@schkola/student/StudentApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';

@Injectable()
export class SchoolYearDataService extends TableDataService {
    itemName = 'schoolyear';

    pageName = 'SchoolYearComponent';

    getFirst() {
        return new SchoolYear();
    }
}




