import {Grade} from '../student/StudentApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class GradeViewComponent extends TableDataService {
    itemName = 'grade';

    pageName = 'GradeComponent';

    getFirst() {
        return new Grade();
    }
}




