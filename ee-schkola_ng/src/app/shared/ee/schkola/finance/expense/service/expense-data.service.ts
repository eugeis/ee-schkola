import {Expense} from '@schkola/finance/FinanceApiBase'

import {Injectable} from '@angular/core';
import {TableDataService} from '../../../../template/services/data.service';

@Injectable()
export class ExpenseDataService extends TableDataService {
    itemName = 'expense';

    pageName = 'ExpenseComponent';

    getFirst() {
        return new Expense();
    }
}




