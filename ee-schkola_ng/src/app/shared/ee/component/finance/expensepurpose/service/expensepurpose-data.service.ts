import {Injectable} from '@angular/core';
import {TableDataService} from '../../../template/services/data.service';

@Injectable()
export class ExpensePurposeViewComponent extends TableDataService {
    itemName = 'expensepurpose';

    pageName = 'ExpensePurposeComponent';

    getFirst() {
        return new ExpensePurpose();
    }
}




