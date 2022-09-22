import {Course} from '@schkola/student/StudentApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';

@Injectable()
export class CourseDataService extends TableDataService {
    itemName = 'course';

    pageName = 'CourseComponent';

    getFirst() {
        return new Course();
    }
}




