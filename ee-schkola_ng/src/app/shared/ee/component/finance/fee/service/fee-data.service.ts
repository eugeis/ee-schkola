import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class FeeViewComponent extends TableDataService {
    itemName = 'fee';

    pageName = 'FeeComponent';

    getFirst() {
        return new Fee();
    }
}




