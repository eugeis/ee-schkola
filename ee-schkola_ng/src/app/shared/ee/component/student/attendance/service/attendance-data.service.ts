import {Attendance} from '../student/StudentApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class AttendanceViewComponent extends TableDataService {
    itemName = 'attendance';

    pageName = 'AttendanceComponent';

    getFirst() {
        return new Attendance();
    }
}




