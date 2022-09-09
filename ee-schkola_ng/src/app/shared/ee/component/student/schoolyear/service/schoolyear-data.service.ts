import {SchoolYear} from '../student/StudentApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class SchoolYearViewComponent extends TableDataService {
    itemName = 'schoolyear';

    pageName = 'SchoolYearComponent';

    getFirst() {
        return new SchoolYear();
    }
}




