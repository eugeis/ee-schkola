import {ExpensePurpose} from '@schkola/finance/FinanceApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';

@Injectable()
export class ExpensePurposeDataService extends TableDataService {
    itemName = 'expensepurpose';

    pageName = 'ExpensePurposeComponent';

    getFirst() {
        return new ExpensePurpose();
    }
}




