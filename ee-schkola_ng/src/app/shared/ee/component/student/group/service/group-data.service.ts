import {Group} from '../student/StudentApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class GroupViewComponent extends TableDataService {
    itemName = 'group';

    pageName = 'GroupComponent';

    getFirst() {
        return new Group();
    }
}




