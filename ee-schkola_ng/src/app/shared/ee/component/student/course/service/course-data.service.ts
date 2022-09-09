import {Course} from '../student/StudentApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class CourseViewComponent extends TableDataService {
    itemName = 'course';

    pageName = 'CourseComponent';

    getFirst() {
        return new Course();
    }
}




