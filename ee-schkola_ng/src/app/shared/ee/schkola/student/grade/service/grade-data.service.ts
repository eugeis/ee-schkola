import {Grade} from '@schkola/student/StudentApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';

@Injectable()
export class GradeDataService extends TableDataService {
    itemName = 'grade';

    pageName = 'GradeComponent';

    getFirst() {
        return new Grade();
    }
}




